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
	fmt.Println("Good morning!")

	var wG sync.WaitGroup

	getReady(bob, &wG)
	getReady(alice, &wG)
	wG.Wait()

	setAlarm()

	takeShoes(bob, &wG)
	takeShoes(alice, &wG)
	wG.Wait()

	fmt.Println("Exiting and locking the door")

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

func setAlarm() {
	fmt.Println("Set alarm")

	time.AfterFunc(60*time.Millisecond, func() {
		fmt.Println("Alarms rings")
	})
}

func rD(seed rand.Source, s, e int64) time.Duration {
	r := rand.New(seed)
	t := r.Int63n(e-s) + s
	d := time.Duration(t)
	return (d * time.Millisecond)
}
