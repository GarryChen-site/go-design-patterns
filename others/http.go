package others

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Store struct {
	dataChan chan dataRequest
}

type dataRequest struct {
	op    string
	key   string
	value string
	resp  chan dataResponse
}

type dataResponse struct {
	value string
	ok    bool
}

func NewStore() *Store {
	s := &Store{
		dataChan: make(chan dataRequest),
	}
	go s.run()
	return s
}

func (s *Store) run() {
	data := make(map[string]string)
	for req := range s.dataChan {
		switch req.op {
		case "get":
			value, ok := data[req.key]
			req.resp <- dataResponse{value: value, ok: ok}
		case "put":
			data[req.key] = req.value
			req.resp <- dataResponse{ok: true}
		}
	}
}

func (s *Store) Get(key string) (string, bool) {
	respChan := make(chan dataResponse)
	s.dataChan <- dataRequest{op: "get", key: key, resp: respChan}
	resp := <-respChan
	return resp.value, resp.ok
}

func (s *Store) Put(key, value string) {
	respChan := make(chan dataResponse)
	s.dataChan <- dataRequest{op: "put", key: key, value: value, resp: respChan}
	<-respChan
}

func main() {
	store := NewStore()

	mux := http.NewServeMux()

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value, ok := store.Get(key)
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, value)
	})

	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		store.Put(key, value)
		fmt.Fprint(w, "OK")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe error: %v\n", err)
		}
	}()

	fmt.Println("Server is running on port 8080...")
	<-stop // Wait for interrupt signal
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown error: %v\n", err)
	}
	fmt.Println("Server stopped gracefully.")

}
