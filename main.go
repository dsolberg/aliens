package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const alienName = "Bob"

const moveLimit = 10000


func loadCityMap() map[string]map[string]string {
	// Assertion - City map will always be in the format described in the challenge document including spacing and ='s'

	// Open and read city map line at a time
	var city_map = map[string]map[string]string{}

	file, err := os.Open("maps/fullcitymap.txt")
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
		cut := len(city[0]) + 1

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

			endpoint := directionsplit[1]
			direction := directionsplit[0]

			city_map[city[0]][direction] = endpoint

		}
	}

	return city_map
}

func createAliens(total_aliens int, city_map map[string]map[string]string) map[string]string {

	var aliens = map[string]string{}

	// Cycle through the total number of aliens
	for i := 0; i < total_aliens; i++ {

		// Cycle through city map to find cities for aliens to blow up
		for city, _ := range city_map {

			// Assign each alien a city from the map which is naturally random
			aliens[alienName+strconv.Itoa(i)] = city
		}
	}

	return aliens
}

func moveAliens(aliens map[string]string, city_map map[string]map[string]string) map[string]string {

	// Get every living alien
	for alien, _ := range aliens {

		// Obtain every destination alive
		for _, destination := range city_map[aliens[alien]] {

			// Alien moves to next live city
			aliens[alien] = destination
			break
		}
	}

	return aliens
}

func blowUpSameCityAliens(alien1Str string, aliens map[string]string) map[string]string {

	// Check to see if city has blown up other aliens
	for alienBlasted, _ := range aliens {

		// Cond 1: Ensure the other aliens are not already blown up
		// Cond 2: Check to see if the other alien is not just the attacker
		// Cond 3: Detecting an alien in the same city

		_, notBlasted := aliens[alienBlasted]
		if (notBlasted) && (alien1Str != alienBlasted) && (aliens[alien1Str] == aliens[alienBlasted]) {

			// Destroying alien found in the city
			fmt.Println(" - " + alienBlasted + " exploded in the crossfire at " + aliens[alienBlasted])
			delete(aliens, alienBlasted)
		}
	}

	return aliens
}

func blowUpDestroyedCityRoutes(alien1Str string, aliens map[string]string, city_map map[string]map[string]string) (map[string]string, map[string]map[string]string) {

	// Obtain every alive city connected to city blowing up
	for _, aliveCity := range city_map[aliens[alien1Str]] {

		// Obtain every route in the connected cities
		for deadRoute, deadCity := range city_map[aliveCity] {

			// Check if the route exists and if it does blow it up too
			if aliens[alien1Str] == deadCity {

				// Exploding city blows up the route to connected
				delete(city_map[aliveCity], deadRoute)

			}
		}
	}

	return aliens, city_map
}

func attackSequence(alien1Str string, aliens map[string]string, city_map map[string]map[string]string) (map[string]string, map[string]map[string]string) {

	// Attack alien begins to look for second alien to compare
	for alien2Str, _ := range aliens {

		_, notBlasted2 := aliens[alien2Str]

		// Cond 1: Make sure the second alien is not already blasted
		// Cond 2: Avoid alien tendencies to attack themselves when they see their reflection in a shiney surface
		// Cond 3: Check if aliens are in the same city

		if (notBlasted2) && (alien1Str != alien2Str) && (aliens[alien1Str] == aliens[alien2Str]) {

			// Aliens are in the same city!
			fmt.Println("Match: " + alien1Str + " vs " + alien2Str + " in " + aliens[alien1Str] + " - Result: all destroyed")

			// Attacker alien destroys his target - pew pew
			delete(aliens, alien2Str)

			// Blow up the routes to the destroyed city
			aliens, city_map = blowUpDestroyedCityRoutes(alien1Str, aliens, city_map)

			// City engages Samson Option and blows up the city so the aliens couldn't take it over
			delete(city_map, aliens[alien1Str])

			// Every alien in the city explodes!
			aliens = blowUpSameCityAliens(alien1Str, aliens)

			// Attacker alien dies in the blast
			delete(aliens, alien1Str)

			fmt.Println()

			// Stop comparing with this city since it and all occupants are now destroyed
			break
		}
	}

	return aliens, city_map
}

func printResults(report string, aliens map[string]string, city_map map[string]map[string]string) {
	if report == "cities" {

		// Obtain every city alive
		for city, _ := range city_map {

			directionLine := ""

			// Obtain every direction per alive city
			for directions, _ := range city_map[city] {

				// Write the base line
				directionLine = directionLine + " " + directions + "=" + city_map[city][directions]
			}

			// Print every live city
			cityLine := city + directionLine
			fmt.Println(cityLine)
		}
	}

	if report == "aliens" {

		// Obtain and print every live alien
		for alien, _ := range aliens {

			alienLine := alien + " " + aliens[alien]
			fmt.Println(alienLine)
		}
	}

	return
}

func main() {

	var total_aliens int
	
	// Import the CLI parameters
	flag.IntVar(&total_aliens, "aliens", 2000, "Total aliens to create")
	flag.Parse()

	city_map := loadCityMap()

	aliens := createAliens(total_aliens, city_map)

	// Begin alien invasion
	for tick := 0; tick < moveLimit; tick++ {
		aliens := moveAliens(aliens, city_map)

		// Cycle through each alien
		for alien1Str, _ := range aliens {

			// Make sure the first alien is not already blasted
			if _, notBlasted1 := aliens[alien1Str]; notBlasted1 {

				// Find second alien to fight
				attackSequence(alien1Str, aliens, city_map)
			}
		}

		fmt.Print("Total Aliens Alive: ")
		fmt.Println(len(aliens))

		fmt.Print("Total Cities Remaining: ")
		fmt.Println(len(city_map))

		fmt.Print("Round #: ")
		fmt.Println(tick + 1)

		if len(aliens) == 0 {
			fmt.Println("All aliens are dead.")
			break
		}
	}

	printResults("cities", aliens, city_map)

	if len(aliens) != 0 {
		printResults("aliens", aliens, city_map)
	}

}
