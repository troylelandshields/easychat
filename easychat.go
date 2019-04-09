package easychat

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

type ChatClient struct {
	username         string
	outgoingMessages *json.Encoder
	incomingMessages *json.Decoder
}

func JoinChatRoom(ipAddr string, username string) *ChatClient {
	conn, err := net.Dial("tcp", ipAddr+":1234")
	if err != nil {
		log.Fatal("Error connecting to chat server:", err)
	}

	// send connection information
	enc := json.NewEncoder(conn)
	err = enc.Encode(username)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(conn)

	return &ChatClient{
		username:         username,
		outgoingMessages: enc,
		incomingMessages: dec,
	}
}

type ChatMessage struct {
	From string
	Body string
	Time time.Time
}

func (chat *ChatClient) SendMessage(messageText string) {
	msg := ChatMessage{
		From: chat.username,
		Body: messageText,
		Time: time.Now(),
	}

	err := chat.outgoingMessages.Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func (chat *ChatClient) ReceiveMessage() (ChatMessage, bool) {
	var incomingMsg ChatMessage

	err := chat.incomingMessages.Decode(&incomingMsg)
	if err != nil {
		return ChatMessage{}, false
	}

	return incomingMsg, true
}
