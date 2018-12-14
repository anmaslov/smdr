package smdr

import "testing"

var tests = Request{sync: 22, ident: '1',
	device: [2]rune{'0', '0'}, parity: 252,
	len: [5]rune{'0', '0', '0', '0', '2'}}

func TestData(t *testing.T) {

	data := dataRequest()

	if tests != data {
		t.Error("Expected data, got", data)
	}
}

func TestSetRequestData(t *testing.T)  {
	test := []byte{22, 49, 48, 48, 48, 48, 50, 48, 48, 252}
	req := setRequest(tests)

	if len(test) != len(req) {
		t.Error("Len of slice is expected", req)
	}
}
