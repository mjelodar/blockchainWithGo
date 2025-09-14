package main

import (
    "github.com/PRACTICING-GO/blockchain/network"
    "time"
)

func main() {
    trLocal := network.NewLocalTransport("LOCAL")
    trRemote := network.NewLocalTransport("REMOTE")

    trLocal.Connect(trRemote)
    trRemote.Connect(trLocal)

    go func(){
        for {
            trRemote.SendMessage(trLocal.Addr(), []byte("Hello from Local to Remote!"))
            time.Sleep(2 * time.Second)
        }
    }()

    ops := network.ServerOps{
        Transports: []network.Transport{trLocal},
    }

    s:= network.NewServer(ops)
    s.Start()
}