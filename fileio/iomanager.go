package fileio

// IOManager:
// Generic IO Manager Interface
type IOManager interface {
	// Read: read from disk
	Read([]byte, int64) (int, error)

	// Write: write to the disk
	Write([]byte) (int, error)

	// Sync
	Sync() error

	// Close: close the file
	Close() error
}
