package move

import (
	"fmt"
	dt "github.com/faith/alien-invasion/pkg/Destroy"
	st "github.com/faith/alien-invasion/pkg/Structure"
	"math/rand"
	"time"
)

func RemoveFromCitiesWith1Alien(cities []st.City, aliens []st.Alien, oldCity *(st.City), cityGroup map[int]int, citiesWith1Alien []int) []int {

	// O(1) time complexity.
	// Remove with pointers so that the internal structure of aliens does not get destroyed.
	ind2 := cityGroup[(oldCity).Aliens[0].Index] - 1
	delete(cityGroup, (oldCity).Aliens[0].Index)

	l3 := len(citiesWith1Alien) - 1
	inx2 := cities[citiesWith1Alien[l3]].Aliens[0].Index

	if ind2 < l3 {
		citiesWith1Alien[ind2], citiesWith1Alien[l3] = citiesWith1Alien[l3], citiesWith1Alien[ind2]
		cityGroup[inx2] = ind2 + 1
	}

	if len(aliens) > 0 {
		citiesWith1Alien = citiesWith1Alien[:l3]
	}

	//fmt.Println(citiesWith1Alien)
	return citiesWith1Alien
}

func AddInCitiesWith1Alien(cities []st.City, aliens []st.Alien, oldCity *(st.City), world map[string]int, cityGroup map[int]int, citiesWith1Alien []int) []int {

	// O(1) time complexity.
	// Add with pointers so that the internal structure of aliens does not get destroyed.
	citiesWith1Alien = append(citiesWith1Alien, world[oldCity.Name]-1)
	cityGroup[(oldCity).Aliens[0].Index] = len(citiesWith1Alien)

	//fmt.Println(citiesWith1Alien)
	return citiesWith1Alien
}

func Remove2(cities []st.City, aliens []st.Alien, newCity *(st.City), world map[string]int, community map[int]int, cityGroup map[int]int, destroyedMap map[string]int, citiesWith1Alien []int) ([]st.City, []st.Alien, []int) {

	// O(1) time complexity.
	newCity.IsDestroyed = true
	// delete first alien from the alien array of the city.
	newCity.Aliens[0].IsDead = true
	cities, aliens = dt.DestroyAlien(cities, aliens, &cities[world[newCity.Name]-1], 0, community)	//  O(1) time complexity.

	// delete the second alien from the alien array of the city.
	newCity.Aliens[1].IsDead = true
	cities, aliens = dt.DestroyAlien(cities, aliens, &cities[world[newCity.Name]-1], 1, community)	
	r := world[newCity.Name] - 1
	fmt.Printf("city with index %v, was destroyed\n", r)
	cities, aliens, citiesWith1Alien = dt.DestroyCity(cities, aliens, r, world, cityGroup, destroyedMap, citiesWith1Alien)	// O(1) time complexity.

	// remove the newCity from the citiesWith1Alien too

	return cities, aliens, citiesWith1Alien
}

func AliensMove(cities []st.City, aliens []st.Alien, a []int, world map[string]int, community map[int]int, cityGroup map[int]int, destroyedMap map[string]int, citiesWith1Alien []int, aliensDoneMoving int) ([]st.City, []st.Alien, []int, int) {

	// moves an alien at a time.
	// O(4*len(citiesWith1Alien)) time complexity.
	var oldCity *(st.City)
	var alien *(st.Alien)
	fmt.Println("The random sequence is ", a)

	// aliens move totally randomly one alien at a time, an alien may get contigous chances though(which is also random).
	for i := 0; i < len(citiesWith1Alien); i++ { 
		fmt.Println("value of i is ", i)
		fmt.Println("Index of current city is ", citiesWith1Alien[a[i]])
		oldCity = &cities[citiesWith1Alien[a[i]]]
		fmt.Println("Index of current alien is ", oldCity.Aliens[0].Index)
		alien = &(oldCity.Aliens[0])

		// randomly selecting in which direction to move, also checking trapped condition at the same time.
		var nextCity string
		rng := 4
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(rng)
		// fmt.Println(r)
		p := 1 // a flag to check the trapped condition.

		for i := 0; i < 4; i++ {
			if oldCity.North != "" && destroyedMap[oldCity.North] == 0 {
				p = 0
				if r == 0 {
					nextCity = oldCity.North
					break
				}
				r -= 1
			}
			if oldCity.East != "" && destroyedMap[oldCity.East] == 0 {
				p = 0
				if r== 0 {
					nextCity = oldCity.East
					break
				}
				r -= 1
			}
			if oldCity.South != "" && destroyedMap[oldCity.South] == 0 {
				p = 0
				if r == 0 {
					nextCity = oldCity.South
					break
				}
				r -= 1
			}
			if oldCity.West != "" && destroyedMap[oldCity.West] == 0 {
				p = 0
				if r == 0 {
					nextCity = oldCity.West
					break
				}
				r -= 1
			}
		}

		if p == 1 {
			alien.IsTrapped = true
			fmt.Println("An alien is trapped in", oldCity.Name)
			//remove the alien from 1 alien city slices.
			citiesWith1Alien = RemoveFromCitiesWith1Alien(cities, aliens, oldCity, cityGroup, citiesWith1Alien)

			return cities, aliens, citiesWith1Alien, aliensDoneMoving
		} else {
			var newCity *(st.City)
			newCity = &cities[world[nextCity]-1]
			alien.Moves += 1
			
			// add the alien in the newCity.
			fmt.Println("From ", oldCity.Name)
			fmt.Println("To ", newCity.Name)
			newCity.Aliens = append(newCity.Aliens, *alien)

			citiesWith1Alien = RemoveFromCitiesWith1Alien(cities, aliens, oldCity, cityGroup, citiesWith1Alien)
			fmt.Println(newCity.Aliens)
			fmt.Println(oldCity.Aliens)
			oldCity.Aliens = oldCity.Aliens[:0]

			// check that the new city already have an alien
			if len(newCity.Aliens) == 2 {

				fmt.Printf("city %v is destroyed by alien%v and alien%v \n", newCity.Name, newCity.Aliens[0].Index, newCity.Aliens[1].Index)
				citiesWith1Alien = RemoveFromCitiesWith1Alien(cities, aliens, newCity, cityGroup, citiesWith1Alien)
				fmt.Println("removed the new city from the city with 1 Alien.")
				
				//update citiesWith1alien after destroying a city for the city that came in place of that destroyed city.
				cities, aliens, citiesWith1Alien = Remove2(cities, aliens, newCity, world, community, cityGroup, destroyedMap, citiesWith1Alien)
				//fmt.Println("print the cities \n", cities)
				return cities, aliens, citiesWith1Alien, aliensDoneMoving
			}

			// if alien has completed all it's moves(10,000) remove from alien array and CityWith1AlienArray.
			if alien.Moves >= 10000 {
				alien.IsDoneMoving = true
				aliensDoneMoving+=1
				fmt.Println("An alien has completed 10,000 moves and is done moving. The alien stopped in city ", newCity.Name)

				//fmt.Println("print the cities \n", cities)
				return cities, aliens, citiesWith1Alien, aliensDoneMoving
			} else {
				// Add the current city in the one alien city group.
				citiesWith1Alien = AddInCitiesWith1Alien(cities, aliens, newCity, world, cityGroup, citiesWith1Alien)
				//fmt.Println("print the cities \n", cities)
			}
		}
	}

	return cities, aliens, citiesWith1Alien, aliensDoneMoving
}
