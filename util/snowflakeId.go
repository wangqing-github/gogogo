package util

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	// 开始时间戳（自定义，通常设置为系统上线的时间戳）
	epoch = int64(1609459200000) // 2021-01-01 00:00:00 UTC

	// 机器ID所占的位数
	workerIDBits = 5
	// 数据中心ID所占的位数
	datacenterIDBits = 5
	// 支持的最大机器ID数量，结果是31 (这个移位算法可以很快地计算出几位二进制数所能表示的最大十进制数)
	maxWorkerID = -1 ^ (-1 << workerIDBits)
	// 支持的最大数据中心ID数量，结果是31
	maxDatacenterID = -1 ^ (-1 << datacenterIDBits)
	// 序列在ID中占的位数
	sequenceBits = 12
	// 机器ID左移12位
	workerIDShift = sequenceBits
	// 数据中心ID左移17位(12+5)
	datacenterIDShift = sequenceBits + workerIDBits
	// 时间戳左移22位(5+5+12)
	timestampLeftShift = sequenceBits + workerIDBits + datacenterIDBits
	// 生成序列的掩码，这里为4095 (0b111111111111=0xfff=4095)
	sequenceMask = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	datacenterID  int64
	sequence      int64
}

var sf = NewSnowflake(1, 1)

// NewSnowflake 创建一个新的雪花ID生成器
func NewSnowflake(workerID, datacenterID int64) *Snowflake {
	if workerID < 0 || workerID > maxWorkerID {
		return nil
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil
	}
	return &Snowflake{
		lastTimestamp: -1,
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
	}
}

// NextID 生成下一个ID
func (sf *Snowflake) NextID() (int64, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if timestamp < sf.lastTimestamp {
		return 0, errors.New("clock moved backwards. Refusing to generate id")
	}

	if sf.lastTimestamp == timestamp {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			for timestamp <= sf.lastTimestamp {
				timestamp = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampLeftShift) |
		(sf.datacenterID << datacenterIDShift) |
		(sf.workerID << workerIDShift) |
		sf.sequence

	return id, nil
}

func NewSnowflakeId() int64 {
	id, err := sf.NextID()
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return -1
	}
	return id
}
