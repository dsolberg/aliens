package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var city_map = map[string]map[string]string{}

func loadcitymap() {

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


func main() {
	loadcitymap()
	fmt.Println(city_map)
}