package userwho_engine

import "github.com/pborman/uuid"

func (person *PersonImpl) GetId() string {
	return person.Id
}
func (person *PhysicalPersonImpl) GetId() string {
	return person.Id
}
func (person *FirmPersonImpl) GetId() string {
	return person.Id
}

func initPerson(person *PersonImpl, name string) {
	person.Id = uuid.New()
	person.Name = name
	return
}

func NewPhysicalPerson(name string, surname string, secondSurname string) *PhysicalPersonImpl {
	person := &PhysicalPersonImpl{}
	person.Name = name
	person.Surname = surname
	person.SecondSurname = secondSurname
	return person
}

func NewFirmPerson(name string) *FirmPersonImpl {
	person := &FirmPersonImpl{}
	initPerson(&person.PersonImpl, name)
	return person
}
