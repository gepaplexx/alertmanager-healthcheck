package logging

import (
	"go.uber.org/zap"
	"encoding/json"
)

// Wrapper for a Logging Library
// Currently Wraps Uber Zap
type Logger struct {
	logger *zap.Logger
}

// Creates a new Logger Instance
func NewLogger() Logger {

	//JSON Config Data
	//TODO: Move to file
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	}`)
	
	var wrapper Logger
	wrapper.logger = instantiateLogger(rawJSON)
	return wrapper
}

// Instantiates a Zap Logger
// Returns a zap.Logger Instance
func instantiateLogger(loggerConf []byte) *zap.Logger {
	var cfg zap.Config
  	if err := json.Unmarshal(loggerConf, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
	   panic(err)
	}

	return logger 
}

// Logs on the Info Level
func (logger Logger) Info(msg string)  {
	defer logger.logger.Sync()
	logger.logger.Info(msg)
}

// Logs on the Warn Level
func (logger Logger) Warn(msg string)  {
	logger.logger.Warn(msg)
}

// Logs on the Error Level
func (logger Logger) Error(msg string)  {
	logger.logger.Error(msg)
}