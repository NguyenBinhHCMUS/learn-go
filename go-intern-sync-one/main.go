package main

import (
	"fmt"
	"sync"
	"time"
)

type Singleton struct {}

var (
	instance *Singleton
	once     sync.Once
)

// Without sync one
// func GetInstance() *Singleton {
// 	if instance == nil {
// 		fmt.Println("Creating new instance")
// 		instance = &Singleton{}
// 	}
// 	return instance
// }

//  With sync one
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating new instance")
		instance = &Singleton{}
	})
	return instance
}

func main() {
	// s1 := GetInstance()
	// fmt.Printf("Instance 1: %p\n", s1)

	// time.Sleep(2 * time.Second)

	// s2 := GetInstance()
	// fmt.Printf("Instance 2: %p\n", s2)

	time.Sleep(2 * time.Second)

	// High traffic simulation
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := GetInstance()
			fmt.Printf("Goroutine Instance: %p\n", s)
		}()
	}
	wg.Wait()
}