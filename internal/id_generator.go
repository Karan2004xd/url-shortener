package internal

import (
	"sync"
	"time"
)

const (
	customEpoch = 1704067200000
	maxSequence = 4095
	dataCenterId = 0
	machineId = 0
)

var (
	lastTimeStamp int64
	sequence int64
	mutex sync.Mutex
)

func setSequence() {
	mutex.Lock()
	defer mutex.Unlock()

	currentTime := time.Now().UnixMilli() - customEpoch

	if currentTime == lastTimeStamp {
		sequence++

		if sequence > maxSequence {
			for currentTime <= lastTimeStamp {
				currentTime = time.Now().UnixMilli() - customEpoch
			}
			sequence = 0
		}
	} else {
		sequence = 0
	}

	lastTimeStamp = currentTime
}

func GenerateId() int64 {
	setSequence()

	currentTime := time.Now().UnixMilli() - customEpoch
	id := (currentTime << 22) | (dataCenterId << 17) | (machineId << 12) | sequence

	return id
}
