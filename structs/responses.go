package structs

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

// StructuredLogger is a simple, but powerful implementation of a custom structured
// logger backed on logrus. I encourage users to copy it, adapt it and make it their
// own. Also take a look at https://github.com/pressly/lg for a dedicated pkg based
// on this work, designed for context-based http routers.
func NewRequestLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger *logrus.Logger
}

func NewLogger() StructuredLogger {
	Log := logrus.New()
	Log.SetReportCaller(true)
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.DebugLevel)
	return StructuredLogger{
		Logger: Log,
	}
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}
	// user information
	logFields["remote_addr"] = r.RemoteAddr
	// important headers
	logFields["x-api-key"] = r.Header.Get("x-api-key")
	logFields["uri"] = fmt.Sprintf("%s%s", r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)
	entry.Logger.Infoln("request started")

	return entry
}

func (l *StructuredLogger) WithUser(userID string) *logrus.Entry {
	fields := logrus.Fields{}
	if userID != "" {
		fields["userID"] = userID
	}
	return l.Logger.WithFields(fields)
}

func (l *StructuredLogger) WithRequest(r *http.Request) *logrus.Entry {
	fields := logrus.Fields{
		"http_method": r.Method,
		"remote_addr": r.RemoteAddr,
		"uri":         fmt.Sprintf("%s%s", r.Host, r.RequestURI),
	}
	return l.Logger.WithFields(fields)
}

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Infoln("request complete")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}
