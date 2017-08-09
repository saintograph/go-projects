package main

type Petstore struct {
	Name string
	Location string
	Dogs []*Pet
	Cats []*Pet
}

type Pet struct {
	Name string
	Breed string
}