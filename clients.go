package main

import (
	"fmt"
	"time"

	funcs "github.com/denizedizcan/week-5-homework-1-denizedizcan/funcs"
)

func main() {
	//clients
	nginxServer := funcs.NewNginxServer()
	getLicenceURL := "/licence"
	logutLicenceURL := "/licence/loguot"

	for i := 0; i < 15; i++ {
		// get licence from server concurrently
		go func() {
			httpCode, body := nginxServer.HandleRequest(getLicenceURL, "GET")
			fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", getLicenceURL, httpCode, body)
		}()
	}
	for i := 0; i < 3; i++ {
		go func() {
			// gave back licence concurrently
			httpCode, body := nginxServer.HandleRequest(logutLicenceURL, "POST")
			fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", logutLicenceURL, httpCode, body)
		}()
	}
	defer fmt.Printf("%T & %#v", nginxServer, nginxServer)
	time.Sleep(1 * time.Second)

}
