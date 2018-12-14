package smdr

import (
	"errors"
	"strconv"
)

type Conversation struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
}

type CDR struct {
	Length   		string
	Sequence 		int
	Tp       		string

	TrunkOut 		string
	TrunkInc 		string

	Id       		string
	Tenant   		string
	Called   		string
	CvsStart 		Conversation
	CvsEnd   		Conversation

	TenantTwo    	string
	Condition    	string
	Route1       	string
	Route2       	string
	Phone        	string
	CallMetering 	string
}

// parse raw data from PBX system
func (r *CDR) Parser(b []byte) error {

	r.Length = string(b[2:7])
	s := string(b[9:10])
	r.Sequence, _ = strconv.Atoi((s))

	if len(b) < 12 {
		return errors.New("can't parse responded data")
	}

	b = b[12:]

	r.Tp = string(b[2:3])
	r.TrunkOut = string(b[3:6])
	r.TrunkInc = string(b[6:9])
	r.Id = string(b[9:10])
	r.Tenant = string(b[10:12])
	r.Called = string(b[12:18])

	r.TenantTwo = string(b[48:51])
	r.Condition = string(b[51:54])

	r.Route1 = string(b[54:57])
	r.Route2 = string(b[57:60])
	r.Phone = string(b[60:71])

	r.CvsStart.Year = string(b[114:116])
	r.CvsStart.Month = string(b[18:20])
	r.CvsStart.Day = string(b[20:22])
	r.CvsStart.Hour = string(b[22:24])
	r.CvsStart.Minute = string(b[24:26])
	r.CvsStart.Second = string(b[26:28])

	r.CvsEnd.Year = string(b[116:118])
	r.CvsEnd.Month = string(b[28:30])
	r.CvsEnd.Day = string(b[30:32])
	r.CvsEnd.Hour = string(b[32:34])
	r.CvsEnd.Minute = string(b[34:36])
	r.CvsEnd.Second = string(b[36:38])

	r.CallMetering = string(b[92:96])

	return nil
}