package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func socket(w http.ResponseWriter, r *http.Request) {
	server := &Server{
		conn: make(map[*websocket.Conn]bool),
	}

	server.handleWebSocket(w, r)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	v, found := cache[req.URL.Path]
	mutex.RUnlock()
	if !found {
		mutex.Lock()
		defer mutex.Unlock()
		fileName := "./static" + req.URL.Path
		fmt.Print(fileName)
		f, err := os.Open(fileName)
		defer f.Close()
		if err != nil {
			http.NotFound(res, req)
			return
		}
		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		r := bytes.NewReader(b.Bytes())
		info, _ := f.Stat()
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}

	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
