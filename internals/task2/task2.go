package internals_task2

import (
	"fmt"
	"sync"
)

func ShowTask2() {
	var wg sync.WaitGroup
	ch := make(chan Messages)

	wg.Add(3)
	go sendMessages("Opet", "Dimana feb", ch, &wg)
	go sendMessages("radif", "Titip Nasi padang pep", ch, &wg)
	go sendMessages("ryan", "sini bandung pep!", ch, &wg)

	go whiteBoard(ch)

	wg.Wait()
	close(ch)

	fmt.Println("All messages have been displayed")

}

type Messages struct {
	Sender  string
	MsgDesc string
}

func sendMessages(sender string, message string, ch chan Messages, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- Messages{Sender: sender, MsgDesc: message}
}

func whiteBoard(msg chan Messages) {
	for messages := range msg {
		fmt.Printf("Sender : %s\nMessage: %s\n", messages.Sender, messages.MsgDesc)
	}
}
