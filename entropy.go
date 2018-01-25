package main

import (
	"math/rand"
	"time"
)

var entropy = rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano()))
