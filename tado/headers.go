package tado

import "fmt"

//GenerateHeaders generate headers required by tado's API
func GenerateHeaders(token string) map[string]string {

	headers := make(map[string]string)
	headers["Referer"] = "https://my.tado.com/app/de/main/home"
	headers["DNT"] = "1"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"

	return headers
}
