package smdr

import (
	"reflect"
	"testing"
)

// Failed parse result test
func TestParserFailed(t *testing.T) {
	b := []byte{22, 49, 48, 48, 48, 48}
	p := New()
	err := p.Parser(b)
	if err != nil && err.Error() != `can't parse responded data` {
		t.Errorf("can't parse responded data: %#v", err)
	}

	b = []byte{22, 49, 48, 48, 48, 48, 50, 48, 48, 1, 1, 252, 12}
	p = New()
	err = p.Parser(b)
	if err != nil && err.Error() != `can't parse responded call data` {
		t.Errorf("can't parse responded data: %#v", err)
	}
}

// Parse result test
func TestParser(t *testing.T) {
	//some test data
	b := []byte{22, 49, 48, 48, 48, 48, 50, 48, 48, 1,
		1, 252, 1, 2, 55, 55, 55, 65, 55, 55,
		50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
		51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
		52, 52, 52, 52, 52, 52, 52, 52, 52, 52,
		53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
		54, 54, 54, 54, 54, 54, 54, 54, 54, 54,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 47, 47, 47, 47, 47, 47,
		59, 59, 59, 59, 59, 59, 59, 59, 59, 59,
		51, 60, 60, 60, 60, 60, 49, 48, 54, 53,
		52, 59, 59, 59, 59, 59, 59, 59, 59, 59,
		53, 60, 60, 55, 57, 48, 57, 53, 49, 57,
		52, 55, 53, 54, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
	p := New()
	_ = p.Parser(b)

	testData := &CDR{
		Length:   "00002",
		Sequence: 0,
		Tp:       "7",
		TrunkOut: "77A",
		TrunkInc: "772",
		Id:       "2",
		Tenant:   "22",
		Called:   "222222",
		CvsStart: Conversation{
			Year:   "10",
			Month:  "33",
			Day:    "33",
			Hour:   "33",
			Minute: "33",
			Second: "33",
		},
		CvsEnd: Conversation{
			Year:   "65",
			Month:  "44",
			Day:    "44",
			Hour:   "44",
			Minute: "44",
			Second: "44",
		},
		TenantTwo:    "666",
		Condition:    "666",
		Route1:       "666",
		Route2:       "6  ",
		Phone:        "79095194756                     ",
		CallMetering: "////",
	}

	if !reflect.DeepEqual(testData, p) {
		t.Error("Expected data, got", p)
	}
}
