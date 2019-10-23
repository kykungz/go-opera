package littleendian

import "encoding/binary"

// Int64ToBytes converts uint64 to bytes.
func Int64ToBytes(n uint64) []byte {
	var res [8]byte
	binary.LittleEndian.PutUint64(res[:], n)
	return res[:]
}

// BytesToInt64 converts uint64 from bytes.
func BytesToInt64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

// Int32ToBytes converts uint32 to bytes.
func Int32ToBytes(n uint32) []byte {
	var res [4]byte
	binary.LittleEndian.PutUint32(res[:], n)
	return res[:]
}

// BytesToInt32 converts uint32 from bytes.
func BytesToInt32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
