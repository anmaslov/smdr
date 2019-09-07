package smdr

import (
	"reflect"
	"testing"
)

var requestData = &Request{
	sync:   22,
	ident:  '1',
	device: [2]rune{'0', '0'},
	parity: 252,
	len:    [5]rune{'0', '0', '0', '0', '2'},
}

// TestDataRequest test request data
func TestDataRequest(t *testing.T) {
	data := DataRequest()

	if !reflect.DeepEqual(requestData, data) {
		t.Error("Expected data, got", data)
	}
}

// TestClientResponse test client response data
func TestClientResponse(t *testing.T) {
	testData := &Request{
		sync:   22,
		ident:  '4',
		len:    [5]rune{'0', '0', '0', '0', '4'},
		device: [2]rune{'0', '0'},
		seq:    byte(5 + '0'),
		ack:    6,
		parity: 200,
	}
	data := ClientResponse(5)

	if !reflect.DeepEqual(testData, data) {
		t.Error("Expected data, got", data)
	}
}

// TestClientDisconnect test client disconnect
func TestClientDisconnect(t *testing.T) {
	testData := &Request{
		sync:   22,
		ident:  '6',
		len:    [5]rune{'0', '0', '0', '0', '3'},
		device: [2]rune{'0', '0'},
		ack:    6,
		parity: 252,
	}

	data := ClientDisconnect()

	if !reflect.DeepEqual(testData, data) {
		t.Error("Expected data, got", data)
	}
}

// TestSetRequestData test set request data
func TestSetRequestData(t *testing.T) {
	test := []byte{22, 49, 48, 48, 48, 48, 50, 48, 48, 252}
	req := SetRequest(requestData)

	if len(test) != len(req) {
		t.Error("Len of slice is expected", req)
	}
}
