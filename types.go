package main

import (
	"io"
	"sync"
	"time"
)

type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile
var mutex = new(sync.RWMutex)
