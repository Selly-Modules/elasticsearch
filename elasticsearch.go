package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Selly-Modules/natsio"
)

// Client ...
type Client struct {
	Config        Config
	natsServer    natsio.Server
	natsJetStream natsio.JetStream
	Queue         Queue
	Pull          Pull
	Request       Request
	Push          Push
}

var (
	client *Client
)

// NewClient
// Init client elasticsearch
func NewClient(config Config) (*Client, error) {
	if config.ApiKey == "" {
		return nil, errors.New("api key is required")
	}
	if config.Nats.URL == "" {
		return nil, errors.New("nats url is required")
	}
	if err := natsio.Connect(config.Nats); err != nil {
		return nil, fmt.Errorf("nats connect failed: %v", err)
	}

	client = &Client{
		Config:        config,
		natsServer:    natsio.GetServer(),
		natsJetStream: natsio.GetJetStream(),
		Queue:         Queue{},
		Pull:          Pull{},
		Request:       Request{},
		Push:          Push{},
	}

	return client, nil
}

func GetClient() *Client {
	return client
}

// RequestNats
// publish message to nats and waiting response
func (c *Client) RequestNats(subject string, data []byte) (*Response, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   data,
		}
		res *Response
	)
	msg, err := c.natsServer.Request(subject, toBytes(req))
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(msg.Data, &res); err != nil {
		return nil, err
	}
	if res.Message != "" {
		return nil, errors.New(res.Message)
	}
	return res, nil
}

// PublishWithJetStream
// Sync data to services ES with JetStream
func (c *Client) PublishWithJetStream(streamName, subject string, data []byte) (bool, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   data,
		}
	)
	err := c.natsJetStream.Publish(streamName, subject, toBytes(req))
	if err != nil {
		return false, err
	}

	return true, nil
}

func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
