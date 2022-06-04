package db

type API interface {
	GetUser() []int
}

type Login interface {
	Signing(username string, password string) bool
}
