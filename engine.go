package userwho_engine

import (
	"time"
)

const (
	firmKey="firm" // Key for Firm in type Person
	physicalKey="physical" // key for Physical in type person

	addressFiscalKey="fiscal" // Key for Fiscal Address
	//addressPostalKey="postal" Key for postal Adress
)


//Init common data of a Person
func (person *Person) initPerson(personType PersonType,nationality Country,residence Country,document Document,fiscalAdress Address) {
	person.Id =document.ID()
	person.Type = personType
	person.Nationality = nationality
	person.ResidenceCountry = residence
	person.addAddress(fiscalAdress,addressFiscalKey)
	person.addDocument(document)
}

func (d *Document) ID() (string) {
	 return d.Number+"-"+string(d.Type)+"-" + string(d.IssueCountry)
}

//Create a new Physicalperson
func NewPhysicalPerson(name string, surname string, secondSurname string,nationality Country,residence Country,document Document,fiscalAdress Address) PhysicalPerson {
	person := PhysicalPerson{}
	person.initPerson(physicalKey,nationality,residence,document,fiscalAdress)
	person.Name = name
	person.Surname = surname
	person.SecondSurname = secondSurname
	return person
}

//Create a new Firm Person
func NewFirmPerson(name string,nationality Country,residence Country,document Document,fiscalAdress Address) FirmPerson {
	person := FirmPerson{}
	person.initPerson(firmKey,nationality,residence,document,fiscalAdress)
	person.Name = name
	return person
}

//Create a new address
func NewAddress(country Country, postalCode string,province string,town string,streetType StreetType,street string, streetNumber string,complementary string) Address {
	address:=Address{}
	address.Country=country
	address.PostalCode=postalCode
	address.Province=province
	address.Town=town
	address.StreetType=streetType
	address.Street=street
	address.StreetNumber=streetNumber
	address.Complementary=complementary
	return address
}


//Create a document
func NewDocument(number string, documentType DocumentType,issueCountry Country,issueDate *time.Time,maturityDate *time.Time) Document {
	document:=Document{}
	document.Number=number
	document.Type=documentType
	document.IssueCountry=issueCountry
	document.IssueDate=issueDate
	document.MaturityDate=maturityDate
	return document
}

