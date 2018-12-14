package smdr

import (
	"errors"
	"strconv"
)

type conversation struct {
	year string
	month string
	day string
	hour string
	minute string
	second string
}

type CDR struct {
	length 		string
	sequence 	int
	tp 			string

	trunkOut 	string
	trunkInc 	string

	id       	string
	tenant   	string
	called   	string
	cvsStart 	conversation
	cvsEnd   	conversation

	tenantTwo    string
	condition    string
	route1       string
	route2       string
	phone        string
	callMetering string
}

// parse raw data from PBX system
func (r *CDR) parser(b []byte) error {

	r.length = string(b[2:7])
	s := string(b[9:10])
	r.sequence, _ = strconv.Atoi((s))

	if len(b) < 12 {
		return errors.New("can't parse responded data")
	}

	b = b[12:]

	r.tp = string(b[2:3])
	r.trunkOut = string(b[3:6])
	r.trunkInc = string(b[6:9])
	r.id = string(b[9:10])
	r.tenant = string(b[10:12])
	r.called = string(b[12:18])

	r.tenantTwo = string(b[48:51])
	r.condition = string(b[51:54])

	r.route1 = string(b[54:57])
	r.route2 = string(b[57:60])
	r.phone = string(b[60:71])

	r.cvsStart.year = string(b[114:116])
	r.cvsStart.month = string(b[18:20])
	r.cvsStart.day = string(b[20:22])
	r.cvsStart.hour = string(b[22:24])
	r.cvsStart.minute = string(b[24:26])
	r.cvsStart.second = string(b[26:28])

	r.cvsEnd.year = string(b[116:118])
	r.cvsEnd.month = string(b[28:30])
	r.cvsEnd.day = string(b[30:32])
	r.cvsEnd.hour = string(b[32:34])
	r.cvsEnd.minute = string(b[34:36])
	r.cvsEnd.second = string(b[36:38])

	r.callMetering = string(b[92:96])

	return nil
}