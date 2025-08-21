package internals_task1

import (
	"fmt"
	"sync"
	"time"
)

func ShowTask1() {
	var wg sync.WaitGroup

	wg.Add(1)
	go wakeUp(&wg)
	wg.Wait()
	wg.Add(1)
	go tidyUpBed(&wg)
	wg.Wait()
	wg.Add(1)
	go takeABath(&wg)
	wg.Wait()
	wg.Add(1)
	go breakfast(&wg)

	wg.Wait()
	fmt.Println("Time to works !")
}

func wakeUp(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Alarm its ringing ... ")
	time.Sleep(2 * time.Second)
	println("Turn off alarm")
	time.Sleep(500 * time.Millisecond)
	println("Finish turn off alarm")
}

func tidyUpBed(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Merapikan tempat tidur...")
	time.Sleep(2 * time.Second)
	println("Selesai Merapikan tempat tidur")
}

func takeABath(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Mandi...")
	time.Sleep(2 * time.Second)
	println("Selesai mandi")
}

func breakfast(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Prepare the breakfast...")
	time.Sleep(2 * time.Second)
	println("Finished breakfast")
}
