package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//HTTPGetStringResponse makes a GET request to a given URL using provided headers
func HTTPGetStringResponse(url string, headers map[string]string) string {

	res, err := http.Get(url)
	if err != nil {
		FailOnError(err, fmt.Sprint("a problem occured fetching GET %s", url))
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return BodyBytesToString(res.Body)
}

//HTTPGet makes a GET request to a given URL using provided headers
func HTTPGet(url string, headers map[string]string) ([]byte, error) {

	var err error
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	FailOnError(err, "could execute request")

	//add headers to your request
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	//execute the request
	res, err := client.Do(req)
	FailOnError(err, "a problem, occured")

	if res.StatusCode != 200 {
		err := fmt.Errorf("a problem, occured: %d on %s", res.StatusCode, url)
		FailOnError(err, "")
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)

	return bodyBytes, err
}

//HTTPPost makes a POST request to a given URL using provided headers
func HTTPPost(url string, payload string, headers map[string]string) ([]byte, error) {

	var err error
	client := &http.Client{}
	content := strings.NewReader(payload)
	req, err := http.NewRequest("POST", url, content)
	FailOnError(err, "could execute request")

	//add headers to your request
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	//execute the request
	res, err := client.Do(req)
	FailOnError(err, "a problem, occured")

	if res.StatusCode < 200 || res.StatusCode > 299 {
		err := fmt.Errorf("a problem, occured: %d on %s", res.StatusCode, url)
		FailOnError(err, "")
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)

	return bodyBytes, err
}
