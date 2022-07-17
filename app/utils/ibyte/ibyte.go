package ibyte

import (
	"bytes"
	"encoding/binary"
)

func Int64ToBytes(i int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, i)
	return buf[:n]
}

func BytesToInt64(buf []byte) (int64, error) {
	i, err := binary.ReadVarint(bytes.NewReader(buf))
	return i, err
}

func Int64ToBytes1(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt641(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
