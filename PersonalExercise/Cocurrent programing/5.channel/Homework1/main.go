package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	personChannel := make(chan Person)
	go func() {
		for i := 0; i < 10; i++ {
			person := Person{
				Name:    generateRandomString(5),
				Age:     generateRandomAge(18, 60),
				Address: generateRandomString(10),
			}
			personChannel <- person
		}
		close(personChannel)
	}()

	for person := range personChannel {
		fmt.Printf("Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)
	}
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}

func generateRandomAge(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
