package structure

import ()

type Alien struct {
	IsTrapped    bool
	IsDead       bool
	IsDoneMoving bool
	Index        int
	Moves        int
}

type City struct {
	IsDestroyed bool
	Name        string
	North       string
	East        string
	South       string
	West        string
	Aliens      []Alien
}
