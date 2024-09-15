package fileio

// IOManager:
// Generic IO Manager Interface
type IOManager interface {
	// Read: read from disk
	Read([]byte, int64) (int, error)

	// Write: write to the disk
	Write([]byte) (int, error)

	// Sync: commits the current contents of the file to stable storage.
	Sync() error

	// Close: close the file
	Close() error
}
