package satu

type Ss struct {
	ID      int
	Name    string
	isAdmin bool
	detail  string
}

func (s Ss) Detail() string {
	if !s.isAdmin {
		return "no access"
	}
	return s.detail
}

func NewSs(id int, name string, detail string, isAdmin bool) Ss {
	return Ss{id, name, isAdmin, detail}
}
