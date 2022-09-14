package queue

import "time"

type PetNode struct {
	Name    string
	Arrival time.Time
}

type PetQueue struct {
	head, tail *PetNode
}
