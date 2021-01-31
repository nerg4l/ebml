// Code generated by go run make_schema.go. DO NOT EDIT.

package matroska

import "time"

type DocTypeExtension struct {
	DocTypeExtensionName    string `ebml:"0x4283"`
	DocTypeExtensionVersion uint   `ebml:"0x4284"`
}

type EBML struct {
	EBMLVersion        uint               `ebml:"0x4286"`
	EBMLReadVersion    uint               `ebml:"0x42F7"`
	EBMLMaxIDLength    uint               `ebml:"0x42F2"`
	EBMLMaxSizeLength  uint               `ebml:"0x42F3"`
	DocType            string             `ebml:"0x4282"`
	DocTypeVersion     uint               `ebml:"0x4287"`
	DocTypeReadVersion uint               `ebml:"0x4285"`
	DocTypeExtension   []DocTypeExtension `ebml:"0x4281"`
}

type Seek struct {
	SeekID       []byte `ebml:"0x53AB"`
	SeekPosition uint   `ebml:"0x53AC"`
}

type SeekHead struct {
	Seek []Seek `ebml:"0x4DBB"`
}

type ChapterTranslate struct {
	ChapterTranslateEditionUID []uint `ebml:"0x69FC"`
	ChapterTranslateCodec      uint   `ebml:"0x69BF"`
	ChapterTranslateID         []byte `ebml:"0x69A5"`
}

type Info struct {
	SegmentUID       []byte             `ebml:"0x73A4"`
	SegmentFilename  []byte             `ebml:"0x7384"`
	PrevUID          []byte             `ebml:"0x3CB923"`
	PrevFilename     []byte             `ebml:"0x3C83AB"`
	NextUID          []byte             `ebml:"0x3EB923"`
	NextFilename     []byte             `ebml:"0x3E83BB"`
	SegmentFamily    [][]byte           `ebml:"0x4444"`
	ChapterTranslate []ChapterTranslate `ebml:"0x6924"`
	TimestampScale   uint               `ebml:"0x2AD7B1"`
	Duration         float64            `ebml:"0x4489"`
	DateUTC          time.Time          `ebml:"0x4461"`
	Title            []byte             `ebml:"0x7BA9"`
	MuxingApp        []byte             `ebml:"0x4D80"`
	WritingApp       []byte             `ebml:"0x5741"`
}

type SilentTracks struct {
	SilentTrackNumber []uint `ebml:"0x58D7"`
}

type BlockMore struct {
	BlockAddID      uint   `ebml:"0xEE"`
	BlockAdditional []byte `ebml:"0xA5"`
}

type BlockAdditions struct {
	BlockMore []BlockMore `ebml:"0xA6"`
}

type TimeSlice struct {
	LaceNumber      uint `ebml:"0xCC"`
	FrameNumber     uint `ebml:"0xCD"`
	BlockAdditionID uint `ebml:"0xCB"`
	Delay           uint `ebml:"0xCE"`
	SliceDuration   uint `ebml:"0xCF"`
}

type Slices struct {
	TimeSlice []TimeSlice `ebml:"0xE8"`
}

type ReferenceFrame struct {
	ReferenceOffset    uint `ebml:"0xC9"`
	ReferenceTimestamp uint `ebml:"0xCA"`
}

type BlockGroup struct {
	Block             []byte         `ebml:"0xA1"`
	BlockVirtual      []byte         `ebml:"0xA2"`
	BlockAdditions    BlockAdditions `ebml:"0x75A1"`
	BlockDuration     uint           `ebml:"0x9B"`
	ReferencePriority uint           `ebml:"0xFA"`
	ReferenceBlock    []int          `ebml:"0xFB"`
	ReferenceVirtual  int            `ebml:"0xFD"`
	CodecState        []byte         `ebml:"0xA4"`
	DiscardPadding    int            `ebml:"0x75A2"`
	Slices            Slices         `ebml:"0x8E"`
	ReferenceFrame    ReferenceFrame `ebml:"0xC8"`
}

type Cluster struct {
	Timestamp      uint         `ebml:"0xE7"`
	SilentTracks   SilentTracks `ebml:"0x5854"`
	Position       uint         `ebml:"0xA7"`
	PrevSize       uint         `ebml:"0xAB"`
	SimpleBlock    [][]byte     `ebml:"0xA3"`
	BlockGroup     []BlockGroup `ebml:"0xA0"`
	EncryptedBlock [][]byte     `ebml:"0xAF"`
}

type BlockAdditionMapping struct {
	BlockAddIDValue     uint   `ebml:"0x41F0"`
	BlockAddIDName      string `ebml:"0x41A4"`
	BlockAddIDType      uint   `ebml:"0x41E7"`
	BlockAddIDExtraData []byte `ebml:"0x41ED"`
}

