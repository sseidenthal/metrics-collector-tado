package tado

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sseidenthal/tado-collector/helpers"
)

//Login into tado api
func Login(username string, password string, client_id string, client_secret string) Token {

	apiURL := "https://auth.tado.com/oauth/token"

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("scope", "home.user")
	data.Set("client_id", client_id)
	data.Set("client_secret", client_secret)
	data.Set("password", password)
	data.Set("username", username)

	u, _ := url.ParseRequestURI(apiURL)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	helpers.FailOnError(err, "a problem, occured")

	if resp.StatusCode != 200 {
		err := fmt.Errorf("a problem, occured: %d on %s", resp.StatusCode, apiURL)
		helpers.FailOnError(err, "")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	log.Println("login: success")

	var t Token
	json.Unmarshal(bodyBytes, &t)

	return t
}

//Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}
