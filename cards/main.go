package main

import (
	"fmt"

	"github.com/jeffreylowy/intro-to-go/henlo"
	"github.com/jeffreylowy/intro-to-go/loops"
)

func main() {
	//randNum := generateRandomNumber(generateSeedValue(), 100)
	askQuestion()

	kals := henlo.Doggo{
		Name:  "Kali",
		Color: "gray and white",
		Age:   8,
	}

	k := fmt.Sprintf("My dog, %s, is %s. She is %v years old", kals.PrintDoggosName(), kals.Color, kals.Age)
	fmt.Println(k)
	fmt.Println("Chico is", henlo.Chico.Age)

	// kals.Age = 1
	if kals.Age > 7 {
		fmt.Println("Kali is", kals.Age, "years old")
	} else {
		fmt.Println("Are you sure Kals is younger than 7?")
	}

	l := loops.Looper{
		MaxCount:     10,
		CurrentCount: 0,
	}

	l.Loop()
}
