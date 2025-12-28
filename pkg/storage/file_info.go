package storage

import "time"

// FileInfo represents metadata about a stored file
type FileInfo struct {
	Key          string    // Unique identifier/path of the file
	Size         int64     // Size in bytes
	LastModified time.Time // Last modification time
	ContentType  string    // MIME type
}
