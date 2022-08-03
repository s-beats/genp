package main

import "time"

type User interface {
	Create(name string, endedAt time.Time) error
}

type user struct{}

func (self *user) Create(name string, endedAt time.Time) error {
	panic("not implemented")
}
