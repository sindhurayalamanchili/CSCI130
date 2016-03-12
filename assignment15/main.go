
package main

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

func main() {

	// The handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// Generating a new ID
		id, err := uuid.NewV4()

		// Logging possible errors generating the ID
		logError(err)

		// Setting the cookie
		cookie := &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}
		// Setting the cookie on the response back to the client
		http.SetCookie(res, cookie)
	})

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
