package client

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// type SavcConfig struct {
// 	ForceSignalLevel     bool    `json:"force_signal_level,omitempty"`
// 	BufsizeRatio         float64 `json:"bufsize_ratio,omitempty"`
// 	RcInitOccupancy      float64 `json:"rc_init_occupancy,omitempty"`
// 	QualitySpeedOverride string  `json:"quality_speed_override,omitempty"`
// }

// TOOD à remettre ?
// type Advanced struct {
// 	Profile             string `json:"profile,omitempty"`
// 	Level               string `json:"level,omitempty"`
// 	Quality             string `json:"quality,omitempty"`
// 	EncodingMode        string `json:"encoding_mode,omitempty"`
// 	EncodingQuality     int    `json:"encoding_quality,omitempty"`
// 	QualityOptimization string `json:"quality_optimization,omitempty"`
// 	ClosedGop           bool   `json:"closed_gop,omitempty"`
// 	GopSize             int    `json:"gop_size,omitempty"`
// 	GopMaxSize          int    `json:"gop_max_size,omitempty"`
// 	Bframe              bool   `json:"bframe,omitempty"`
// 	BframeNumber        int    `json:"bframe_number,omitempty"`
// 	KeyFrameIntervalMs  int    `json:"key_frame_interval_ms,omitempty"`
// 	HorizontalSharpness int    `json:"horizontal_sharpness,omitempty"`
// 	VerticalSharpness   int    `json:"vertical_sharpness,omitempty"`
// 	LogoEnabled         bool   `json:"logo_enabled,omitempty"`
// 	// SavcConfig          SavcConfig `json:"savc_config,omitempty"`
// }

type VideoMedia struct {
	Label      string     `json:"label"`
	Codec      string     `json:"codec"`
	Coder      string     `json:"coder,omitempty"`
	Bitrate    int        `json:"bitrate"`
	Framerate  string     `json:"framerate"`
	Resolution Resolution `json:"resolution"`
	// TOOD à remettre ?
	// Advanced   Advanced   `json:"advanced"`
}

type AudioMedia struct {
	Uuid             string `json:"uuid,omitempty"`
	Codec            string `json:"codec"`
	Channels         string `json:"channels"`
	Bitrate          int    `json:"bitrate"`
	Samplerate       string `json:"samplerate"`
	Track            string `json:"track"`
	Output           string `json:"output"`
	OutputLabel      string `json:"output_label"`
	AudioDescription bool   `json:"audio_description"`
	Label            string `json:"label,"`
}

type SubtitleMedia struct {
	Uuid                 string `json:"uuid,omitempty"`
	Track                string `json:"track,omitempty"`
	Bitrate              int    `json:"bitrate"`
	Output               string `json:"output"`
	OutputLabel          string `json:"output_label"`
	DeafAndHardOfHearing bool   `json:"deaf_and_hard_of_hearing"`
}

type ProcessingPresets struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name"`
	// Identifier     string          `json:"identifier,omitempty"`
	// PoolUuid       string          `json:"pool_uuid"`
	// Published      bool            `json:"published,omitempty"`
	VideoMedias    []VideoMedia    `json:"video_medias,omitempty"`
	AudioMedias    []AudioMedia    `json:"audio_medias,omitempty"`
	SubtitleMedias []SubtitleMedia `json:"subtitle_medias,omitempty"`
	// Labels         []string        `json:"labels,omitempty"`
	ModifiedAt string `json:"modified_at,omitempty"`
}
