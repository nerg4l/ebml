package ebml

var defaultHeader = EBML{
	EBMLVersion:        1,
	EBMLReadVersion:    1,
	EBMLMaxIDLength:    4,
	EBMLMaxSizeLength:  8,
	DocType:            "ebml",
	DocTypeVersion:     1,
	DocTypeReadVersion: 1,
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

type DocTypeExtension struct {
	DocTypeExtensionName    string `ebml:"0x4283"`
	DocTypeExtensionVersion uint   `ebml:"0x4284"`
}