package data

// LogRecordPos used to locate the log record
type LogRecordPos struct {
	Fid    uint32 // File id, used to represent the file in which the record is saved
	Offset int64  // Offset, used to denote the position of the record in the file
}
