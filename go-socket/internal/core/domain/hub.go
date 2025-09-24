package domain

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateRoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
