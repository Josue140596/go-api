package user

import "log"

// This can used by others packages
type Service interface{
	Create(firstName, lastName, email, password ,phone string) error
}

//This is private you can use it in this current package
type service struct {

}

// It'll return a Service  
func NewService() Service{
	return &service{}
}


func (s service) Create(firstName, lastName, email, password ,phone string) error{
	log.Print("Aqui desde el método")
	return nil
}