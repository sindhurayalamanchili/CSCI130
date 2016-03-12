
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// Registering the URL path and binding it to surferPage handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Name: "+req.FormValue("n"))
	})

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}
