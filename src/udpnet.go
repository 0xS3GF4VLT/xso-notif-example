package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type XSOConnect struct {
	conn net.Conn
}

type XSONotif struct {
	Title, //Notification title, supports Rich Text Formatting
	AudioPath, //File path to .ogg audio file. Can be "default", "error", or "warning". Notification will be silent if left empty.
	Content, //Notification content, supports Rich Text Formatting, if left empty, notification will be small.
	Icon, //Base64 Encoded image, or file path to image. Can also be "default", "error", or "warning"
	SourceApp string //Somewhere to put your app name for debugging purposes
	MessageType, // 1 = Notification Popup, 2 = MediaPlayer Information, will be extended later on.
	Index int //Only used for Media Player, changes the icon on the wrist.
	Timeout, //How long the notification will stay on screen for in seconds
	Height, //Height notification will expand to if it has content other than a title. Default is 175
	Volume float32 // Notification sound volume.
	UseBase64Icon bool //Set to true if using Base64 for the icon image
}

func (x *XSOConnect) Connect() {
	con, err := net.Dial("udp", "127.0.0.1:42069")

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	x.conn = con
}

func (x *XSOConnect) SendNotif(msg string, time float32) {
	if msg == "" {
		msg = "Hello World"
	}

	if time == 0 {
		time = 1
	}

	notif := XSONotif{
		AudioPath:     "default",
		Title:         "XSONotif",
		Content:       msg,
		Icon:          "default",
		SourceApp:     "XSONotif",
		MessageType:   1,
		Index:         0,
		Timeout:       time,
		Height:        175,
		Volume:        0.7,
		UseBase64Icon: false,
	}

	obj, err := json.Marshal(notif)

	if err != nil {
		fmt.Printf("JSON Marshalling error: %s\n", err)
	}

	x.conn.Write(obj)
}
