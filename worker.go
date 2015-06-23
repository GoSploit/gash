package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
)

type Worker struct {
	ID       string
	CurChunk *Chunk
}

func (w *Worker) GetChunk() *Chunk {
	var chunk *Chunk
	for _, p := range Projects {
		chunk = p.GetNextChunk()
		if chunk != nil {
			break
		}
	}
	w.CurChunk = chunk
	return chunk
}

type WorkerMap struct {
	workermap map[string]*Worker
	lock      sync.Mutex
}

func (wm *WorkerMap) Get(id string) *Worker {
	wm.lock.Lock()
	defer wm.lock.Unlock()
	return wm.workermap[id]
}

func (wm *WorkerMap) Add(id string) *Worker {
	wm.lock.Lock()
	defer wm.lock.Unlock()
	bid := make([]byte, 20)
	rand.Read(bid)
	id := base64.StdEncoding.EncodeToString(bid)
	w := &Worker{
		ID: id,
	}
	wm.workermap[id] = w
	return w
}

func GetWork(rw http.ResponseWriter, req *http.Request) {
	workerID := req.FormValue("workerID")
	for _, w := range Workers {

	}
}
