package util

import (
	"bufio"
	"fmt"
	st "github.com/faith/alien-invasion/pkg/Structure"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetRandomSequence(n int) []int {

	// This function is a very well known shuffle function for generating random permutations of numbers in O(n) time complexity.
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}

func MakeMap(filename string) ([]st.City, map[string]int, error) {

	// Time Complexity of O(4*(no. of lines in text or no. of cities))
	var cities [](st.City)
	var world = make(map[string]int)
	var err error
	file, err := os.Open(filename)

	if err != nil {
		return cities, world, err
	}

	var N = make(map[string]string) // {city name, city name} value is the city north of key.
	var E = make(map[string]string) // {city name, city name} value is the city east of key.
	var S = make(map[string]string) // {city name, city name} value is the city south of key.
	var W = make(map[string]string) // {city name, city name} value is the city west of key.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var Present = make(map[string]int)
		txt := scanner.Text()

		str := strings.Split(txt, " ")
		city := st.City{}

		city.Name = str[0]
		Present[city.Name] = 1
		str = str[1:]
		//scan the strings according to contains function and build the map.
		for _, adj := range str {
			if strings.Contains(adj, "north=") {
				city.North = strings.Replace(adj, "north=", "", 1)
				if N[city.Name] == "" && Present[city.North] == 0 {
					N[city.Name] = city.North
					Present[city.North] = 1
				} else {
					fmt.Println("There can't be the a city name repeated in the same line of the map file.")
					os.Exit(0)
				}
			} else if strings.Contains(adj, "east=") {
				city.East = strings.Replace(adj, "east=", "", 1)
				if E[city.Name] == "" && Present[city.East] == 0 {
					E[city.Name] = city.East
					Present[city.East] = 1
				} else {
					fmt.Println("There can't be the a city name repeated in the same line of the map file.")
					os.Exit(0)
				}
			} else if strings.Contains(adj, "south=") {
				city.South = strings.Replace(adj, "south=", "", 1)
				if S[city.Name] == "" && Present[city.South] == 0 {
					S[city.Name] = city.South
					Present[city.South] = 1
				} else {
					fmt.Println("There can't be the a city name repeated in the same line of the map file.")
					os.Exit(0)
				}
			} else if strings.Contains(adj, "west=") {
				city.West = strings.Replace(adj, "west=", "", 1)
				if W[city.Name] == "" && Present[city.West] == 0 {
					W[city.Name] = city.West
					Present[city.West] = 1
				} else {
					fmt.Println("There can't be the a city name repeated in the same line of the map file.")
					os.Exit(0)
				}
			}
		}

		//initialize the world map.

		if world[city.Name] == 0 {
			cities = append(cities, city)
			world[city.Name] = len(cities)
		}

	}

	// check all maps and their corresponding inverse maps.
	// if a key, value pair is present in one and absent in the other then it is incomplete map.
	for key, val := range N {
		if S[val] != key {
			fmt.Println("The map provided is not complete.")
			os.Exit(0)
		}
	}

	for key, val := range E {
		if W[val] != key {
			fmt.Println("The map provided is not complete.")
			os.Exit(0)
		}
	}

	for key, val := range S {
		if N[val] != key {
			fmt.Println("The map provided is not complete.")
			os.Exit(0)
		}
	}

	for key, val := range W {
		if E[val] != key {
			fmt.Println("The map provided is not complete.")
			os.Exit(0)
		}
	}

	return cities, world, err
}

func Util(cities []st.City, aliens []st.Alien, world map[string]int, cityGroup map[int]int, citiesWith1Alien []int) ([]st.City, []st.Alien, []int) {

	// This functions aim is to initialize the world map, cityGroup map and citiesWith1Alien slice.
	// time complexity of O(len(cities))
	for i := 0; i < len(cities); i++ {
		world[cities[i].Name] = i + 1
		if len(cities[i].Aliens) == 1 {
			citiesWith1Alien = append(citiesWith1Alien, i)
			cityGroup[cities[i].Aliens[0].Index] = len(citiesWith1Alien)
		}
	}

	return cities, aliens, citiesWith1Alien

}

func PrintRemainingMap(filename string, destroyedMap map[string]int) error {

	// Re-read the map from file and print the remaining map.
	// Time Complexity of O(4*(no. of lines in text or no. of cities))
	fmt.Println("\n\n\nPRINTING THE REMAINING MAP!")
	var err error
	//open map file
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		// split txt according to the " ".
		str := strings.Split(txt, " ")
		city := st.City{}

		city.Name = str[0]
		if destroyedMap[city.Name] == 1 {
			continue
		}

		fmt.Printf("%v ", str[0])
		str = str[1:]

		// scan the strings according to contains function and print if it is part of the remaining map.
		for _, adj := range str {
			if strings.Contains(adj, "north=") {
				city.North = strings.Replace(adj, "north=", "", 1)
				if destroyedMap[city.North] == 0 {
					fmt.Printf("north=%v ", city.North)
				}
			} else if strings.Contains(adj, "east=") {
				city.East = strings.Replace(adj, "east=", "", 1)
				if destroyedMap[city.East] == 0 {
					fmt.Printf("east=%v ", city.East)
				}
			} else if strings.Contains(adj, "south=") {
				city.South = strings.Replace(adj, "south=", "", 1)
				if destroyedMap[city.South] == 0 {
					fmt.Printf("south=%v ", city.South)
				}
			} else if strings.Contains(adj, "west=") {
				city.West = strings.Replace(adj, "west=", "", 1)
				if destroyedMap[city.West] == 0 {
					fmt.Printf("west=%v ", city.West)
				}
			}
		}
		fmt.Printf("\n")
	}
	return err
}

func CheckGameOver(cities []st.City, aliens []st.Alien, destroyedMap map[string]int, citiesWith1Alien []int, mapFile string, aliensDoneMoving int) (int){
	if(len(citiesWith1Alien) == 1 && aliensDoneMoving == 0) {
		fmt.Println("Only 1 active alien in the remaining map and all others are trapped or dead, it is hence trapped!, GAME OVER")
		//fmt.Println("print the aliens \n", aliens)
		return 1
	}
	
	if len(citiesWith1Alien) == 0 {
		fmt.Println("All the aliens are dead or trapped or have moved atleast 10,000 times!, GAME OVER")
		//fmt.Println("print the cities \n", cities)
		return 1
	}

	if len(cities) == 0 {
		fmt.Println("The world has been destroyed!, GAME OVER")
		//fmt.Println("print the aliens \n", aliens)
		return 2
	}
	
	return 0
}
