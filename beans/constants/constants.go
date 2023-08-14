package constants

// log level
type LOG_LEVEL int

const (
	LOG_LEVEL_ERROR LOG_LEVEL = 1 + iota
	LOG_LEVEL_INFO
	LOG_LEVEL_DEBUG
)

const USER_TYPE_MASTER = "M"

const USER_TYPE_VISITOR = "V"
