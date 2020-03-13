package formatter

import (
	"github.com/GameWith/gwlog"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func mockEntry(level logrus.Level, message string, fields *logrus.Fields) *logrus.Entry {
	logger := logrus.New()
	entry := logrus.NewEntry(logger)
	entry.Level = level
	entry.Message = message
	jst, _ := time.LoadLocation("Asia/Tokyo")
	entry.Time = time.Date(2000, 1, 1, 0, 0, 0, 0, jst)
	if fields != nil {
		entry.Data = *fields
	}
	return entry
}

func ExampleJSONFormatter_Format() {
	// singleton instance
	gwlog.GetLogger().SetOutput(os.Stdout)
	gwlog.GetLogger().SetFormatter(&JSONFormatter{})
	gwlog.GetLogger().WithFields(map[string]interface{}{
		FieldKeyTime: "2000-01-01T00:00:00+09:00", // fixed time
	}).Info("aaa")
	// Output: {"level":"INFO","message":"aaa","time":"2000-01-01T00:00:00+09:00"}
}

func ExampleJSONFormatter_Format_withFields() {
	logger := gwlog.GetLogger()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&JSONFormatter{})
	logger.WithFields(map[string]interface{}{
		FieldKeyTime: "2000-01-01T00:00:00+09:00", // fixed time
		"hoge":       "hoge",
	}).Info("aaa")
	// Output: {"hoge":"hoge","level":"INFO","message":"aaa","time":"2000-01-01T00:00:00+09:00"}
}

func ExampleJSONFormatter_Format_changeDefaultField() {
	logger := gwlog.GetLogger()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&JSONFormatter{
		FieldMap: FieldMap{
			FieldKeyMsg:   "@message",
			FieldKeyLevel: "@level",
		},
	})
	logger.WithFields(map[string]interface{}{
		FieldKeyTime: "2000-01-01T00:00:00+09:00", // fixed time
	}).Info("aaa")
	// Output: {"@level":"INFO","@message":"aaa","time":"2000-01-01T00:00:00+09:00"}
}

func TestJSONFormatter_Format(t *testing.T) {
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name    string
		j       *JSONFormatter
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "default",
			j:       &JSONFormatter{},
			args:    args{entry: mockEntry(logrus.InfoLevel, "hoge", nil)},
			want:    []byte(`{"level":"INFO","message":"hoge","time":"2000-01-01T00:00:00+09:00"}` + "\n"),
			wantErr: false,
		},
		{
			name:    "empty message",
			j:       &JSONFormatter{},
			args:    args{entry: mockEntry(logrus.InfoLevel, "", nil)},
			want:    []byte(`{"level":"INFO","message":"","time":"2000-01-01T00:00:00+09:00"}` + "\n"),
			wantErr: false,
		},
		{
			name:    "with_fields",
			j:       &JSONFormatter{},
			args:    args{entry: mockEntry(logrus.InfoLevel, "hoge", &logrus.Fields{"hoge": "fuga"})},
			want:    []byte(`{"hoge":"fuga","level":"INFO","message":"hoge","time":"2000-01-01T00:00:00+09:00"}` + "\n"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.Format(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONFormatter.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("JSONFormatter.Format() = %s, want %s", got, tt.want)
			}
		})
	}
}
