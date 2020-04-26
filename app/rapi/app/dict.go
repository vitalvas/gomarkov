package app

// Message godoc
type Message struct {
	Message string `json:"message"`
}

// BatchMessage godoc
type BatchMessage struct {
	Messages []string `json:"messages"`
}
