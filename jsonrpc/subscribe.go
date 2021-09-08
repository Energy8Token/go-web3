package jsonrpc

import (
	"fmt"

	"github.com/Energy8Token/go-web3/jsonrpc/transport"
)

// SubscriptionEnabled returns true if the subscription endpoints are enabled
func (c *Client) SubscriptionEnabled() bool {
	_, ok := c.transport.(transport.PubSubTransport)
	return ok
}

// Subscribe starts a new subscription
func (c *Client) Subscribe(callback func(b []byte), params ...interface{}) (func() error, error) {
	pub, ok := c.transport.(transport.PubSubTransport)
	if !ok {
		return nil, fmt.Errorf("Transport does not support the subscribe method")
	}
	close, err := pub.Subscribe(callback, params...)
	return close, err
}
