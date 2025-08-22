package internals_task2

import (
	"fmt"
	"sync"
)

func ShowTask2() {
	var wg sync.WaitGroup
	ch := make(chan Messages)

	allMessages := []*Messages{
		NewMessages("Dana", "dimana bro"),
		NewMessages("Ara", "dimana bro"),
		NewMessages("Genji", "POSISI bro"),
	}

	go whiteBoard(ch)

	wg.Add(len(allMessages))

	for _, msg := range allMessages {
		go sendMessages(msg, ch, &wg)
	}

	wg.Wait()
	close(ch)

	fmt.Println("All messages have been displayed")

}

type Messages struct {
	Sender  string
	MsgDesc string
}

func NewMessages(sender, msg string) *Messages {
	return &Messages{
		Sender:  sender,
		MsgDesc: msg,
	}
}

func sendMessages(newMsg *Messages, ch chan Messages, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- Messages{Sender: newMsg.Sender, MsgDesc: newMsg.MsgDesc}
}

func whiteBoard(msg chan Messages) {
	for messages := range msg {
		fmt.Println("=================New Messages=================")
		fmt.Printf("Sender : %s\nMessage: %s\n", messages.Sender, messages.MsgDesc)
	}
}
