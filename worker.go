package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sync"
)

type Worker struct {
	ID       string
	CurChunk *Chunk
}

func (w *Worker) GetChunk() *Chunk {
	if w.CurChunk != nil {
		return w.CurChunk
	}
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

func (w *Worker) FinishChunk() {
	if w.CurChunk != nil {
		w.CurChunk.lossTimer.Stop()
		w.CurChunk.RemoveFromCurrent()
		w.CurChunk = nil
	}
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

func (wm *WorkerMap) Add() *Worker {
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
	var chunk *Chunk
	workerID := req.FormValue("workerID")
	worker := Workers.Get(workerID)
	rw.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(rw)
	if worker != nil {
		chunk = worker.GetChunk()
		enc.Encode(chunk)
	} else {
		enc.Encode(struct {
			Error string
		}{
			Error: "No worker found by that ID",
		})
	}
}

func Register(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(rw)
	if req.FormValue("token") == "swordfish" {
		w := Workers.Add()
		enc.Encode(w)
	} else {
		enc.Encode(struct {
			Error string
		}{
			Error: "Invalid Token",
		})
	}
}

func FinishWork(rw http.ResponseWriter, req *http.Request) {
	var chunk *Chunk
	workerID := req.FormValue("workerID")
	worker := Workers.Get(workerID)
	rw.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(rw)
	if worker != nil {
		worker.FinishChunk()
		chunk = worker.GetChunk()
		enc.Encode(chunk)
	} else {
		enc.Encode(struct {
			Error string
		}{
			Error: "No worker found by that ID",
		})
	}
}
