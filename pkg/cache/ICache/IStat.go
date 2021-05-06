package ICache

import "log"

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) Add(k string, v []byte) {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
	log.Println("缓存写入", &s)
}

func (s *Stat) Remove(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}