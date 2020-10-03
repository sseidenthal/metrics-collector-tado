package tado

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sseidenthal/tado-collector/helpers"
)

//Zones represent a collection of Zone
type Zones []Zone

//Zone is Tado Zone
type Zone struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//State of a Zone, contains values temperature and humidity
type State struct {
	Zone    Zone
	Setting struct {
		Temperature struct {
			Celsius float64 `json:"celsius"`
		} `json:"temperature"`
	} `json:"setting"`
	SensorDataPoints struct {
		InsideTemperature struct {
			Celsius float64 `json:"celsius"`
		} `json:"insideTemperature"`
		Humidity struct {
			Percentage float64 `json:"percentage"`
		} `json:"humidity"`
	} `json:"sensorDataPoints"`
}

//GetCurrentTemperature ...
func (s *State) GetCurrentTemperature() float64 {
	return s.SensorDataPoints.InsideTemperature.Celsius
}

//GetCurrentHumidity ...
func (s *State) GetCurrentHumidity() float64 {
	return s.SensorDataPoints.Humidity.Percentage
}

//GetRequestedTemperature ...
func (s *State) GetRequestedTemperature() float64 {
	return s.Setting.Temperature.Celsius
}

//GetZoneName ...
func (s *State) GetZoneName() string {
	return s.Zone.Name
}

//GetZoneID ...
func (s *State) GetZoneID() int {
	return s.Zone.ID
}

//GetZones from tado's api
func GetZones(token string, home int) Zones {

	url := fmt.Sprintf("https://my.tado.com/api/v2/homes/%d/zones", home)
	r, err := helpers.HTTPGet(url, GenerateHeaders(token))

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}

	var zones Zones
	json.Unmarshal(r, &zones)

	log.Println("fetch zones: success")
	return zones
}

//GetZonesState ...
func GetZonesState(token string, home int, zones Zones) map[int]State {

	states := make(map[int]State)

	for _, z := range zones {
		url := fmt.Sprintf("https://my.tado.com/api/v2/homes/%d/zones/%d/state", home, z.ID)
		r, err := helpers.HTTPGet(url, GenerateHeaders(token))

		if err != nil {
			fmt.Println(err)
			log.Fatal(err.Error())
		}

		var state State
		json.Unmarshal(r, &state)
		state.Zone = z

		states[state.Zone.ID] = state
	}

	log.Println("fetch states: success")

	return states
}
