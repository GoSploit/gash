package main

import (
	"time"
)

type Chunk struct {
	Project   *Project
	KeyStart  uint64
	KeyEnd    uint64
	lossTimer *time.Timer
	worker    *Worker
}

func NewChunk(p *Project, ks uint64, ke uint64) *Chunk {
	chunk := &Chunk{
		Project:  p,
		KeyStart: ks,
		KeyEnd:   ke,
	}
	return chunk
}

func (c *Chunk) CatchLoss() {
	if c.lossTimer != nil {
		c.lossTimer.Stop()
	}
	c.lossTimer = time.AfterFunc(5*time.Minute, func() {
		c.worker.CurChunk = nil
		c.Project.LostChunks = append(c.Project.LostChunks, c)
		c.RemoveFromCurrent()
		c.worker = nil
	})
}

func (c *Chunk) ResetLoss() {
	if c.lossTimer != nil {
		c.lossTimer.Reset(5 * time.Minute)
	}
}

func (c *Chunk) RemoveFromCurrent() {
	found := -1
	for i, v := range c.Project.CurrentChunks {
		if v == c {
			found = i
			break
		}
	}
	if found >= 0 {
		c.Project.CurrentChunks = append(c.Project.CurrentChunks[0:found], c.Project.CurrentChunks[found+1:]...)
	}

}
