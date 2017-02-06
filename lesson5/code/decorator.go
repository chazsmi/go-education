package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chazsmi/decorator-pattern/client"
)

func main() {
	// START OMIT
	// Add the default http client and call pass in our Auth method to add the token
	cli := client.Decorate(
		http.DefaultClient,
		client.Log(log.New(os.Stdout, "client: ", log.LstdFlags)),
	)

	req, err := http.NewRequest("GET", "http://www.sainsburys.co.uk", nil)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	resp, err := cli.Do(req)
	// END OMIT
}
