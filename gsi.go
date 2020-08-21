package gsi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Server struct {
	handlers []Handler
	addr     string
}

func NewServer(addr string) *Server {
	return &Server{
		handlers: make([]Handler, 0, 4),
		addr:     addr,
	}
}

func (s *Server) AddHandler(h Handler) {
	s.handlers = append(s.handlers, h)
}

func (s *Server) Listen() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var body struct {
			Provider Provider `json:"provider"`
		}
		json.Unmarshal(buf, &body)
		for _, h := range s.handlers {
			for _, id := range h.AppID() {
				if id == body.Provider.AppID {
					h.Receive(buf)
				}
			}
		}
	})
	server := &http.Server{Addr: s.addr, Handler: mux}
	return server.ListenAndServe()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf := bytes.NewBuffer(make([]byte, 0, 4096))
		io.Copy(buf, r.Body)
		defer r.Body.Close()
		fmt.Println(buf.String())
	})
	http.ListenAndServe(":58591", nil)
}
