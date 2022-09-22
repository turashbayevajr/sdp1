package main

import (
	"awesomeProject2/ch"
	"fmt"
)

func main() {
	f1 := &ch.Follower{Name: "Follower1"}
	f2 := &ch.Follower{Name: "Follower2"}
	f3 := &ch.Follower{Name: "Follower3"}
	fN := &ch.Follower{Name: "FollowerN"}
	channel := ch.Channel{
		Name:      "channel",
		Followers: map[string]ch.Observer{},
	}
	channel.Subcribe(f1)
	channel.Subcribe(f2)
	channel.Subcribe(f3)
	channel.Subcribe(fN)
	channel.SendAll()
	fmt.Printf("UnSubcribe follower FollowerN ")
	channel.UnSubcribe(fN)
	channel.SendAll()
}
