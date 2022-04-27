package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age int
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}
	sort.SliceStable(people, func(j, i int) bool { return people[j].first > people[i].first })
	fmt.Println(people)
	sort.SliceStable(people, func(j, i int) bool { return people[j].age > people[i].age })
	fmt.Println(people)

}