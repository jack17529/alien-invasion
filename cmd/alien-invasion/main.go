package main

import (
	"flag"
	"fmt"
	cm "github.com/faith/alien-invasion/pkg/Come"
	mv "github.com/faith/alien-invasion/pkg/Move"
	ut "github.com/faith/alien-invasion/pkg/Util"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println(`Alien-Invasion

    Reads in pre-defined map data from given map file, then runs simulation for n aliens wandering in cities.

    Usage: main [flags]
    				`)

		flag.PrintDefaults()
	}

	mapFile := flag.String("map", "storage/sample_map", "map file path")
	numAliens := flag.Int("aliens", 1, "number of aliens to be deployed")

	flag.Parse()

	var cityGroup = make(map[int]int)
	var destroyedMap = make(map[string]int)
	var citiesWith1Alien []int

	// Base condition.
	if *numAliens <= 0 {
		fmt.Println("The map requires at least one alien!")
		os.Exit(0)
	}

	// creating a map.
	cities, world, err := ut.MakeMap(*mapFile)
	if err != nil {
		log.Fatal(err)
	}

	// the number of aliens who could not come the world.
	remAliens := *numAliens - len(cities)*2
	if remAliens > 0 {
		*numAliens = len(cities) * 2
	}

	// simulates the coming of aliens in the map.
	cities, aliens, community := cm.AliensCome(cities, *numAliens, world, cityGroup, destroyedMap)
	if err != nil {
		log.Fatal(err)
	}

	if remAliens > 0 {
		fmt.Printf("the whole world got destroyed and %v aliens couldn't even come to the world \n", remAliens)
	}

	// prints the initial arrays of aliens and cities.
	fmt.Println("print the aliens \n", aliens)
	fmt.Println("print the cities \n", cities)

	// a function to intialize.
	cities, aliens, citiesWith1Alien = ut.Util(cities, aliens, world, cityGroup, citiesWith1Alien)
	var lv int
	if(len(citiesWith1Alien)>0){
		lv = 100000000/(4*(len(citiesWith1Alien)))
	}
	
	// first and last check before entering into the loop.
	aliensDoneMoving := 0
	val:=ut.CheckGameOver(cities, aliens, destroyedMap, citiesWith1Alien, *mapFile, aliensDoneMoving)
	if(val>0){
		goto done
	}

	//until we hit the game's end (infinite loop escaped by 'os.Exit once conditions are met')
	// Executing 10^8 instructions take around 1 second on an average machine.
	// 10^9 instructions can be executed in a few seconds in good servers.
	for i := 1; i <= lv; i++ {
		fmt.Println("print the cities with 1 alien \n", citiesWith1Alien)
		randSeq := ut.GetRandomSequence(len(citiesWith1Alien)) // time complexity of O(len(citiesWith1Alien))

		// move all eligible aliens in citiesWith1Alien
		cities, aliens, citiesWith1Alien, aliensDoneMoving = mv.AliensMove(cities, aliens, randSeq, world, community, cityGroup, destroyedMap, citiesWith1Alien, aliensDoneMoving) // time complexity of O(4*len(citiesWith1Alien))

		// check game over.
		val=ut.CheckGameOver(cities, aliens, destroyedMap, citiesWith1Alien, *mapFile, aliensDoneMoving)	// time complexity O(1)
		if(val>0){
			goto done
		}
	}

	done:
	if(val==1){
		err := ut.PrintRemainingMap(*mapFile, destroyedMap)
		if err != nil {
			log.Fatal(err)
		}
	}
	if(val==0){
		fmt.Println("The data provided by you is overwhelming, do adjustment according to the lv value provided on Line 70 in main.go or no known algorithm to humans can solve the problem for you.")
	}
}
