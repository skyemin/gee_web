package main

import "sync"

func main() {
	s := sync.Map{}
	s.Store("a", "1")
	s.Store("b", 2)
	value, ok := s.Load("a")
	if ok {
		print(value.(string))
	}
}
