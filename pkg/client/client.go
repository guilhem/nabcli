package client

import (
	"encoding/json"
	"fmt"
	"net"
)

// Client is a object able to talk a remote footloose API server.
type Client struct {
	Decoder *json.Decoder
	Encoder *json.Encoder
}

func New(conn net.Conn) *Client {
	decode := json.NewDecoder(conn)
	encode := json.NewEncoder(conn)
	return &Client{Decoder: decode, Encoder: encode}
}

func (c *Client) Send(v interface{}) error {
	err := c.Encoder.Encode(v)
	if err != nil {
		return err
	}

	var r Response
	err = c.Decoder.Decode(&r)

	if err != nil {
		return err
	}

	if r.Status == "error" {
		return fmt.Errorf("Server return error")
	}

	return nil
}

func (c *Client) Wakeup() error {
	return c.Send(NewWakeup())
}

func (c *Client) Sleep() error {
	return c.Send(NewSleep())
}
