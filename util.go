package nadeshiko

import (
	"log"
	"net/http"
)

//StartServer is just probably most copy and pasted bit of code that I finally decided to extract
func StartServer() {
	address := ":3000"
	log.Printf("Starting server on %s", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Panic(err)
	}
}
