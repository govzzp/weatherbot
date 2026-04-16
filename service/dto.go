package service

type FeishuMessage struct {
	MsgType string     `json:"msg_type"`
	Card    FeishuCard `json:"card"`
}

type FeishuCard struct {
	Config   CardConfig `json:"config"`
	Header   CardHeader `json:"header"`
	Elements []Element  `json:"elements"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type CardHeader struct {
	Title    TextModule `json:"title"`
	Template string     `json:"template"`
}

type Element struct {
	Tag  string      `json:"tag"`
	Text *TextModule `json:"text,omitempty"`
}

type TextModule struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}
