package ebml

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"reflect"
	"time"
)

// An DecodeTypeError describes an EBML value that was
// not appropriate for a value of a specific Go type.
type DecodeTypeError struct {
	EBMLType string       // description of EBML type - "integer", "binary", "master"
	Type     reflect.Type // type of Go value it could not be assigned to
	Offset   int64        // error occurred after reading Offset bytes
	Path     string       // the full path from root node to the field
}

func (e *DecodeTypeError) Error() string {
	if e.Path != "" {
		return "ebml: cannot unmarshal " + e.EBMLType + " into Go struct field " + e.Path + " of type " + e.Type.String()
	}
	return "ebml: cannot unmarshal " + e.EBMLType + " into Go value of type " + e.Type.String()
}

func (e *DecodeTypeError) extendError(p string) {
	if e.Path == "" {
		e.Path = p
		return
	}
	e.Path = p + "." + e.Path
}

// An InvalidDecodeError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidDecodeError struct {
	Type reflect.Type
}

func (e *InvalidDecodeError) Error() string {
	if e.Type == nil {
		return "ebml: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "ebml: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "ebml: Unmarshal(nil " + e.Type.String() + ")"
}

// DecodeHeader decodes the document header.
func (d *Decoder) DecodeHeader() (EBML, error) {
	var v EBML
	val := reflect.ValueOf(&v)

	if err := d.decodeRoot(val.Elem(), HeaderDocType); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return EBML{}, err
	}
	return v, nil
}

// DecodeBody decodes the EBML Body and stores the result in the value
// pointed to by v. If v is nil or not a pointer, DecodeBody returns
// an InvalidDecodeError.
func (d *Decoder) DecodeBody(header EBML, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return &InvalidDecodeError{reflect.TypeOf(v)}
	}
	bodyDef, err := getDefinition(header.DocType)
	if err != nil {
		return err
	}
	if err := d.decodeRoot(val.Elem(), bodyDef); err != nil {
		if err == io.EOF {
			err = nil
		}
		return err
	}
	return nil
}

var (
	timeType        = reflect.TypeOf(time.Time{})
	thirdMillennium = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
)

func findField(val reflect.Value, tinfo *typeInfo, name string) (fieldv reflect.Value, found bool) {
	for i := range tinfo.fields {
		finfo := tinfo.fields[i]
		if name != finfo.name {
			continue
		}
		found = true
		fieldv = val.Field(finfo.idx[0])
		break
	}
	return
}

func (d *Decoder) decodeRoot(val reflect.Value, def Definition) error {
	// Load value from interface, but only if the result will be
	// usefully addressable.
	if val.Kind() == reflect.Interface && !val.IsNil() {
		e := val.Elem()
		if e.Kind() == reflect.Ptr && !e.IsNil() {
			val = e
		}
	}

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}

	el, err := d.element([]Definition{def})
	if err != nil {
		return err
	}
	if el.Definition.ID == CRC32.ID {
		return errors.New("ebml: unexpected crc-32 element")
	}
	if el.Definition.ID == Void.ID {
		// TODO: skip void elements
		return errors.New("ebml: unexpected void element")
	}
	if val.Kind() != reflect.Struct {
		return errors.New("ebml: unknown root element type: " + val.Type().String())
	}

	if err := d.decodeSingle(el, true, val, el.Definition.Children); err != nil {
		return err
	}
	return nil
}

func (d *Decoder) decodeMaster(val reflect.Value, ds dataSize, defs []Definition) error {
	// Load value from interface, but only if the result will be
	// usefully addressable.
	if val.Kind() == reflect.Interface && !val.IsNil() {
		e := val.Elem()
		if e.Kind() == reflect.Ptr && !e.IsNil() {
			val = e
		}
	}

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}

	switch v := val; v.Kind() {
	default:
		return errors.New("unknown master element type: " + val.Type().String())
	case reflect.Slice:
		// TODO: Consider checking max / min occurrence.
		e := v.Type().Elem()
		n := v.Len()
		v.Set(reflect.Append(v, reflect.Zero(e)))
		val = v.Index(n)
	case reflect.Struct:
		// Everything is ok
	}
	typ := val.Type()
	tinfo, err := getTypeInfo(typ)
	if err != nil {
		return err
	}

	occurrences := make(map[string]int, len(defs))
	start := d.r.Position()
	for {
		// Check size first because it can be 0
		if ds.Known() {
			offset := d.r.Position() - start
			if offset > ds.Size() {
				return errors.New("ebml: element overflow")
			} else if offset == ds.Size() {
				break
			}
		}

		el, err := d.element(defs)
		if err != nil {
			if err == errInvalidId {
				continue
			}
			var e *UnknownElementError
			if !ds.Known() && errors.As(err, &e) {
				d.elCache = &e.el
				break
			}
			return err
		}
		occurrences[el.ID]++
		fieldv, found := findField(val, tinfo, el.Definition.Name)

		if err := d.decodeSingle(el, found, fieldv, el.Definition.Children); err != nil {
			if e, ok := err.(*DecodeTypeError); ok {
				e.extendError(val.Type().Name())
			}
			return err
		}
	}

	for i := range defs {
		def := defs[i]
		if def.Default == nil || occurrences[def.ID] > 0 {
			continue
		}
		if def.Type == TypeMaster {
			// TODO: catch this when Doc Type is registered.
			panic("ebml: Master Elements MUST NOT declare a default value.")
		}
		fieldv, found := findField(val, tinfo, def.Name)
		if !found {
			continue
		}
		if v := fieldv; v.Kind() == reflect.Slice {
			e := v.Type().Elem()
			if !(def.Type == TypeBinary && e.Kind() == reflect.Uint8) {
				n := v.Len()
				v.Set(reflect.Append(v, reflect.Zero(e)))
				fieldv = v.Index(n)
			}
		}

		if err := validateReflectType(fieldv, def, d.r.Position()); err != nil {
			if e, ok := err.(*DecodeTypeError); ok {
				e.extendError(def.Name)
				e.extendError(val.Type().Name())
			}
			return err
		}
		fieldv.Set(reflect.ValueOf(def.Default))
	}
	return nil
}