type TrackTranslate struct {
	TrackTranslateEditionUID []uint `ebml:"0x66FC"`
	TrackTranslateCodec      uint   `ebml:"0x66BF"`
	TrackTranslateTrackID    []byte `ebml:"0x66A5"`
}

type MasteringMetadata struct {
	PrimaryRChromaticityX   float64 `ebml:"0x55D1"`
	PrimaryRChromaticityY   float64 `ebml:"0x55D2"`
	PrimaryGChromaticityX   float64 `ebml:"0x55D3"`
	PrimaryGChromaticityY   float64 `ebml:"0x55D4"`
	PrimaryBChromaticityX   float64 `ebml:"0x55D5"`
	PrimaryBChromaticityY   float64 `ebml:"0x55D6"`
	WhitePointChromaticityX float64 `ebml:"0x55D7"`
	WhitePointChromaticityY float64 `ebml:"0x55D8"`
	LuminanceMax            float64 `ebml:"0x55D9"`
	LuminanceMin            float64 `ebml:"0x55DA"`
}

type Colour struct {
	MatrixCoefficients      uint              `ebml:"0x55B1"`
	BitsPerChannel          uint              `ebml:"0x55B2"`
	ChromaSubsamplingHorz   uint              `ebml:"0x55B3"`
	ChromaSubsamplingVert   uint              `ebml:"0x55B4"`
	CbSubsamplingHorz       uint              `ebml:"0x55B5"`
	CbSubsamplingVert       uint              `ebml:"0x55B6"`
	ChromaSitingHorz        uint              `ebml:"0x55B7"`
	ChromaSitingVert        uint              `ebml:"0x55B8"`
	Range                   uint              `ebml:"0x55B9"`
	TransferCharacteristics uint              `ebml:"0x55BA"`
	Primaries               uint              `ebml:"0x55BB"`
	MaxCLL                  uint              `ebml:"0x55BC"`
	MaxFALL                 uint              `ebml:"0x55BD"`
	MasteringMetadata       MasteringMetadata `ebml:"0x55D0"`
}

type Projection struct {
	ProjectionType      uint    `ebml:"0x7671"`
	ProjectionPrivate   []byte  `ebml:"0x7672"`
	ProjectionPoseYaw   float64 `ebml:"0x7673"`
	ProjectionPosePitch float64 `ebml:"0x7674"`
	ProjectionPoseRoll  float64 `ebml:"0x7675"`
}

type Video struct {
	FlagInterlaced  uint       `ebml:"0x9A"`
	FieldOrder      uint       `ebml:"0x9D"`
	StereoMode      uint       `ebml:"0x53B8"`
	AlphaMode       uint       `ebml:"0x53C0"`
	OldStereoMode   uint       `ebml:"0x53B9"`
	PixelWidth      uint       `ebml:"0xB0"`
	PixelHeight     uint       `ebml:"0xBA"`
	PixelCropBottom uint       `ebml:"0x54AA"`
	PixelCropTop    uint       `ebml:"0x54BB"`
	PixelCropLeft   uint       `ebml:"0x54CC"`
	PixelCropRight  uint       `ebml:"0x54DD"`
	DisplayWidth    uint       `ebml:"0x54B0"`
	DisplayHeight   uint       `ebml:"0x54BA"`
	DisplayUnit     uint       `ebml:"0x54B2"`
	AspectRatioType uint       `ebml:"0x54B3"`
	ColourSpace     []byte     `ebml:"0x2EB524"`
	GammaValue      float64    `ebml:"0x2FB523"`
	FrameRate       float64    `ebml:"0x2383E3"`
	Colour          Colour     `ebml:"0x55B0"`
	Projection      Projection `ebml:"0x7670"`
}

type Audio struct {
	SamplingFrequency       float64 `ebml:"0xB5"`
	OutputSamplingFrequency float64 `ebml:"0x78B5"`
	Channels                uint    `ebml:"0x9F"`
	ChannelPositions        []byte  `ebml:"0x7D7B"`
	BitDepth                uint    `ebml:"0x6264"`
}

type TrackPlane struct {
	TrackPlaneUID  uint `ebml:"0xE5"`
	TrackPlaneType uint `ebml:"0xE6"`
}

type TrackCombinePlanes struct {
	TrackPlane []TrackPlane `ebml:"0xE4"`
}

type TrackJoinBlocks struct {
	TrackJoinUID []uint `ebml:"0xED"`
}

type TrackOperation struct {
	TrackCombinePlanes TrackCombinePlanes `ebml:"0xE3"`
	TrackJoinBlocks    TrackJoinBlocks    `ebml:"0xE9"`
}

type ContentCompression struct {
	ContentCompAlgo     uint   `ebml:"0x4254"`
	ContentCompSettings []byte `ebml:"0x4255"`
}

