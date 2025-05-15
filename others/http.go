package others

import (
	"fmt"
	"net/http"
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

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value, ok := store.Get(key)
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, value)
	})

	http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		store.Put(key, value)
		fmt.Fprint(w, "OK")
	})
	http.ListenAndServe(":8080", nil)
}
