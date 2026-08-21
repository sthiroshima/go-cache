package internal

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	network      string
	port         int
	storage      *Storage
	handler      *Handler
	ln           net.Listener
	connWg       sync.WaitGroup
	acceptWg     sync.WaitGroup
	ctx          context.Context
	mu           sync.Mutex
	connRegistry map[net.Conn]struct{}
}

func NewServer(network string, port int) *Server {
	storage := NewStorage()
	handler := NewHandler(storage)

	return &Server{
		network:      network,
		port:         port,
		storage:      storage,
		handler:      handler,
		connRegistry: make(map[net.Conn]struct{}),
	}
}

func (s *Server) Run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	s.ctx = ctx

	listener, err := net.Listen(s.network, fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	s.ln = listener

	s.acceptWg.Add(1)
	go s.acceptConnections()

	log.Println("Server is started.")

	<-ctx.Done()

	if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
		log.Println("Error when closing listener: ", err)
	}

	s.mu.Lock()
	for conn, _ := range s.connRegistry {
		if err := conn.Close(); err != nil {
			log.Println("Error close connection from connRegistry: ", err)
		}
	}
	s.mu.Unlock()

	s.acceptWg.Wait()
	s.connWg.Wait()

	log.Println("Server is stopped.")

	return nil
}

func (s *Server) acceptConnections() {
	defer s.acceptWg.Done()
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}

			log.Println("Error accepting connection:", err)
			continue
		}

		s.mu.Lock()
		s.connRegistry[conn] = struct{}{}
		s.mu.Unlock()

		select {
		case <-s.ctx.Done():
			conn.Close()
			return
		default:
			s.connWg.Add(1)
			go s.handleConnection(conn)
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer s.connWg.Done()
	defer conn.Close()

	defer func() {
		s.mu.Lock()
		delete(s.connRegistry, conn)
		s.mu.Unlock()
	}()

	reader := bufio.NewReader(conn)

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			cmd, err := ParseCommand(reader)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}

				log.Printf("Parse command error: %v", err)
				continue
			}

			response := s.handler.Handle(cmd)
			_, err = conn.Write([]byte(response))
			if err != nil {
				log.Printf("Server write error: %v", err)
				return
			}
		}
	}
}