type ContentEncAESSettings struct {
	AESSettingsCipherMode uint `ebml:"0x47E8"`
}

type ContentEncryption struct {
	ContentEncAlgo        uint                  `ebml:"0x47E1"`
	ContentEncKeyID       []byte                `ebml:"0x47E2"`
	ContentEncAESSettings ContentEncAESSettings `ebml:"0x47E7"`
	ContentSignature      []byte                `ebml:"0x47E3"`
	ContentSigKeyID       []byte                `ebml:"0x47E4"`
	ContentSigAlgo        uint                  `ebml:"0x47E5"`
	ContentSigHashAlgo    uint                  `ebml:"0x47E6"`
}

type ContentEncoding struct {
	ContentEncodingOrder uint               `ebml:"0x5031"`
	ContentEncodingScope uint               `ebml:"0x5032"`
	ContentEncodingType  uint               `ebml:"0x5033"`
	ContentCompression   ContentCompression `ebml:"0x5034"`
	ContentEncryption    ContentEncryption  `ebml:"0x5035"`
}

type ContentEncodings struct {
	ContentEncoding []ContentEncoding `ebml:"0x6240"`
}

type TrackEntry struct {
	TrackNumber                 uint                   `ebml:"0xD7"`
	TrackUID                    uint                   `ebml:"0x73C5"`
	TrackType                   uint                   `ebml:"0x83"`
	FlagEnabled                 uint                   `ebml:"0xB9"`
	FlagDefault                 uint                   `ebml:"0x88"`
	FlagForced                  uint                   `ebml:"0x55AA"`
	FlagLacing                  uint                   `ebml:"0x9C"`
	MinCache                    uint                   `ebml:"0x6DE7"`
	MaxCache                    uint                   `ebml:"0x6DF8"`
	DefaultDuration             uint                   `ebml:"0x23E383"`
	DefaultDecodedFieldDuration uint                   `ebml:"0x234E7A"`
	TrackTimestampScale         float64                `ebml:"0x23314F"`
	TrackOffset                 int                    `ebml:"0x537F"`
	MaxBlockAdditionID          uint                   `ebml:"0x55EE"`
	BlockAdditionMapping        []BlockAdditionMapping `ebml:"0x41E4"`
	Name                        []byte                 `ebml:"0x536E"`
	Language                    string                 `ebml:"0x22B59C"`
	LanguageIETF                string                 `ebml:"0x22B59D"`
	CodecID                     string                 `ebml:"0x86"`
	CodecPrivate                []byte                 `ebml:"0x63A2"`
	CodecName                   []byte                 `ebml:"0x258688"`
	AttachmentLink              uint                   `ebml:"0x7446"`
	CodecSettings               []byte                 `ebml:"0x3A9697"`
	CodecInfoURL                []string               `ebml:"0x3B4040"`
	CodecDownloadURL            []string               `ebml:"0x26B240"`
	CodecDecodeAll              uint                   `ebml:"0xAA"`
	TrackOverlay                []uint                 `ebml:"0x6FAB"`
	CodecDelay                  uint                   `ebml:"0x56AA"`
	SeekPreRoll                 uint                   `ebml:"0x56BB"`
	TrackTranslate              []TrackTranslate       `ebml:"0x6624"`
	Video                       Video                  `ebml:"0xE0"`
	Audio                       Audio                  `ebml:"0xE1"`
	TrackOperation              TrackOperation         `ebml:"0xE2"`
	TrickTrackUID               uint                   `ebml:"0xC0"`
	TrickTrackSegmentUID        []byte                 `ebml:"0xC1"`
	TrickTrackFlag              uint                   `ebml:"0xC6"`
	TrickMasterTrackUID         uint                   `ebml:"0xC7"`
	TrickMasterTrackSegmentUID  []byte                 `ebml:"0xC4"`
	ContentEncodings            ContentEncodings       `ebml:"0x6D80"`
}

type Tracks struct {
	TrackEntry []TrackEntry `ebml:"0xAE"`
}

type CueReference struct {
	CueRefTime       uint `ebml:"0x96"`
	CueRefCluster    uint `ebml:"0x97"`
	CueRefNumber     uint `ebml:"0x535F"`
	CueRefCodecState uint `ebml:"0xEB"`
}

type CueTrackPositions struct {
	CueTrack            uint           `ebml:"0xF7"`
	CueClusterPosition  uint           `ebml:"0xF1"`
	CueRelativePosition uint           `ebml:"0xF0"`
	CueDuration         uint           `ebml:"0xB2"`
	CueBlockNumber      uint           `ebml:"0x5378"`
	CueCodecState       uint           `ebml:"0xEA"`
	CueReference        []CueReference `ebml:"0xDB"`
}

