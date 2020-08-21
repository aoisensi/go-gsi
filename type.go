package gsi

type Provider struct {
	Name    string
	AppID   int
	Version int
	//Timestamp TODO
}

type Handler interface {
	AppID() []int
	Receive([]byte) error
}
