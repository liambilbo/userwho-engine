package userwho_engine

import (
	"time"
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
	Id() string
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
	Nationality       Country
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
