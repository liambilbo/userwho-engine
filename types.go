package userwho_engine

import (
	"time"
)

type Country string
type DocumentType string
type Sex string
type MaritalStatus string
type AddressType string

type Person interface {
	GetId() string
}

type Document struct {
	Number       string
	Type         DocumentType
	IssueCountry Country
	IssueDate    time.Time
	MaturityDate time.Time
}
type Address struct {
	Country       string
	PostalCode    string
	ProvinceCode  string
	Province      string
	Town          string
	Street        string
	StreetNumber  string
	StreetType    string
	Complementary string
}

type PersonImpl struct {
	Id           string
	Name         string
	BirthDate    time.Time
	BirthCountry Country
	Documents    map[DocumentType]Document
	Addresses    map[AddressType]Address
}

type PhysicalPersonImpl struct {
	PersonImpl
	Surname           string
	SecondSurname     string
	BirthCity         string
	Nationality       Country
	SecondNationality Country
	Sex               Sex
	MaritalStatus     MaritalStatus
}

type FirmPersonImpl struct {
	PersonImpl
}
