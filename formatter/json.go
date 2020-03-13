package formatter

import (
	"github.com/sirupsen/logrus"

	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONFormatter is custom json formatter for logrus
type JSONFormatter struct {
	FieldMap FieldMap
}

// Format is custom json format
//
// Setup:
// 	gwlog.GetLogger().SetFormatter(&formatter.JSONFormatter{})
//
// Options:
//	// Change default field name.
// 	gwlog.GetLogger().SetFormatter(&formatter.JSONFormatter{
//		FieldMap: FieldMap{
// 			FieldKeyMsg:   "@message",
//			FieldKeyLevel: "@level",
//		}
// 	})
// Notice:
// 	If FieldName of FieldMap and FieldName of WithFields are the same, WithFields takes precedence.
func (j *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	// Set level
	if !data.conflict(FieldKeyLevel, j.FieldMap) {
		data[j.FieldMap.get(FieldKeyLevel)] = strings.ToUpper(entry.Level.String())
	}

	// Set Message
	if !data.conflict(FieldKeyMsg, j.FieldMap) {
		data[j.FieldMap.get(FieldKeyMsg)] = entry.Message
	}

	// Set Time
	if !data.conflict(FieldKeyTime, j.FieldMap) {
		data[j.FieldMap.get(FieldKeyTime)] = entry.Time.Format(defaultTimestampFormat)
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
