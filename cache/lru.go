package cache

import "container/list"

type Value interface {
	Len() int
}

type Cache struct {
	maxBytes  int64
	useBytes  int8
	ll        *list.List
	cache     map[string]list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}
