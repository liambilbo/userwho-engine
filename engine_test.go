package userwho_engine

import "testing"

func TestNewFirmPerson(t *testing.T) {
	address:=NewAddress("ESP","28027","Madrid","Madrid","CL","Rocadragon","4","10-F")

	firmPerson:=NewFirmPerson("Lorea S.A.","ESP","ESP")
    firmPerson.addAddress(address,addressFiscalKey)

	addresses:=firmPerson.Address(addressFiscalKey)

	if addresses==nil || len(addresses)==0 {
		t.Errorf("Adding a fiscal address , should return a fiscal address")
	} else if addresses[0].Street!="Rocadragon"{
		t.Errorf("Adding a fiscal address , should return a fiscal address with the same street")
	}
}
