package MFRCScan

type MFRCScan struct {
	Secret    string `toml:"SECRET"`
	Tolerance int    `toml:"TOLERANCE"`
}

type msg struct {
	UID       string `json:"UID"`
	SAK       uint8  `json:"SAK"`
	Timestamp int32  `json:"timestamp"`
}
