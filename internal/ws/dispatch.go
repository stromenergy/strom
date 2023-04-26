package ws

type Dispatcher interface {
	Broadcast(message []byte)
	Register(client *Client)
	Unregister(client *Client)
}