type CuePoint struct {
	CueTime           uint                `ebml:"0xB3"`
	CueTrackPositions []CueTrackPositions `ebml:"0xB7"`
}

type Cues struct {
	CuePoint []CuePoint `ebml:"0xBB"`
}

type AttachedFile struct {
	FileDescription   []byte `ebml:"0x467E"`
	FileName          []byte `ebml:"0x466E"`
	FileMimeType      string `ebml:"0x4660"`
	FileData          []byte `ebml:"0x465C"`
	FileUID           uint   `ebml:"0x46AE"`
	FileReferral      []byte `ebml:"0x4675"`
	FileUsedStartTime uint   `ebml:"0x4661"`
	FileUsedEndTime   uint   `ebml:"0x4662"`
}

type Attachments struct {
	AttachedFile []AttachedFile `ebml:"0x61A7"`
}

type ChapterTrack struct {
	ChapterTrackUID []uint `ebml:"0x89"`
}

type ChapterDisplay struct {
	ChapString       []byte   `ebml:"0x85"`
	ChapLanguage     []string `ebml:"0x437C"`
	ChapLanguageIETF []string `ebml:"0x437D"`
	ChapCountry      []string `ebml:"0x437E"`
}

type ChapProcessCommand struct {
	ChapProcessTime uint   `ebml:"0x6922"`
	ChapProcessData []byte `ebml:"0x6933"`
}

type ChapProcess struct {
	ChapProcessCodecID uint                 `ebml:"0x6955"`
	ChapProcessPrivate []byte               `ebml:"0x450D"`
	ChapProcessCommand []ChapProcessCommand `ebml:"0x6911"`
}

type ChapterAtom struct {
	ChapterUID               uint             `ebml:"0x73C4"`
	ChapterStringUID         []byte           `ebml:"0x5654"`
	ChapterTimeStart         uint             `ebml:"0x91"`
	ChapterTimeEnd           uint             `ebml:"0x92"`
	ChapterFlagHidden        uint             `ebml:"0x98"`
	ChapterFlagEnabled       uint             `ebml:"0x4598"`
	ChapterSegmentUID        []byte           `ebml:"0x6E67"`
	ChapterSegmentEditionUID uint             `ebml:"0x6EBC"`
	ChapterPhysicalEquiv     uint             `ebml:"0x63C3"`
	ChapterTrack             ChapterTrack     `ebml:"0x8F"`
	ChapterDisplay           []ChapterDisplay `ebml:"0x80"`
	ChapProcess              []ChapProcess    `ebml:"0x6944"`
}

type EditionEntry struct {
	EditionUID         uint          `ebml:"0x45BC"`
	EditionFlagHidden  uint          `ebml:"0x45BD"`
	EditionFlagDefault uint          `ebml:"0x45DB"`
	EditionFlagOrdered uint          `ebml:"0x45DD"`
	ChapterAtom        []ChapterAtom `ebml:"0xB6"`
}

type Chapters struct {
	EditionEntry []EditionEntry `ebml:"0x45B9"`
}

type Targets struct {
	TargetTypeValue  uint   `ebml:"0x68CA"`
	TargetType       string `ebml:"0x63CA"`
	TagTrackUID      []uint `ebml:"0x63C5"`
	TagEditionUID    []uint `ebml:"0x63C9"`
	TagChapterUID    []uint `ebml:"0x63C4"`
	TagAttachmentUID []uint `ebml:"0x63C6"`
}

type SimpleTag struct {
	TagName         []byte `ebml:"0x45A3"`
	TagLanguage     string `ebml:"0x447A"`
	TagLanguageIETF string `ebml:"0x447B"`
	TagDefault      uint   `ebml:"0x4484"`
	TagString       []byte `ebml:"0x4487"`
	TagBinary       []byte `ebml:"0x4485"`
}

type Tag struct {
	Targets   Targets     `ebml:"0x63C0"`
	SimpleTag []SimpleTag `ebml:"0x67C8"`
}

type Tags struct {
	Tag []Tag `ebml:"0x7373"`
}

type Segment struct {
	SeekHead    []SeekHead  `ebml:"0x114D9B74"`
	Info        Info        `ebml:"0x1549A966"`
	Cluster     []Cluster   `ebml:"0x1F43B675"`
	Tracks      Tracks      `ebml:"0x1654AE6B"`
	Cues        Cues        `ebml:"0x1C53BB6B"`
	Attachments Attachments `ebml:"0x1941A469"`
	Chapters    Chapters    `ebml:"0x1043A770"`
	Tags        []Tags      `ebml:"0x1254C367"`
}

type Document struct {
	EBML    EBML    `ebml:"0x1A45DFA3"`
	Segment Segment `ebml:"0x18538067"`
}