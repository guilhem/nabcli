package client

//{"status": "ok", "type": "response"}

type Packet struct {
	Type      string `json:"type"`
	RequestID string `json:"request_id,omitempty"`
}

type Wakeup struct {
	Packet
}

func NewWakeup() Packet {
	return Packet{
		Type: "wakeup",
	}
}

func NewSleep() Packet {
	return Packet{
		Type: "sleep",
	}
}

type Response struct {
	Packet
	Status  string `json:"status,omitempty"`
	Class   string `json:"class,omitempty"`
	Message string `json:"message,omitempty"`
}
