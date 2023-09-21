package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the function returns
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func hell(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the function returns
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup    // Create a WaitGroup object
	wg.Add(2)                // Set the counter to 2, the number of goroutines to wait for
	go say("world", &wg)     // Pass the address of the WaitGroup object to the function
	go hell("children", &wg) // Pass the address of the WaitGroup object to the function
	wg.Wait()                // Block the main function until the counter becomes zero

}
