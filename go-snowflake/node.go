package main

import (
	"errors"
	"sync"
	"time"
)

var (
	Epoch     int64 = 1288834974657
	NodeBits  uint8 = 10
	StepBits  uint8 = 12
	nodeMax   int64 = -1
	nodeMask        = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)

type Node struct {
	mu        sync.Mutex
	epoch     time.Time
	time      int64
	node      int64
	step      int64
	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

func NewSnowflakeNode(node int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("Remember, you have a total 22 bits to share between Node/Step")
	}

	n := Node{}

	return &n, nil
}

func (n *Node) GenerateID() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	// now := time.Since(n.epoch).Milliseconds()
}
