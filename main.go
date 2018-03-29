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
    "time"
)

const alienName = "Bob"

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
			aliens[alienName + strconv.Itoa(i)] = city
		}
	}
}

func moveAliens() {

	var randdir int

	// Set direction map to map direction to random number
	directions := map[int]string{0: "north", 1: "south", 2: "east", 3: "west"}

	// Cycle through total aliens
	for i := 0; i < total_aliens; i++ {
		randdir = rand.Intn(4)

		// Check if destination city exists before moving
		if city_map[city_map[aliens[alienName + strconv.Itoa(i)]][directions[randdir]]] == nil {

			// Change existing city associated with each alien to a random city
			aliens[alienName + strconv.Itoa(i)] = city_map[aliens[alienName + strconv.Itoa(i)]][directions[randdir]]
		}
		// TODO - Try to move to other non-destroyed city
	}
}

func fightAliens() {

	// Loop for first alien to compare
	for alien1 := 0; alien1 < total_aliens; alien1++ {

		// Check to see if city has already blown up
		if city_map[aliens[alienName + strconv.Itoa(alien1)]] != nil {

			// Loop for second alien to compare
			for alien2:= 0; alien2 < total_aliens; alien2++ {

				// Check if alien is dead
				if aliens[alienName + strconv.Itoa(alien1)] != "dead" {
					
					// Check if alien is being compared against itself
					if alien1 != alien2 {

						// Check if aliens are in the same city
						if aliens[alienName + strconv.Itoa(alien1)] == aliens[alienName + strconv.Itoa(alien2)] {
							fmt.Println("Match: " + alienName + strconv.Itoa(alien1) + " vs " + alienName + strconv.Itoa(alien2) + " in " + aliens[alienName + strconv.Itoa(alien1)] + " - Result: all non-destroyed" )
							
							// Blow up city
							delete(city_map,aliens[alienName + strconv.Itoa(alien1)])

							// Aliens kill eachother in the smouldering city explosion - pew pew
							aliens[alienName + strconv.Itoa(alien1)] = "dead"
							aliens[alienName + strconv.Itoa(alien2)] = "dead"

						}
					}
				}	
			}

		} else {

			// Alien in blown up city - presumed dead - collatoral damage
			aliens[alienName + strconv.Itoa(alien1)] = "dead"
			fmt.Println("	- Poor " + alienName + strconv.Itoa(alien1) + " died in the crossfire")
		}
	}
}


func main() {

	gatherCliParameters()
	loadCityMap()
    createAliens()
    moveAliens()

    // Begin alien invasion
    for tick := 1; tick < 2; tick++ {
    	fmt.Println(tick)
		moveAliens()
    	fightAliens()
    	time.Sleep(2 * time.Second)
    }
    //fmt.Println(aliens)

}