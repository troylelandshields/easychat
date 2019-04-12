package easychat

import "fmt"

func ExampleJoinChatRoom() {
	chatClient, err := JoinChatRoom("128.0.0.1", "ExampleUser")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	msg := chatClient.ReceiveMessage()
	fmt.Println("Message from", msg.From, msg.Body)

	chatClient.SendMessage("Hi!")
}
