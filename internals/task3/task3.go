package internals_task3

import (
	"fmt"
	"sync"
	"time"
)

func ShowTask3() {
	var wg sync.WaitGroup
	var mt sync.Mutex
	personUse := make(chan string)

	person := []string{"Dragon", "Luffy", "Aokiji", "Akainu"}
	go func() {
		for _, user := range person {
			personUse <- user
		}
		close(personUse)
	}()

	wg.Add(len(person))
	for i := 0; i < len(person); i++ {
		go useMicrowave(personUse, &wg, &mt)
	}

	wg.Wait()
	println(">> Microwave dimatikan!")

}

func useMicrowave(personName chan string, wg *sync.WaitGroup, mt *sync.Mutex) {
	defer wg.Done()
	for name := range personName {
		fmt.Printf("%s sedang antre menggunakan microwave\n", name)
		mt.Lock()
		fmt.Printf(">> %s sedang menggunakan microwave\n", name)
		time.Sleep(2 * time.Second)
		fmt.Printf("Microwave telah selesai digunakan oleh %s\n", name)
		mt.Unlock()
	}
}
