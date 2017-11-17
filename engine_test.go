package userwho_engine

import (
	"testing"
)

func TestNewFirmPerson(t *testing.T) {
	address:=NewAddress("ESP","28027","Madrid","Madrid","CL","Rocadragon","4","10-F")
    document:=NewDocument("0123456","CIF","ESP",nil,nil)
	firmPerson:=NewFirmPerson("Lorea S.A.","ESP","ESP",document,address)

	addresses:=firmPerson.Address(addressFiscalKey)

	if addresses==nil || len(addresses)==0 {
		t.Errorf("Adding a fiscal address , should return a fiscal address")
	} else if addresses[0].Street!="Rocadragon"{
		t.Errorf("Adding a fiscal address , should return a fiscal address with the same street")
	}

	passport:=firmPerson.Document("CIF")
	if passport.Number != "0123456" {
		t.Errorf("CIF is not equal")
	}

}


func TestNewPhysicalPerson(t *testing.T) {
	address:=NewAddress("ESP","28027","Madrid","Madrid","CL","Rocadragon","4","10-F")
	document:=NewDocument("K12312","passport","ESP",nil,nil)
	physicalPerson:=NewPhysicalPerson("Juan","Rodrigo","","ESP","ESP",document,address)

	addresses:=physicalPerson.Address(addressFiscalKey)

	if addresses==nil || len(addresses)==0 {
		t.Errorf("Adding a fiscal address , should return a fiscal address")
	} else if addresses[0].Street!="Rocadragon"{
		t.Errorf("Adding a fiscal address , should return a fiscal address with the same street")
	}

	passport:=physicalPerson.Document("passport")
	if passport.Number != "K12312" {
		t.Errorf("Passport is not equal")
	}

	fullname:=physicalPerson.FullName()
	if fullname != "Rodrigo , Juan" {
		t.Errorf("FullName is not equal")
	}

}


