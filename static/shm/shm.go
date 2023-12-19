package main

import (
	"fmt"
	"github.com/hidez8891/shm"
)

func main() {
	w, _ := shm.Create("shm_name", 256)
	defer w.Close()

	r, _ := shm.Open("shm_name", 256)
	defer r.Close()

	wbuf := []byte("Hello World")
	w.Write(wbuf)

	rbuf := make([]byte, len(wbuf))
	r.Read(rbuf)
	// rbuf == wbuf

	fmt.Println(string(rbuf))
}
