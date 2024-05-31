package engine

type key int

const (
	// ContextOriginalPath holds the original request URL
	ContextOriginalPath key = iota
	// ContextRequestStart hold the request start time
	ContextRequestStart
	// ContextOriginalPath hold the user ID (this is just for demo)
	ContextUserID
)
