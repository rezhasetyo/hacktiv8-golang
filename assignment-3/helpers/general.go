package helpers

import (
	models "assignment3/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

func GenerateRandomNumberRange() int {
	min := 0
	max := 20
	randomNumb := rand.Intn(max-min) + min
	return randomNumb
}

func GetStatus(w models.Weather) string {

	if w.Water < 5 && w.Wind < 6 {
		return "Aman"
	} else if (w.Water >= 6 && w.Water <= 8) || (w.Wind >= 7 && w.Wind <= 15) {
		return "Siaga"
	} else if (w.Water > 8) || (w.Wind > 15) {
		return "Bahaya"
	}

	return "Failed"
}

func WriteFile(data interface{}, pathName string) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(pathName, file, 0644)

	return
}

func Stringify(m interface{}) string {
	json, _ := json.Marshal(m)
	return string(json)
}

func CreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
}
