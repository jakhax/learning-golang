package main

import (
	"fmt"
	"net/http"
)

//this is not a real site checker many factors could be wrong such as a network issue
// just using this to understand go routines and channels

var links = []string{
	"http://google.com",
	"http://facebook.com",
	"http://golang.org",
	"http://stackoverflow.com",
}

type Status struct {
	state bool
	link  string
}

func linkCheck(l string, c chan Status) {
	_, err := http.Get(l)
	if err != nil {
		c <- Status{state: false, link: l}
	}
	c <- Status{state: true, link: l}
}

func main() {
	c := make(chan Status)
	for _, l := range links {
		go linkCheck(l, c)
	}

	for i := 0; i < len(links); i++ {
		s := <-c
		if s.state == true {
			fmt.Printf("Success: %s seems up\n", s.link)
		} else {
			fmt.Printf("Error: %s seems down\n", s.link)
		}

	}
}
