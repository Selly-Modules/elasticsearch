package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Selly-Modules/natsio"
)

type Client struct {
	Config        Config
	natsServer    natsio.Server
	natsJetStream natsio.JetStream
}

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

	c := &Client{
		Config:        config,
		natsServer:    natsio.GetServer(),
		natsJetStream: natsio.GetJetStream(),
	}

	return c, nil
}

func (c *Client) SyncData(data SyncData) (bool, error) {
	var (
		res Response
	)
	msg, err := c.natsServer.Request(SubjectSyncData, toBytes(data))
	if err != nil {
		return false, err
	}
	if err = json.Unmarshal(msg.Data, &res); err != nil {
		return false, err
	}
	if res.Message != "" {
		return false, errors.New(res.Message)
	}
	return res.Success, nil
}

func (c *Client) Search(query ESQuery) ([]string, error) {
	var (
		res Response
	)
	msg, err := c.natsServer.Request(SubjectSearch, toBytes(query))
	if err != nil {
		return res.Data, err
	}
	if err = json.Unmarshal(msg.Data, &res); err != nil {
		return res.Data, err
	}
	if res.Message != "" {
		return res.Data, errors.New(res.Message)
	}
	return res.Data, nil
}

func (c *Client) UpdateDocument(query UpdateDataPayload) ([]string, error) {
	var (
		res Response
	)
	msg, err := c.natsServer.Request(SubjectUpdateDocument, toBytes(query))
	if err != nil {
		return res.Data, err
	}
	if err = json.Unmarshal(msg.Data, &res); err != nil {
		return res.Data, err
	}
	if res.Message != "" {
		return res.Data, errors.New(res.Message)
	}
	return res.Data, nil
}

func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
