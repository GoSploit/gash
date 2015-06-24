package main

type Project struct {
	Hashes []*Hash
	Attack *Attack

	Keyspace uint64

	CurrentChunks []*Chunk
	LostChunks    []*Chunk

	NextKeyStart uint64
	KeyChunkSize uint64
}

func (p *Project) GetNextChunk() *Chunk {
	var chunk *Chunk
	if len(p.LostChunks) > 0 {
		chunk = p.LostChunks[0]
		p.LostChunks = p.LostChunks[1:]
		return chunk
	}
	if p.NextKeyStart >= p.Keyspace {
		return nil
	}
	chunk = NewChunk(p, p.NextKeyStart, p.KeyChunkSize)
	if chunk.KeyEnd > p.Keyspace {
		chunk.KeyEnd = p.Keyspace
	}
	p.NextKeyStart = p.KeyChunkSize
	p.CurrentChunks = append(p.CurrentChunks, chunk)
	return chunk
}
