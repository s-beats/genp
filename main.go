package main

import "io/ioutil"

func main() {
	ioutil.WriteFile("test.go", []byte("This is test"), 0644)
}
