package transport

import (
	"os"
	"strings"
)

// Transport is an inteface for transport methods to send jsonrpc requests
type Transport interface {
	// Call makes a jsonrpc request
	Call(method string, out interface{}, params ...interface{}) error

	// Close closes the transport connection if necessary
	Close() error
}

// PubSubTransport is a transport that allows subscriptions
type PubSubTransport interface {
	// Subscribe starts a subscription to a new event
	Subscribe(callback func(b []byte), params ...interface{}) (func() error, error)
}

const (
	wsPrefix  = "ws://"
	wssPrefix = "wss://"
)

// NewTransport creates a new transport object
func NewTransport(url string) (Transport, error) {
	if strings.HasPrefix(url, wsPrefix) || strings.HasPrefix(url, wssPrefix) {
		t, err := newWebsocket(url)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	if _, err := os.Stat(url); err == nil {
		// path exists, it could be an ipc path
		t, err := newIPC(url)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return newHTTP(url), nil
}
