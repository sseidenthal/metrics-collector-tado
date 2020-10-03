package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tkanos/gonfig"

	"github.com/sseidenthal/tado-collector/helpers"
	"github.com/sseidenthal/tado-collector/tado"
)

//Config hold configuration read from config.json
var Config Configuration

func start() {

	token := tado.Login(Config.TadoUsername, Config.TadoPassword, Config.TadoClientID, Config.TadoClientSecret)
	home := extractHomeID(token.AccessToken)
	zones := tado.GetZones(token.AccessToken, home)
	states := tado.GetZonesState(token.AccessToken, home, zones)

	var headers map[string]string

	for _, state := range states {

		payload1 := fmt.Sprintf("temperature_current,zone_id=%d,zone_name=%s value=%f", state.GetZoneID(), state.GetZoneName(), state.GetCurrentTemperature())
		helpers.HTTPPost(Config.InfluxdbDsn, payload1, headers)
		log.Println("deliver temperature_current: success")

		payload2 := fmt.Sprintf("temperature_requested,zone_id=%d,zone_name=%s value=%f", state.GetZoneID(), state.GetZoneName(), state.GetRequestedTemperature())
		helpers.HTTPPost(Config.InfluxdbDsn, payload2, headers)
		log.Println("deliver temperature_requested: success")

		payload3 := fmt.Sprintf("humidity,zone_id=%d,zone_name=%s value=%f", state.GetZoneID(), state.GetZoneName(), state.GetCurrentTemperature())
		helpers.HTTPPost(Config.InfluxdbDsn, payload3, headers)
		log.Println("deliver humidity: success")

	}

}

func main() {

	Config = GetConfig()

	start()

	ticker := time.NewTicker(Config.Interval * time.Second)

	for range ticker.C {
		start()
	}
}

//this function is a massacre ..
//i could not yet find a clean way to extract custom claim tado_homes from the JTW
//feedback is welcome
func extractHomeID(tokenString string) int {
	home := 0
	if token, _ := jwt.Parse(tokenString, nil); token != nil {
		claims := token.Claims.(jwt.MapClaims)
		home = int(claims["tado_homes"].([]interface{})[0].(map[string]interface{})["id"].(float64))

	}
	return home
}

//GetConfig read config.json
func GetConfig(params ...string) Configuration {
	config := Configuration{}
	fileName := "./config.json"
	gonfig.GetConf(fileName, &config)
	return config
}

//Configuration ..
type Configuration struct {
	InfluxdbDsn      string        `json:"influxdb_dsn"`
	Interval         time.Duration `json:"interval"`
	TadoClientID     string        `json:"tado_client_id"`
	TadoClientSecret string        `json:"tado_client_secret"`
	TadoPassword     string        `json:"tado_password"`
	TadoUsername     string        `json:"tado_username"`
}

type JWT struct {
	Sub       string `json:"sub"`
	TadoHomes []struct {
		ID int `json:"id"`
	} `json:"tado_homes"`
	Iss          string   `json:"iss"`
	Locale       string   `json:"locale"`
	Aud          []string `json:"aud"`
	Nbf          int      `json:"nbf"`
	TadoScope    []string `json:"tado_scope"`
	TadoUsername string   `json:"tado_username"`
	Name         string   `json:"name"`
	Exp          int      `json:"exp"`
	Iat          int      `json:"iat"`
	TadoClientID string   `json:"tado_client_id"`
	Jti          string   `json:"jti"`
	Email        string   `json:"email"`
}
