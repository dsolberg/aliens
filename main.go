package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
	"strconv"
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

			_ = endpoint
        	_ = direction
        	_ = directionsplit
       
        }
    }
}


func createAliens() {

	for i := 0; i < total_aliens; i++ {
	    aliens["Bob" + strconv.Itoa(i)] = "City"
	}
}


func main() {

	gatherCliParameters()
	loadCityMap()
	//fmt.Println(city_map)
    createAliens()
    fmt.Println(aliens["Bob2"])

}