func validateReflectType(v reflect.Value, def Definition, position int64) error {
	switch def.Type {
	default:
		return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}

	case TypeMaster:
		switch v.Kind() {
		default:
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		case reflect.Struct:
			// valid type
		}

	case TypeBinary:
		switch v.Kind() {
		default:
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		case reflect.Slice:
			e := v.Type().Elem()
			if e.Kind() != reflect.Uint8 {
				return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
			}
		}

	case TypeDate:
		switch v.Type() {
		default:
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		case timeType:
			// valid type
		}

	case TypeFloat:
		switch v.Kind() {
		default:
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		case reflect.Float32, reflect.Float64:
			// valid type
		}

	case TypeInteger:
		switch v.Kind() {
		default:
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		case reflect.Int, reflect.Int64, reflect.Int32:
			// valid type
		}

	case TypeUinteger:
		switch v.Kind() {
		default:
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		case reflect.Uint, reflect.Uint64, reflect.Uint32:
			// valid type
		}

	case TypeString, TypeUTF8:
		if v.Kind() != reflect.String {
			return &DecodeTypeError{EBMLType: def.Type, Type: v.Type(), Offset: position}
		}
	}
	return nil
}

func (d *Decoder) decodeSingle(el Element, found bool, val reflect.Value, defs []Definition) error {
	if v := val; v.Kind() == reflect.Slice {
		e := v.Type().Elem()
		if !(el.Definition.Type == TypeBinary && e.Kind() == reflect.Uint8) {
			n := v.Len()
			v.Set(reflect.Append(v, reflect.Zero(e)))
			val = v.Index(n)
		}
	}
	if found {
		if err := validateReflectType(val, el.Definition, d.r.Position()); err != nil {
			if e, ok := err.(*DecodeTypeError); ok {
				e.extendError(el.Definition.Name)
			}
			return err
		}
	}

	switch el.Definition.Type {
	case TypeMaster:
		if err := d.decodeMaster(val, el.DataSize, defs); err != nil {
			return err
		}

	case TypeBinary:
		b, err := d.readByteSlice(el.DataSize.Size())
		if err != nil {
			return err
		}
		if found {
			val.SetBytes(b)
		}

	case TypeDate:
		t, err := d.readDate(el.DataSize.Size())
		if err != nil {
			return err
		}
		if found {
			val.Set(reflect.ValueOf(t))
		}

	case TypeFloat:
		f, err := d.readFloat(el.DataSize.Size())
		if err != nil {
			return err
		}
		if found {
			val.SetFloat(f)
		}

	case TypeInteger:
		i, err := d.readInt(el.DataSize.Size())
		if err != nil {
			return err
		}
		if found {
			val.SetInt(i)
		}

	case TypeUinteger:
		i, err := d.readUint(el.DataSize.Size())
		if err != nil {
			return err
		}
		if found {
			val.SetUint(i)
		}

	case TypeString, TypeUTF8:
		str, err := d.readString(el.DataSize.Size())
		if err != nil {
			return err
		}
		if found {
			val.SetString(str)
		}
	}
	return nil
}

func (d *Decoder) readByteSlice(ds int64) ([]byte, error) {
	b := make([]byte, ds)
	_, err := io.ReadFull(d.r, b)
	return b, err
}

func (d *Decoder) readDate(ds int64) (time.Time, error) {
	i, err := d.readInt(ds)
	if err != nil {
		return time.Time{}, err
	}
	return thirdMillennium.Add(time.Nanosecond * time.Duration(i)), nil
}

func (d *Decoder) readString(ds int64) (string, error) {
	b, err := d.readByteSlice(ds)
	if err != nil {
		return "", err
	}
	// TODO: detect value greater than VINTMAX
	return string(b), err
}

func (d *Decoder) readFloat(ds int64) (float64, error) {
	b, err := d.readByteSlice(ds)
	if err != nil {
		return 0, err
	}
	// A Float Element MUST declare a length of either
	// zero octets (0 bit), four octets (32 bit),
	// or eight octets (64 bit).
	switch len(b) {
	case 0:
		return 0, nil
	case 4:
		return float64(math.Float32frombits(binary.BigEndian.Uint32(b))), nil
	case 8:
		return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
	default:
		return 0, errors.New("ebml: data length must be 0 bit, 32 bit or 64 bit for a float")
	}
}

func (d *Decoder) readInt(ds int64) (int64, error) {
	b, err := d.readByteSlice(ds)
	if err != nil {
		return 0, err
	}
	if len(b) > 8 {
		return 0, errors.New("ebml: max length for an unsigned integer is eight octets")
	}
	i := int64(0)
	for _, bb := range b {
		i = (i << 8) | int64(bb)
	}
	return i, nil
}

func (d *Decoder) readUint(ds int64) (uint64, error) {
	b, err := d.readByteSlice(ds)
	if err != nil {
		return 0, err
	}
	if len(b) > 8 {
		return 0, errors.New("ebml: max length for an unsigned integer is eight octets")
	}
	i := uint64(0)
	for _, bb := range b {
		i = (i << 8) | uint64(bb)
	}
	return i, nil
}
