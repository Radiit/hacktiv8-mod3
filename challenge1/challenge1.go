package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for true {

		RandomWater := rand.Intn(100-1) + 1
		randomWind := rand.Intn(100-1) + 1

		data := map[string]interface{}{
			"water": RandomWater,
			"wind":  randomWind,
		}

		JsonRequest, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(JsonRequest))

		req.Header.Set("Content-type", "application/json")
		if err != nil {
			panic(err)
		}

		response, eer := client.Do(req)
		if eer != nil {
			log.Fatalln(err)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))
		result(RandomWater, randomWind)
		time.Sleep(15 * time.Second)
	}

}

func result(InputWater int, InputWind int) {
	water := ""
	wind := ""

	if InputWater <= 5 {
		water = "aman"
	} else if InputWater >= 6 && InputWater <= 8 {
		water = "siaga"
	} else {
		water = "bahaya"
	}

	if InputWind <= 6 {
		wind = "aman"
	} else if InputWind >= 7 && InputWind <= 15 {
		wind = "siaga"
	} else {
		wind = "bahaya"
	}

	fmt.Printf("water: %s\n", water)
	fmt.Printf("wind: %s\n", wind)
}
