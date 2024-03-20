package CQHTTP

type CQHTTP struct {
	Secret       string `toml:"SECRET"`
	GroupID      []int  `toml:"GROUP_ID"`
	KeyWord      string `toml:"KEY_WORD"`
	FastReplayOk string `toml:"FAST_REPLAY_OK"`
	FastReplayNo string `toml:"FAST_REPLAY_NO"`
}

type heartbeat struct {
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	Time          int    `json:"time"`
	SelfId        int    `json:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
			PacketLost      int `json:"packet_lost"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			DisconnectTimes int `json:"disconnect_times"`
			LostTimes       int `json:"lost_times"`
			LastMessageTime int `json:"last_message_time"`
		} `json:"stat"`
	} `json:"status"`
	Interval int `json:"interval"`
}
type groupMsg struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	Time        int    `json:"time"`
	SelfID      int    `json:"self_id"`
	SubType     string `json:"sub_type"`
	GroupID     int    `json:"group_id"`
	RawMessage  string `json:"raw_message"`
	MessageID   int    `json:"message_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Ca       string `json:"ca"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserID   int64  `json:"user_id"`
	} `json:"sender"`
	UserID     int64  `json:"user_id"`
	Anonymous  any    `json:"anonymous"`
	Font       int    `json:"font"`
	Message    string `json:"message"`
	MessageSeq int    `json:"message_seq"`
}
