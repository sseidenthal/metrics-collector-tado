package helpers

import (
	"io"
	"io/ioutil"
	"log"
)

//BodyBytesToString converts http body bytes to string
func BodyBytesToString(body io.Reader) string {

	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyBytes)
}

//FailOnError handles errors in a more compact way
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
