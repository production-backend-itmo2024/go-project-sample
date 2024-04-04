package kafka

type Message struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}
