package smdr

import (
	"fmt"
	"strconv"
	"strings"
)

// Conversation datetime structure
type Conversation struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
}

// CDR structure for response from pbx
type CDR struct {
	Length   string
	Sequence int
	Tp       string

	TrunkOut string
	TrunkInc string

	Id       string
	Tenant   string
	Called   string
	CvsStart Conversation
	CvsEnd   Conversation

	TenantTwo    string
	Condition    string
	Route1       string
	Route2       string
	Phone        string
	CallMetering string
}

// New create new blank cdr structure
func New() *CDR {
	return &CDR{}
}

// Parser parse raw data from PBX system
func (r *CDR) Parser(b []byte) error {
	if len(b) < 12 {
		return fmt.Errorf("can't parse responded data")
	}

	r.Length = string(b[2:7])
	s := string(b[9:10])
	r.Sequence, _ = strconv.Atoi(s)

	b = b[12:]

	if len(b) < 118 {
		return fmt.Errorf("can't parse responded call data")
	}

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
	r.Phone = string(b[60:92])

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

	//Extended NEAX 2400 IMS Format
	if len(b) > 163 && len(strings.Trim(r.Phone, " ")) == 0 {
		r.Phone = string(b[131:163])
	}

	return nil
}
