package main

import (
	"fmt"
	"sync"
)

func insert(intchan chan int, wg *sync.WaitGroup) {

	for i := 1; i <= 100; i++ {
		intchan <- i
	}
	close(intchan)
	wg.Done()
}
func input(resChan chan int, intChan chan int, wg *sync.WaitGroup) {
	for i := 0; i < 8; i++ {
		for num := range intChan {
			sum := 0
			for j := 1; j <= num; j++ {
				sum += j
			}
			resChan <- sum
		}
	}
	close(resChan)
	for num := range resChan {
		fmt.Println(num)
	}

	wg.Done()
}

func out() {

}
func main() {
	intChan := make(chan int, 100)
	resChan := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go insert(intChan, &wg)
	go input(resChan, intChan, &wg)
	wg.Wait()
}
