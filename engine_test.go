package userwho_engine

import (
	"testing"
	"time"
)

func TestNewFirmPerson(t *testing.T) {
	address:=NewAddress("ESP","28027","Madrid","Madrid","CL","Rocadragon","4","10-F")
    document:=NewDocument("PK34","passport","ESP",time.Time{},time.Time{})
	firmPerson:=NewFirmPerson("Lorea S.A.","ESP","ESP",document,address)

	addresses:=firmPerson.Address(addressFiscalKey)

	if addresses==nil || len(addresses)==0 {
		t.Errorf("Adding a fiscal address , should return a fiscal address")
	} else if addresses[0].Street!="Rocadragon"{
		t.Errorf("Adding a fiscal address , should return a fiscal address with the same street")
	}

	passport:=firmPerson.Document("passport")
	if passport.Number != "PK34" {
		t.Errorf("Passport is not equal")

	}

}
