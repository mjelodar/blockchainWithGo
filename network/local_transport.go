package network

import (
    "fmt"
    "sync"
)

type LocalTransport struct {
    addr NetAddr
    consumeCh chan RPC
    lock sync.Mutex
    peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
    return &LocalTransport{
        addr: addr,
        consumeCh: make(chan RPC, 1024), // Buffered channel for RPCs
        peers: make(map[NetAddr]*LocalTransport),
    }
}

func (lt *LocalTransport) Consume() <-chan RPC {
    return lt.consumeCh
}

func (lt *LocalTransport) Connect(tr Transport) error {
    peer, ok := tr.(*LocalTransport)
    if !ok {
        return fmt.Errorf("invalid transport type")
    }
    
    lt.lock.Lock()
    defer lt.lock.Unlock()

    if _, exists := lt.peers[peer.Addr()]; exists {
        return nil // Already connected
    }

    lt.peers[peer.Addr()] = peer
    peer.peers[lt.Addr()] = lt // Add this transport to the peer's list
    return nil
}

func (lt *LocalTransport) SendMessage(addr NetAddr, payload []byte) error {
    lt.lock.Lock()
    defer lt.lock.Unlock()

    peer, exists := lt.peers[addr]
    if !exists {
        return fmt.Errorf("no peer found with address %s", addr)
    }

    rpc := RPC{from: lt.Addr(), payload: payload}
    select {
    case peer.consumeCh <- rpc:
        return nil
    default:
        return fmt.Errorf("peer's consume channel is full")
    }
}

func (lt *LocalTransport) Addr() NetAddr {
    return lt.addr
}
