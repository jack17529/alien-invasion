package destroy

import (
	//"fmt"
	st "github.com/faith/alien-invasion/pkg/Structure"
	"testing"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func TestDestroy(t *testing.T) /*([]st.City, []st.Alien, int, map[string]int,[]int, map[int]int, map[string]int, map[int]int ) */ {
	a0 := st.Alien{IsTrapped: false, IsDead: false, IsDoneMoving: false, Index: 0, Moves: 0}
	a1 := st.Alien{IsTrapped: false, IsDead: false, IsDoneMoving: false, Index: 1, Moves: 0}
	aliens2 := []st.Alien{a0, a1}
	//fmt.Println(aliens2)
	cities2 := []st.City{{IsDestroyed: false, Name: "Foo", North: "Bar", East: "", South: "", West: "", Aliens: []st.Alien{a0, a1}}, {IsDestroyed: false, Name: "Bar", North: "", East: "", South: "Foo", West: "", Aliens: []st.Alien{}}}

	world2 := map[string]int{"Foo": 1, "Bar": 2}
	citiesWith1Alien2 := []int{}
	cityGroup2 := map[int]int{}
	destroyedMap2 := map[string]int{}
	community2 := map[int]int{0: 1, 1: 2}

	cities2, aliens2 = DestroyAlien(cities2, aliens2, &cities2[world2["Foo"]-1], 0, community2)
	if len(cities2) != 2 {
		t.Fatalf("Number of remaining cities differ, got: %d, want: %d", len(cities2), 2)
	}
	if aliens2[0].Index != 1 {
		t.Fatalf("The remaining alien differ, got: %d, want: %d", aliens2[0].Index, 1)
	}
	if community2[0] != 0 {
		t.Fatalf("The remaining community differ after deleting 1st alien, got: %d, want: %d", community2[0], 0)
	}

	//Use(world2, citiesWith1Alien2, cityGroup2, destroyedMap2)

	cities2, aliens2 = DestroyAlien(cities2, aliens2, &cities2[world2["Foo"]-1], 1, community2)

	if len(cities2) != 2 {
		t.Fatalf("Number of remaining cities differ, got: %d, want: %d", len(cities2), 2)
	}
	if len(aliens2) != 0 {
		t.Fatalf("Number of remaining alines differ, got: %d, want: %d", len(aliens2), 0)
	}
	if len(community2) == 0 {
		t.Fatalf("The remaining community differ after deleting 2nd alien, got: %d, want: %d", len(community2), 0)
	}
	r2 := world2["Foo"] - 1
	cities2, aliens2, citiesWith1Alien2 = DestroyCity(cities2, aliens2, r2, world2, cityGroup2, destroyedMap2, citiesWith1Alien2)

	if cities2[0].Name != "Bar" {
		t.Fatalf("Number of remaining cities differ, got: %s, want: %s", cities2[0].Name, "Bar")
	}
	if destroyedMap2["Foo"] != 1 {
		t.Fatalf("The city was not destroyed or not recorded, got: %d, want: %d", destroyedMap2["Foo"], 1)
	}
	if world2["Foo"] != 0 {
		t.Fatalf("The world map still has the destroyed city, got: %d, want: %d", world2["Foo"], 0)
	}
}
