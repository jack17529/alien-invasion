package destroy

import (
	"fmt"
	st "github.com/faith/alien-invasion/pkg/Structure"
)

func DestroyAlien(cities []st.City, aliens []st.Alien, c *(st.City), i int, community map[int]int) ([](st.City), []st.Alien) {
	
	// time complexity O(1)
	//fmt.Println("Aliens in city before destruction",c.Aliens)
	ind := community[c.Aliens[i].Index] - 1
	fmt.Println(ind)
	delete(community, ind)
	l2 := len(aliens) - 1
	
	//index of alien with l2 index
	inx := aliens[l2].Index
	if ind < l2 {
		aliens[ind], aliens[l2] = aliens[l2], aliens[ind]
		community[inx] = ind + 1
	}
	if len(aliens) > 0 {
		aliens = aliens[:l2]
	}
	
	//fmt.Println("Aliens in city after destruction",c.Aliens)
	//fmt.Println("Aliens after destroying an alien", aliens)
	return cities, aliens
}

func DestroyCity(cities []st.City, aliens []st.Alien, r int, world map[string]int, cityGroup map[int]int, destroyedMap map[string]int, citiesWith1Alien []int) ([]st.City, []st.Alien, []int) {

	// delete the city from the array of ciities and map.
	// update cityName too.
	// O(1) time complexity.
	destroyedMap[cities[r].Name] = 1

	delete(world, cities[r].Name)
	l := len(cities) - 1
	if len(citiesWith1Alien) > 0 {
		if len(cities[l].Aliens) == 1 && cityGroup[cities[l].Aliens[0].Index] > 0 {
			indx := cityGroup[cities[l].Aliens[0].Index] - 1
			citiesWith1Alien[indx] = r
		}
	}

	if r < l {
		cities[l], cities[r] = cities[r], cities[l]
		world[cities[r].Name] = r + 1
		//world[cities[l].Name]=0
	}
	if len(cities) > 0 {
		cities = cities[:l]
	}
	
	//fmt.Println(cities)
	return cities, aliens, citiesWith1Alien
}
