// Package smdr provided for extracting SMDR data from NEC SV8100 PBX systems.
package smdr

// Parity used for error detection
type Parity byte

// Request to pbx
type Request struct {
	// Synchronization Character
	sync byte
	// Identifier Kind
	ident rune
	len   [5]rune
	// Device Number
	device [2]rune
	seq    byte
	ack    byte
	// Parity Byte
	parity Parity
}

// DataRequest ident 1 request for get data from PBX
func DataRequest() *Request {
	return &Request{
		sync:   22,
		ident:  1,
		len:    [5]rune{'0', '0', '0', '0', '2'},
		device: [2]rune{'0', '0'},
		parity: 252,
	}
}

// ClientResponse ident 4 client response
func ClientResponse(seq int) *Request {
	return &Request{
		sync:   22,
		ident:  '4',
		len:    [5]rune{'0', '0', '0', '0', '4'},
		device: [2]rune{'0', '0'},
		seq:    byte(seq + '0'),
		ack:    6,
		parity: 200,
	}
}

// ClientDisconnect ident 6 connection Disconnect
func ClientDisconnect() *Request {
	return &Request{
		sync:   22,
		ident:  '6',
		len:    [5]rune{'0', '0', '0', '0', '3'},
		device: [2]rune{'0', '0'},
		ack:    6,
		parity: 252,
	}
}

// SetRequest prepare to send request to PBX system
func SetRequest(p Request) []byte {
	var res []byte
	res = append(res, byte(p.sync))
	res = append(res, byte(p.ident))

	for _, value := range p.len {
		res = append(res, byte(value))
	}

	res = append(res, byte(p.device[0]))
	res = append(res, byte(p.device[1]))

	if p.seq != 0 {
		res = append(res, byte(p.seq))
	}

	if p.ack != 0 {
		res = append(res, p.ack)
	}

	res = append(res, byte(p.parity))

	return res
}
