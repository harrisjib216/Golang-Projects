package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type City struct {
	name string 
	temp float64
}

type CityScore struct {
	name string
	min float64
	avg float64
	max float64
	_hits int
}

func seedRandom() {
	rand.Seed(time.Now().UnixNano())
}

func generateRows() []City {
	cityNames := []string{
		"Alexander City",
		"Andalusia",
		"Anniston",
		"Athens",
		"Atmore",
		"Auburn",
		"Bessemer",
		"Birmingham",
		"Chickasaw",
		"Clanton",
		"Cullman",
		"Decatur",
		"Demopolis",
		"Dothan",
		"Enterprise",
		"Eufaula",
		"Florence",
		"Fort Payne",
		"Gadsden",
		"Greenville",
		"Guntersville",
		"Huntsville",
		"Jasper",
		"Marion",
		"Mobile",
		"Montgomery",
		"Opelika",
		"Ozark",
		"Phenix City",
		"Prichard",
		"Scottsboro",
		"Selma",
		"Sheffield",
		"Sylacauga",
		"Talladega",
		"Troy",
		"Tuscaloosa",
		"Tuscumbia",
		"Tuskegee",
	}

	cities := []City{}

	for i := 0; i < 100000; i++ {
		cities = append(cities, City{
			name: cityNames[rand.Intn(len(cityNames))],
			temp: rand.Float64() * 100,
		})
	}

	return cities
}

func algo(cities []City) []CityScore {
	cache := map[string]CityScore{}

	// score min, max, calc avg ltr
	for _, city := range (cities) {
		if cityScore, ok := cache[city.name]; ok {
			cityScore.min = math.Min(city.temp, cityScore.min)
			cityScore.max = math.Max(city.temp, cityScore.max)
			
			// sum all the temps, average later
			cityScore.avg += city.temp
			cityScore._hits += 1

			cache[city.name] = cityScore
		} else {
			cache[city.name] = CityScore{
				name: city.name,
				min: city.temp,
				max: city.temp,
				avg: city.temp,
				_hits: 0,
			}
		}
	}

	// calc avg
	res := []CityScore{}
	for _, cityScore := range cache {
		cityScore.avg = cityScore.avg / float64(cityScore._hits)
		res = append(res, cityScore)
	}

	return res
}

func main() {
	seedRandom()

	res := algo(generateRows())
	fmt.Println(res)
}