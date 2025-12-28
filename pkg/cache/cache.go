package cache

import (
	"sync"
	"time"
)

type CacheStructure struct {
	Value      interface{}
	Expiration time.Time
}

type Cache struct {
	store sync.Map
}
