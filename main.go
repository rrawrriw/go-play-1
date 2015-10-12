package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	bob   = "Bob"
	alice = "Alice"
)

func main() {
	var wG sync.WaitGroup

	fmt.Println("Good morning!")

	getReady(bob, &wG)
	getReady(alice, &wG)

	// Block until bob and alice are ready to rumble.
	wG.Wait()

	alarm := setAlarm()
	takeShoes(bob, &wG)
	takeShoes(alice, &wG)

	// Start goroutine to wait to finish putting on shoes.
	// No blocking, due to the alarm has the possiblity to ring
	// befor the shoes have been putted on.
	waitShoes(&wG)

	// Block until the alarm chan receive a signal.
	waitAlarm(alarm)

}

func getReady(name string, wG *sync.WaitGroup) {
	seed := rand.NewSource(int64(time.Now().Nanosecond()))
	d := rD(seed, 60, 90)

	wG.Add(1)

	fmt.Println(name, "start to get ready")
	time.AfterFunc(d, func() {
		defer wG.Done()
		fmt.Println(name, "spent", d, "seconds to get ready")
	})
}

func takeShoes(name string, wG *sync.WaitGroup) {
	seed := rand.NewSource(int64(time.Now().Nanosecond()))
	d := rD(seed, 35, 45)

	wG.Add(1)

	fmt.Println(name, "started putting on shoes")
	time.AfterFunc(d, func() {
		defer wG.Done()
		fmt.Println(name, "spent", d, "putting on shoes")
	})
}

func waitShoes(wG *sync.WaitGroup) {
	go func() {
		wG.Wait()
		fmt.Println("Exiting and locking the door.")
	}()
}

func setAlarm() <-chan struct{} {
	fmt.Println("Set alarm")
	done := make(chan struct{})

	time.AfterFunc(60*time.Millisecond, func() {
		fmt.Println("Alarms rings")
		done <- struct{}{}
	})

	return done
}

func waitAlarm(d <-chan struct{}) {
	<-d
}

func rD(seed rand.Source, s, e int64) time.Duration {
	r := rand.New(seed)
	t := r.Int63n(e-s) + s
	d := time.Duration(t)
	return (d * time.Millisecond)
}
