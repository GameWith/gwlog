# gwlog

gamewith custom logrus logger for golang

## Example

```go
func main() {
	// singleton instance
	// same > logger == gwlog.GetLogger() 
	logger := gwlog.GetLogger()
	
	// SetFormatter (implements logrus.Formatter)
	// Default Formatter logrus.TextFormatter
	logger.SetFormatter(&formatter.JSONFormatter{})
	
	logger.Info("abc")
	// => {"level":"INFO","message":"abc","time":"2000-01-01T00:00:00+09:00"}
	
	logger.WithFields(map[string]interface{}{
		"hoge": 1,
		"fuga": "2"
	}).Info("abc")
	// => {"hoge":1,"fuga":"2","level":"INFO","message":"abc","time":"2000-01-01T00:00:00+09:00"}
    
	logger.Type("APP").Info("abc")
	// => {"type":"APP","level":"INFO","message":"abc","time":"2000-01-01T00:00:00+09:00"}
}
```
