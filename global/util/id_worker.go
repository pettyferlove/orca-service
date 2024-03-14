package util

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type IdWorker struct {
	startTime             int64
	workerIdBits          uint
	datacenterIdBits      uint
	maxWorkerId           int64
	maxDatacenterId       int64
	sequenceBits          uint
	workerIdLeftShift     uint
	datacenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	datacenterId          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	idLock                *sync.Mutex
}

func (worker *IdWorker) InitIdWorker(workerId, datacenterId int64) error {

	var baseValue int64 = -1
	worker.startTime = 1463834116272
	worker.workerIdBits = 5
	worker.datacenterIdBits = 5
	worker.maxWorkerId = baseValue ^ (baseValue << worker.workerIdBits)
	worker.maxDatacenterId = baseValue ^ (baseValue << worker.datacenterIdBits)
	worker.sequenceBits = 12
	worker.workerIdLeftShift = worker.sequenceBits
	worker.datacenterIdLeftShift = worker.workerIdBits + worker.workerIdLeftShift
	worker.timestampLeftShift = worker.datacenterIdBits + worker.datacenterIdLeftShift
	worker.sequenceMask = baseValue ^ (baseValue << worker.sequenceBits)
	worker.sequence = 0
	worker.lastTimestamp = -1
	worker.signMask = ^baseValue + 1

	worker.idLock = &sync.Mutex{}

	if worker.workerId < 0 || worker.workerId > worker.maxWorkerId {
		return errors.New(fmt.Sprintf("workerId[%v] is less than 0 or greater than maxWorkerId[%v].", workerId, datacenterId))
	}
	if worker.datacenterId < 0 || worker.datacenterId > worker.maxDatacenterId {
		return errors.New(fmt.Sprintf("datacenterId[%d] is less than 0 or greater than maxDatacenterId[%d].", workerId, datacenterId))
	}
	worker.workerId = workerId
	worker.datacenterId = datacenterId
	return nil
}

func (worker *IdWorker) NextId() (int64, error) {
	worker.idLock.Lock()
	timestamp := time.Now().UnixNano()
	if timestamp < worker.lastTimestamp {
		return -1, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", worker.lastTimestamp-timestamp))
	}

	if timestamp == worker.lastTimestamp {
		worker.sequence = (worker.sequence + 1) & worker.sequenceMask
		if worker.sequence == 0 {
			timestamp = worker.tilNextMillis()
			worker.sequence = 0
		}
	} else {
		worker.sequence = 0
	}

	worker.lastTimestamp = timestamp

	worker.idLock.Unlock()

	id := ((timestamp - worker.startTime) << worker.timestampLeftShift) |
		(worker.datacenterId << worker.datacenterIdLeftShift) |
		(worker.workerId << worker.workerIdLeftShift) |
		worker.sequence

	if id < 0 {
		id = -id
	}

	return id, nil
}

func (worker *IdWorker) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano()
	if timestamp <= worker.lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}
