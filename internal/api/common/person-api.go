package common

// API
type Person interface {
	Update(person *PersonModel) error
	Model() *PersonModel
}

// Builders

// Model

type PersonModel struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (model *PersonModel) Change(newModel *PersonModel) {
	model.FirstName = newModel.FirstName
	model.LastName = newModel.LastName
	model.Email = newModel.Email
	model.Phone = newModel.Phone
}

// Solid

type SolidPerson struct {
	model  *PersonModel
	person Person
}

func NewSolidPerson(model *PersonModel, person Person) SolidPerson {
	return SolidPerson{model, person}
}

func (p SolidPerson) Update(newModel *PersonModel) error {
	p.model.Change(newModel)
	if p.person == nil {
		return nil
	}
	return p.person.Update(newModel)
}

func (p SolidPerson) Model() *PersonModel {
	return p.model
}
