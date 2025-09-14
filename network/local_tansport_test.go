package network

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T){
    tra := NewLocalTransport("A")
    trb := NewLocalTransport("B")

    tra.Connect(trb)
    trb.Connect(tra)

    assert.Equal(t, tra.peers[trb.addr], trb)
    assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
    tra := NewLocalTransport("A")
    trb := NewLocalTransport("B")

    err := tra.Connect(trb)
    assert.NoError(t, err)

    payload := []byte("Hello, B!")
    err = tra.SendMessage(trb.Addr(), payload)
    assert.NoError(t, err)

    select {
    case rpc := <-trb.Consume():
        assert.Equal(t, rpc.from, tra.Addr())
        assert.Equal(t, rpc.payload, payload)
    default:
        t.Fatal("Expected to receive a message but got none")
    }
}