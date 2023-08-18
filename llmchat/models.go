package llmchat

type Question struct {
	HumanChatMessage  string `json:"human_chat_message"`
	SystemChatMessage string `json:"system_chat_message,omitempty"`
	Temperature       int    `json:"temperature,omitempty"`
}

type Answer struct {
	AnswerText *string   `json:"answer_text,omitempty"`
}
