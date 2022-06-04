package service

import "belajar/dua/db"

// interface
type Service struct {
	api   db.API
	login db.Login
}

func New(a db.API, b db.Login) Service {
	return Service{
		api:   a,
		login: b,
	}
}

func (s Service) Login(param1 string, param2 string) {
	s.login.Signing(param1, param2)
}

func (s Service) GetUser() {
	s.api.GetUser()
}
