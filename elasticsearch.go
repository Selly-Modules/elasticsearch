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
}

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

	c := &Client{
		Config:        config,
		natsServer:    natsio.GetServer(),
		natsJetStream: natsio.GetJetStream(),
	}

	return c, nil
}

// SyncData
// Sync data to services ES
func (c *Client) SyncData(data SyncData) (bool, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   toBytes(data),
		}
		res *Response
	)
	msg, err := c.natsServer.Request(SubjectSyncData, toBytes(req))
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

// SyncDataWithJetStream
// Sync data to services ES with JetStream
func (c *Client) SyncDataWithJetStream(data SyncData) (bool, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   toBytes(data),
		}
	)
	err := c.natsJetStream.Publish(JetStreamSearchService, SubjectSyncData, toBytes(req))
	if err != nil {
		return false, err
	}

	return true, nil
}

// Search
// Request search to service es
func (c *Client) Search(query ESQuery) (*Response, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   toBytes(query),
		}
		res *Response
	)
	msg, err := c.natsServer.Request(SubjectSearch, toBytes(req))
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

// UpdateDocument
// Insert or update document to ES
func (c *Client) UpdateDocument(query UpdateDataPayload) (bool, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   toBytes(query),
		}
	)
	err := c.natsJetStream.Publish(JetStreamSearchService, SubjectUpdateDocument, toBytes(req))
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteDocument
// Delete document to ES
func (c *Client) DeleteDocument(payload DeleteDataPayload) (bool, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   toBytes(payload),
		}
	)
	err := c.natsJetStream.Publish(JetStreamSearchService, SubjectUpdateDocument, toBytes(req))
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateIndex
// Create index to ES
func (c *Client) CreateIndex(name string) (bool, error) {
	var (
		req = RequestBody{
			ApiKey: c.Config.ApiKey,
			Body:   toBytes(name),
		}
		res *Response
	)
	msg, err := c.natsServer.Request(SubjectCreateIndex, toBytes(req))
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

func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
