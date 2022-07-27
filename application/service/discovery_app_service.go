package service

import (
	ps "github.com/Rock-liyi/p2pdb-pubsub"
)

func InitSub() {

	var sub ps.PubSub
	sub.SetType("p2pdb")
	var subscription, err = sub.Sub()
	if err != nil {
		panic(err)
	}
	go sub.StartNewSubscribeService(subscription)
}

func InitPub() {

	var pub ps.PubSub
	pub.SetType("p2pdb")
	go pub.InitPub()
}