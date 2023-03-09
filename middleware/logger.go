package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/logrus.log"
	logger := logrus.New()
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0075)
	if err != nil {
		logger.Info("Failed to log to file, using default stderr")
	} else {
		out := io.MultiWriter( /*os.Stdout, */ file)
		logger.SetOutput(out)
	}
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		clientIp := c.ClientIP()
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.URL
		entry := logger.WithFields(logrus.Fields{
			"HostNAme":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}
