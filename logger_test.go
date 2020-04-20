package gwlog

import (
	"bytes"
	"strings"
	"testing"
)

func Test_logger_Type(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := GetLogger()
	logger.SetOutput(buf)
	logger.Type("APP").Info("example")
	if strings.Contains(buf.String(), "type=APP") == false {
		t.Errorf("invalid format: %s", buf.String())
	}
}
