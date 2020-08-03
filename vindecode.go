package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type vinResponse struct {
	Count          int         `json:"Count"`
	Meaasge        string      `json:"Message"`
	SearchCriteria string      `json:"SearchCriteria"`
	Attributes     []Attribute `json:"Results"`
}

// Attribute for a vehicle
type Attribute struct {
	Value      string `json:"Value"`
	ValueID    string `json:"ValueId"`
	Variable   string `json:"Variable"`
	VariableID int    `json:"VariableId"`
}

func findStringAttribute(attributes []Attribute, attributeName string) string {
	for i := 0; i < len(attributes); i++ {
		if attributes[i].Variable == attributeName {
			return string(attributes[i].Value)
		}
	}

	return ""
}

func main() {
	vin := os.Args[1]
	url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/decodevin/%s?format=json", vin)
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var response vinResponse
	json.Unmarshal(body, &response)

	var year        = findStringAttribute(response.Attributes, "Model Year")
	var make        = findStringAttribute(response.Attributes, "Make")
	var model       = findStringAttribute(response.Attributes, "Model")
	var trim        = findStringAttribute(response.Attributes, "Trim")
	var vehicleType = findStringAttribute(response.Attributes, "Vehicle Type")
	var driveType   = findStringAttribute(response.Attributes, "Drive Type")

	log.Println(year)
	log.Println(make)
	log.Println(model)
	log.Println(trim)
	log.Println(vehicleType)
	log.Println(driveType)
}
