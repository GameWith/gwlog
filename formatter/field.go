package formatter

import "time"

const (
	defaultTimestampFormat = time.RFC3339
	// FieldKeyMsg is Default message field key
	FieldKeyMsg = "message"
	// FieldKeyLevel is Default level field key
	FieldKeyLevel = "level"
	// FieldKeyTime is Default time field key
	FieldKeyTime = "time"
)

// FieldMap is Default field key map
type FieldMap map[string]string

func (f FieldMap) get(key string) string {
	if k, ok := f[key]; ok {
		return k
	}
	return string(key)
}

// Fields is custom field map
type Fields map[string]interface{}

func (f Fields) conflict(key string, fm FieldMap) bool {
	fmKey := fm.get(key)
	return f.has(fmKey)
}

func (f Fields) has(key string) bool {
	_, ok := f[key]
	return ok
}
