package network

import (
    "fmt"
    "time"
)

type ServerOps struct{
    Transports []Transport
}

type Server struct {
    ServerOps
    rpcCh chan RPC // Channel to handle incoming RPCs
    quitCh chan struct{} // Channel to signal server shutdown
}

func NewServer(ops ServerOps) *Server {
    return &Server{
        ServerOps: ops,
        rpcCh: make(chan RPC, 1024), // Buffered channel for RPCs
        quitCh: make(chan struct{}, 1),
    }
}

func (s *Server) Start() { 
    s.initTransports()
    ticker := time.NewTicker(5 * time.Second)
free:
    for {
        select {
        case rpc := <-s.rpcCh:
            // Handle incoming RPCs
            fmt.Println("Received RPC from:", rpc.from, "with payload:", string(rpc.payload))
        case <-s.quitCh:
            break free // Shutdown the server
        case <-ticker.C:
            fmt.Println("do some stuff every 5 seconds")
        }
    }
    fmt.Println("Server shutting down...")
}

func (s *Server) initTransports() {
    for _, tr := range s.Transports {
        go func(tr Transport) {
            for rpc := range tr.Consume() {
                s.rpcCh <- rpc
            }
        }(tr)
    }
}