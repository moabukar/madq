package disk

type Disk interface {
	ReadAt(b []byte, off int64) (int, error)
	WriteAt(b []byte, off int64) (int, error)
}

// Disk interface mainly for reading and writing data.
