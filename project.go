package main

type Project struct {
	Hashes []*Hash
	Attack *Attack

	Keyspace uint64

	CurrentChunks *[]Chunk
	LostChunks    *[]Chunk

	NextKeyStart uint64
	KeyChunkSize uint64
}

func (p *Project) GetNextChunk() *Chunk {
	if len(p.LostChunks) > 0 {
		chunk := p.LostChunks[0]
		h.LostChunks = p.LostChunks[1:]
		return chunk
	}
	if p.NextKeyStart >= p.Keyspace {
		return nil
	}
	chunk = &Chunk{
		Hash:     p,
		KeyStart: p.NextKeyStart,
		KeyEnd:   p.KeyChunkSize,
	}
	if chunk.KeyEnd > p.Keyspace {
		chunk.KeyEnd = p.Keyspace
	}
	p.NextKeyStart = p.KeyChunkSize
	return chunk
}
