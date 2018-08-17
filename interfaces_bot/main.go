package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type swBot struct{}
type enBot struct{}

type Bot interface {
	getGreeting() string
}

func (en enBot) getGreeting() string {
	return "Hello"
}

func (sw swBot) getGreeting() string {
	return "Habari"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
func main() {
	eb := enBot{}
	sb := swBot{}
	printGreeting(eb)
	printGreeting(sb)

	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(resp)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Println(string(b))

}

/* explanation for http.Get and the response
##http.Get
func Get(url string) (resp *Response, err error)

## Response

type Response struct {
        Status     string // e.g. "200 OK"
        StatusCode int    // e.g. 200
        Proto      string // e.g. "HTTP/1.0"
        ProtoMajor int    // e.g. 1
        ProtoMinor int    // e.g. 0
        Header Header
		Body io.ReadCloser
		ContentLength int64
		........
}

## Body if of type ReadCloser which is an interface
type ReadCloser interface {
        Reader
        Closer
}

type Reader interface {
        Read(p []byte) (n int, err error)
}

type Closer interface {
        Close() error
}

#hence we can read the value of  the body using ioutil.ReadAll


func ReadAll(r io.Reader) ([]byte, error)
*/
