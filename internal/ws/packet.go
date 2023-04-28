package ws

type Packet struct {
	Client  *Client
	Message []byte
}
