package come

import (
	"fmt"
	dt "github.com/faith/alien-invasion/pkg/Destroy"
	st "github.com/faith/alien-invasion/pkg/Structure"
	"math/rand"
	"time"
)

func Remove1(cities []st.City, aliens []st.Alien, r int, world map[string]int, community map[int]int, cityGroup map[int]int, destroyedMap map[string]int) ([]st.City, []st.Alien) {

	// delete the alien from the array of aliens and map.
	// O(1) time complexity.
	cities[r].IsDestroyed = true
	
	cities[r].Aliens[0].IsDead = true
	var c *(st.City)
	c = &cities[r]
	cities, aliens = dt.DestroyAlien(cities, aliens, c, 0, community)
	
	cities[r].Aliens[1].IsDead = true
	c = &cities[r]
	cities, aliens = dt.DestroyAlien(cities, aliens, c, 1, community)
	
	temp := make([]int, 0)
	cities, aliens, temp = dt.DestroyCity(cities, aliens, r, world, cityGroup, destroyedMap, temp)

	return cities, aliens
}

func AliensCome(cities []st.City, num int, world map[string]int, cityGroup map[int]int, destroyedMap map[string]int) ([]st.City, []st.Alien, map[int]int) {

	// Simulates coming of aliens in the world.
	// As soon as 2 aliens land in the same city they kill each other and destroy the city.
	// No alien start to move until all the aliens have come to the map.
	// time complexity of O(maximum no. of aliens that can come)
	var aliens []st.Alien
	var community = make(map[int]int)

	for i := 0; i < num; i++ {
		fmt.Println(i)
		//create alien
		alien := st.Alien{Index: i}
		
		rng := len(cities)
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(rng)
		fmt.Printf("City randomly selected is %v \n", cities[r].Name)
		
		// placing the alien inside the city.
		cities[r].Aliens = append(cities[r].Aliens, alien)
		aliens = append(aliens, alien)
		//fmt.Println(aliens)
		community[i] = i + 1
		
		if len(cities[r].Aliens) == 2 {
			fmt.Printf("city %v is destroyed by alien%v and alien%v \n", cities[r].Name, cities[r].Aliens[0].Index, cities[r].Aliens[1].Index)

			cities, aliens = Remove1(cities, aliens, r, world, community, cityGroup, destroyedMap)
			continue
		}
	}
	return cities, aliens, community
}
