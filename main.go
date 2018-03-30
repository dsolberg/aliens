package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
	"strconv"
//	"math/rand"
    //"time"
)

const alienName = "Bob"

const moveLimit = 10000

var city_map = map[string]map[string]string{}

var aliens = map[string]string{}

var total_aliens int


func gatherCliParameters() {

	// Import the CLI parameters
	flag.IntVar(&total_aliens, "aliens", 10000, "Total aliens to create")
	flag.Parse()
}


func loadCityMap() {

	// Open and read city map line at a time
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

//	var randdir int

	// Set direction map to map direction to random number

	// Cycle through total aliens

	for alien, _ := range aliens {

//		directions := map[int]string{0: "north", 1: "south", 2: "east", 3: "west"}
//		randdir = rand.Intn(len(directions))
		
		for _, destination := range city_map[aliens[alien]] {

			//println(destination)
			if city_map[destination] != nil {
				aliens[alien] = destination
				delete(city_map[aliens[alien]], destination)
			}


		// Check if destination city exists before moving
//		if city_map[city_map[aliens[alien]][directions[randdir]]] != nil {

			// Change existing city associated with each alien to a random city
//			aliens[alien] = city_map[aliens[alien]][directions[randdir]]
		
//		} else {

			// City in previous direction was destroyed - removing from list
//			delete(directions, randdir)
//			randdir = rand.Intn(len(directions))

//			if city_map[city_map[aliens[alien]][directions[randdir]]] != nil {

				// Change existing city associated with each alien to a random city
//				aliens[alien] = city_map[aliens[alien]][directions[randdir]]

//			} else {

				// City in previous direction was destroyed - removing from list
//				delete(directions, randdir)
//				randdir = rand.Intn(len(directions))

//				if city_map[city_map[aliens[alien]][directions[randdir]]] != nil {

					// Change existing city associated with each alien to a random city
//					aliens[alien] = city_map[aliens[alien]][directions[randdir]]

//				} else {

					// City in previous direction was destroyed - removing from list
//					delete(directions, randdir)
//					randdir = 0

//					if city_map[city_map[aliens[alien]][directions[randdir]]] != nil {

						// Change existing city associated with each alien to a random city
//						aliens[alien] = city_map[aliens[alien]][directions[randdir]]

//					} else {

						// Trapped - all roads destroyed
						//fmt.Println(alien + " is trapped!")
		//				aliens[alien] :=

//					}

//				}
//			}

		}
	}
}

func fightAliens() {

	// Alien gets up in the morning, has some eggs, loads his rifle and goes to find another alien
    for alien1Str, _ := range aliens {
	
    	// Make sure the first alien is not already blasted
		if _, notBlasted1 := aliens[alien1Str]; notBlasted1 { 

			// Attack alien begins to look for second alien to compare
			for alien2Str, _ := range aliens {

				// Make sure the second alien is not already blasted
				if _, notBlasted2 := aliens[alien2Str]; notBlasted2 {

					// Avoid alien tendencies to attack themselves when they see their reflection in a shiney surface
					if alien1Str != alien2Str {

						// Check if aliens are in the same city
						if aliens[alien1Str] == aliens[alien2Str] {
							
							fmt.Println("Match: " + alien1Str + " vs " + alien2Str + " in " + aliens[alien1Str] + " - Result: all destroyed" )
							
							// Attacker alien destroys his target - pew pew
							delete(aliens, alien2Str)

							// City engages Samson protocol and blows up the city so the aliens couldn't take it over
							delete(city_map, aliens[alien1Str])

							// Check to see if city has blown up other aliens
							for alienBlasted, _ := range aliens {
								
								// Ensure the other aliens are not already blown up
								if _, notBlasted := aliens[alienBlasted]; notBlasted {
								
									// Check to see if the other alien is not just the attacker
									if alien1Str == alienBlasted {
									
										// Detecting an alien in the same city 
										if aliens[alien1Str] == aliens[alienBlasted] {

											// Destroying alien found in the city
											fmt.Println(" - " + alienBlasted + " exploded in the crossfire at " + aliens[alienBlasted])
											delete(aliens, alienBlasted)
										}
									}
								}
							}


							// Attacker alien dies in the blast
							delete(aliens, alien1Str)

							fmt.Println()

							// Stop comparing with this city since it and all occupants are now destroyed
							break
						}	
					}
				}
			}
		}
	}
}

func printResult(report string) {
	cities := ""
	if report == cities {
        for _, city := range city_map {

			// Check to see if city has already blown up
			//if city_map[city] != nil {

			//}
			fmt.Println(city)
		}
	}
}


func main() {

	gatherCliParameters()
	loadCityMap()
    createAliens()
    moveAliens()

    // Begin alien invasion
    for tick := 0; tick < moveLimit; tick++ {
 		moveAliens()
    	fightAliens()
    	if len(aliens) == 0 {
    		fmt.Println("All aliens are dead.")
    //		printResult("cities")
    		break
    	}
    	print("Total Aliens Alive: ")
    	println(len(aliens))

    	print("Total Cities Remaining: ")
    	println(len(city_map))
    	
    	print("Round #: ")
    	fmt.Println(tick + 1)

    	//time.Sleep(1 * time.Second)
    }
	printResult("cities")
    //fmt.Println(aliens)

}