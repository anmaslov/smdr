// Package smdr provided for extracting SMDR data from NEC SV8100 PBX systems.
package smdr

//Used for error detection
type Parity byte

type Request struct{
	sync byte
	ident rune
	len [5]rune
	device [2]rune
	seq byte
	ack byte
	parity Parity
}

// ident 1 request for get data from PBX
func dataRequest() (Request){
	var params Request
	params.sync = 22; //Synchronization Character
	params.ident = '1' //Identifier Kind
	params.device = [2]rune{'0', '0'} //Device Number
	params.parity = 252 //Parity Byte
	params.len = [5]rune{'0', '0', '0', '0', '2'}

	return params
}

// ident 4 client response
func clientResponse(seq int) Request {
	var params Request
	params.sync = 22;
	params.ident = '4'
	params.device = [2]rune{'0', '0'}
	params.parity = 200
	params.len = [5]rune{'0', '0', '0', '0', '4'}
	params.seq = byte(seq + '0')
	params.ack = 6

	return params
}
//ident 6 connection Disconnect
func clientDisconect() Request {
	var params Request
	params.sync = 22;
	params.ident = '6'
	params.device = [2]rune{'0', '0'}
	params.parity = 252
	params.len = [5]rune{'0', '0', '0', '0', '3'}
	params.ack = 6

	return params
}

// prepare to send request to PBX system
func setRequest(p Request) []byte {
	var res []byte
	res = append(res, byte(p.sync))
	res = append(res, byte(p.ident))

	for _, value := range p.len{
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