package easychat

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"time"
)

// ChatClient has functionality for sending and receiving messages from an EasyChatServer
type ChatClient struct {
	username         string
	outgoingMessages *json.Encoder
	incomingMessages *json.Decoder
}

// JoinChatRoom requires the IP Address of an EasyChatServer and a username to use for connecting.
// It returns a ChatClient or an error if we can't connect.
//  chatClient, err := JoinChatRoom("128.0.0.1", "ExampleUser")
//  if err != nil {
//	   fmt.Println("Error occurred:", err.Error())
//     return
//  }

func JoinChatRoom(ipAddr string, username string) (*ChatClient, error) {
	conn, err := net.Dial("tcp", ipAddr+":1234")
	if err != nil {
		return nil, errors.New("Error connecting to chat server: " + err.Error())
	}

	// send connection information
	enc := json.NewEncoder(conn)
	err = enc.Encode(username)
	if err != nil {
		return nil, errors.New("Error sending username to chatroom: " + err.Error())
	}

	dec := json.NewDecoder(conn)

	return &ChatClient{
		username:         username,
		outgoingMessages: enc,
		incomingMessages: dec,
	}, nil
}

// ChatMessage contains a message, the author, and the timestamp of the message.
type ChatMessage struct {
	From string
	Body string
	Time time.Time
}

// SendMessage will deliver the messageText in a ChatMessage to the EasyChatServer so it can be sent to everyone in the chatroom
//  chatClient.SendMessage("Hi! I'm sending you a message")
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

// ReceiveMessage will return the next message that has been sent in the chatroom from the EasyChatServer.
//  msg := chatClient.ReceiveMessage()
//  fmt.Println(msg.Body)
func (chat *ChatClient) ReceiveMessage() ChatMessage {
	var incomingMsg ChatMessage

	_ = chat.incomingMessages.Decode(&incomingMsg)

	return incomingMsg
}
