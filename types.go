package userwho_engine

import (
	"time"
	"fmt"
)

type Country string
type DocumentType string
type Sex string
type MaritalStatus string
type AddressType string
type PersonType string
type StreetType string
type Cnae string
type EducationLevel string
type ProfessionalActivity string
type LaboralSituation string


type Actor interface {
	FullName() string
}

type Document struct {
	Number       string
	Type         DocumentType
	IssueCountry Country
	IssueDate    time.Time
	MaturityDate time.Time
}
type Address struct {
	Country       Country
	PostalCode    string
	Province      string
	Town          string
	StreetType    StreetType
	Street        string
	StreetNumber  string
	Complementary string
}

type PersonAddress struct {
	Type AddressType
	Order int
	Address
}

type Documents struct {
	Documents    map[DocumentType]Document
}


type Addresses struct {
	Addresses    []PersonAddress
}

type Person struct {
	Id                  string
	Type                PersonType
	Nationality         Country
	ResidenceCountry    Country
	Documents
	Addresses
}

type PhysicalPerson struct {
	Person
	Name              string
	Surname           string
	SecondSurname     string
	BirthCity         string
	BirthDate         time.Time
	BirthCountry       Country
	SecondNationality Country
	Sex               Sex
	MaritalStatus     MaritalStatus
	EducationLevel    EducationLevel
	ProfessionalActivity ProfessionalActivity
	LaboralSituation LaboralSituation
}

type FirmPerson struct {
	Person
	Name                  string
	SettingUpDate         time.Time
	SettingUpCountry      Country
	Cnae Cnae
}

func (person PhysicalPerson) FullName() string {
	return fmt.Sprintf("%s %s , %s" ,person.SecondSurname,person.Surname,person.Name)
}

func (person FirmPerson) FullName() string {
	return person.Name
}
