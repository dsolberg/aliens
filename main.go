package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
	"strconv"
	"math/rand"
)

var city_map = map[string]map[string]string{}

var total_aliens int

var aliens = map[string]string{}

//type alien struct {
//	id string
//	city string
//}

func gatherCliParameters() {

	// Import the CLI parameters
	flag.IntVar(&total_aliens, "aliens", 10000, "Total aliens to create")
	flag.Parse()
}


func loadCityMap() {

	// Open and read city map line at a time
	file, err := os.Open("maps/citymap.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {

		// Read line into variable
		l := scanner.Text()

		// Split variables
	    city := strings.Split(l, " ")

	    // Cut the city name out
	    cut := len(city[0])+1

	    // Cut the directions out
	    directionstrim := l[cut:len(l)]
	    directionscut := strings.Split(directionstrim, " ")
	    
	    str := ""

	    // Declare city submap
        city_map[city[0]] = make(map[string]string)

        for k, routes := range directionscut {
        	_ = k
        	str = routes
        	directionsplit := strings.Split(str, "=")
        
        	//fmt.Println(directionsplit)
        	endpoint := directionsplit[1]
        	direction := directionsplit[0]

        	city_map[city[0]][direction] = endpoint

		}
    }
}


func createAliens() {

    // Cycle through the total number of aliens
	for i := 0; i < total_aliens; i++ {

		// Cycle through city map to find cities for aliens to blow up
		for city, _ := range city_map {

			// Assign each alien a city from the map which is naturally random
			aliens["Bob" + strconv.Itoa(i)] = city
		}
	}
}

func moveAliens() {

	// Set direction map to map direction to random number
	directions := map[int]string{1: "north", 2: "south", 3: "east", 4: "west"}

	// Cycle through total aliens
	for i := 0; i < total_aliens; i++ {

		// Change existing city associated with each alien to a random city
		aliens["Bob" + strconv.Itoa(i)] = city_map[aliens["Bob" + strconv.Itoa(i)]][directions[rand.Intn(4)]]
	}
}


func main() {

	gatherCliParameters()
	loadCityMap()
    createAliens()
    moveAliens()
    fmt.Println(aliens["Bob2"])

}