package client

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type VideoMedia struct {
	Label      string     `json:"label"`
	Codec      string     `json:"codec"`
	Coder      string     `json:"coder,omitempty"`
	Bitrate    int        `json:"bitrate"`
	Framerate  string     `json:"framerate"`
	Resolution Resolution `json:"resolution"`
}

type AudioMedia struct {
	Uuid             string `json:"uuid,omitempty"`
	Codec            string `json:"codec"`
	Channels         string `json:"channels"`
	Bitrate          int    `json:"bitrate"`
	Samplerate       string `json:"samplerate"`
	Track            string `json:"track"`
	Output           string `json:"output,"`
	OutputLabel      string `json:"output_label"`
	AudioDescription bool   `json:"audio_description"`
	Label            string `json:"label"`
}

type SubtitleMedia struct {
	Uuid                 string `json:"uuid,omitempty"`
	Track                string `json:"track"`
	Bitrate              int    `json:"bitrate"`
	Output               string `json:"output"`
	OutputLabel          string `json:"output_label"`
	DeafAndHardOfHearing bool   `json:"deaf_and_hard_of_hearing"`
}

type ProcessingPresets struct {
	Uuid           string          `json:"uuid,omitempty"`
	Name           string          `json:"name"`
	VideoMedias    []VideoMedia    `json:"video_medias"`
	AudioMedias    []AudioMedia    `json:"audio_medias"`
	SubtitleMedias []SubtitleMedia `json:"subtitle_medias"`
}
