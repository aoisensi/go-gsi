package dota2

import (
	"encoding/json"
	"fmt"
)

func NewHandler(r func(*Data)) *Handler {
	return &Handler{Receiver: r}
}

type Handler struct {
	Receiver func(*Data)
}

func (h *Handler) AppID() []int {
	return []int{570}
}

func (h Handler) Receive(body []byte) error {
	var data *Data
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return err
	}
	h.Receiver(data)
	return nil
}
