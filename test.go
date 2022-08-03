package main

type User interface {
	Create(name string) error
}

type user struct{}

func (self *user) Create(name string) error {
	panic("not implemented")
}
