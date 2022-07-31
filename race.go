package main

import (
	"fmt"
	"time"
)

type Person struct {
	follower int
}

// method for get follower
func (a *Person) getFollower() {
	a.follower += 1
}

// method for lose follower
func (a *Person) loseFollower() {
	if 0 > a.follower {
		fmt.Printf("already 0 follower: %d\n", a.follower)
	} else {
		time.Sleep(10 * time.Millisecond)
		a.follower -= 1
	}
}

func main() {
	var Person *Person = &Person{20}
	fmt.Printf("follower at the beginning: %d\n", Person.follower)

	for i := 1; i <= 10000; i++ {
		go losingFollower(i, Person)
		go gainFollower(i, Person)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("follower at the end: %d\n", Person.follower)
}

//lose function for go routines
func losingFollower(count int, a *Person) {
	fmt.Printf("%d: follower before: %d\n", count, a.follower)
	a.loseFollower()
	fmt.Printf("%d: follower after: %d\n", count, a.follower)
}

// gain function for go routines
func gainFollower(count int, a *Person) {
	a.getFollower()
	fmt.Printf("%d: follower after: %d\n", count, a.follower)
}
