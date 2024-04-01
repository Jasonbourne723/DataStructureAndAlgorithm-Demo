package snowflakeid

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type SnowFlaker struct {
	locker sync.Mutex
	//lastTimestamp int64
	dataCenterId int64
	workId       int64
	seq          int64
}

func (s *SnowFlaker) Next() int64 {
	defer s.locker.Unlock()
	s.locker.Lock()

	timestamp := timeGen()

	fmt.Println(timestamp)
	fmt.Println(strconv.FormatInt(timestamp, 2))
	s.dataCenterId = 0
	s.workId = 1
	s.seq = 1

	return 0<<63 | timestamp<<22 | s.dataCenterId<<17 | s.workId<<12 | s.seq
}

func timeGen() int64 {
	return int64(time.Since(time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)).Milliseconds())
}
