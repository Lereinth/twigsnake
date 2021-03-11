package nlog

import (
	"errors"
	"io"
	"log"
)

// Log severity levels (as defined in RFC 5424 section 6.2.1):
const (
	LOG_EMERG  = iota // Emergency: system is unusable
	LOG_ALERT         // Alert: action must be taken immediately
	LOG_CRIT          // Critical: critical conditions
	LOG_ERROR         // Error: error conditions
	LOG_WARN          // Warning: warning conditions
	LOG_NOTICE        // Notice: normal, but significant conditions
	LOG_INFO          // Informational: informational messages
	LOG_DEBUG         // Debug: debug-level messages
)

// NLog is the logging object itself. Under the hood it has separate log.Logger instance for every severity level. All of them are
// exported, so you can fine-tune them individually (set custom prefix, output and whatever log.Logger allows to to with it).
type NLog struct {
	logLevel int

	// Collection of standard loggers for every severity level:
	EmergLogger   *log.Logger
	AlertLogger   *log.Logger
	CritLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
	NoticeLogger  *log.Logger
	InfoLogger    *log.Logger
	DebugLogger   *log.Logger
}

func checkLogLevel(lvl int) error {
	if !(lvl >= 0 && lvl <= 7) {
		return errors.New("invalid severity level")
	}
	return nil
}

// New creates new NLog instance with specified logging level and output; by default messages of every severity level
// will have its own prefix and output flags of underlying log.Logger set to log.Ldate|log.Ltime|log.Lmsgprefix. Prefixes are:
//	Emergency level 	- [EMERG]
//	Alert level		- [ALERT]
//	Critical level		- [CRIT]
//	Error level		- [ERROR]
//	Warning level		- [WARN]
//	Notification level	- [NOTICE]
//	Informational level	- [INFO]
//	Debug level		- [DEBUG]
func New(lvl int, dest io.Writer) (*NLog, error) {
	if err := checkLogLevel(lvl); err != nil {
		return nil, err

	}

	return &NLog{
		lvl,
		log.New(dest, "[EMERG] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[ALERT] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[CRIT] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[ERROR] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[WARN] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[NOTICE] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[INFO] ", log.Ldate|log.Ltime|log.Lmsgprefix),
		log.New(dest, "[DEBUG] ", log.Ldate|log.Ltime|log.Lmsgprefix),
	}, nil

}

// LogLevel returns current logging level.
func (n *NLog) LogLevel() int {
	return n.logLevel
}

// SetLogLevel sets logging level. Returns error if specified level is incorrect.
func (n *NLog) SetLogLevel(lvl int) error {
	if err := checkLogLevel(lvl); err != nil {
		return err
	}
	n.logLevel = lvl
	return nil
}

// Emerg prints emergency messages. They will appear on any logging level. Handles arguments in the same manner as log.Print.
func (n *NLog) Emerg(v ...interface{}) {
	n.EmergLogger.Print(v...)
}

// Emergf prints emergency messages. They will appear on any logging level. Handles arguments in the same manner as log.Printf.
func (n *NLog) Emergf(format string, v ...interface{}) {
	n.EmergLogger.Printf(format, v...)
}

// Emergln prints emergency messages. They will appear on any logging level. Handles arguments in the same manner as log.Println.
func (n *NLog) Emergln(v ...interface{}) {
	n.EmergLogger.Println(v...)
}

// Alert prints alert messages. They will appear on logging level nlog.LOG_ALERT and higher. Handles arguments in the same manner
// as log.Print.
func (n *NLog) Alert(v ...interface{}) {
	if n.logLevel >= LOG_ALERT {
		n.AlertLogger.Print(v...)
	}
}

// Alertf prints alert messages. They will appear on logging level nlog.LOG_ALERT and higher. Handles arguments in the same manner
// as log.Printf.
func (n *NLog) Alertf(format string, v ...interface{}) {
	if n.logLevel >= LOG_ALERT {
		n.AlertLogger.Printf(format, v...)
	}
}

// Alertln prints alert messages. They will appear on logging level nlog.LOG_ALERT and higher. Handles arguments in the same manner
// as log.Println.
func (n *NLog) Alertln(v ...interface{}) {
	if n.logLevel >= LOG_ALERT {
		n.AlertLogger.Println(v...)
	}
}

// Crit prints critical messages. They will appear on logging level nlog.LOG_CRIT and higher. Handles arguments in the same manner
// as log.Print.
func (n *NLog) Crit(v ...interface{}) {
	if n.logLevel >= LOG_CRIT {
		n.CritLogger.Print(v...)
	}
}

// Critf prints critical messages. They will appear on logging level nlog.LOG_CRIT and higher. Handles arguments in the same manner
// as log.Printf.
func (n *NLog) Critf(format string, v ...interface{}) {
	if n.logLevel >= LOG_CRIT {
		n.CritLogger.Printf(format, v...)
	}
}

// Critln prints critical messages. They will appear on logging level nlog.LOG_CRIT and higher. Handles arguments in the same manner
// as log.Println.
func (n *NLog) Critln(v ...interface{}) {
	if n.logLevel >= LOG_CRIT {
		n.CritLogger.Println(v...)
	}
}

// Error prints error messages. They will appear on logging level nlog.LOG_ERROR and higher. Handles arguments in the same manner
// as log.Print.
func (n *NLog) Error(v ...interface{}) {
	if n.logLevel >= LOG_ERROR {
		n.ErrorLogger.Print(v...)
	}
}

// Errorf prints error messages. They will appear on logging level nlog.LOG_ERROR and higher. Handles arguments in the same manner
// as log.Printf.
func (n *NLog) Errorf(format string, v ...interface{}) {
	if n.logLevel >= LOG_ERROR {
		n.ErrorLogger.Printf(format, v...)
	}
}

// Errorln prints error messages. They will appear on logging level nlog.LOG_ERROR and higher. Handles arguments in the same manner
// as log.Println.
func (n *NLog) Errorln(v ...interface{}) {
	if n.loglLevel >= LOG_ERROR {
		n.ErrorLogger.Println(v...)
	}
}

// Warn prints warning messages. They will appear on logging level nlog.LOG_WARN and higher. Handles arguments in the same manner
// as log.Print.
func (n *NLog) Warn(v ...interface{}) {
	if n.logLevel >= LOG_WARN {
		n.WarningLogger.Print(v...)
	}
}

// Warnf prints warning messages. They will appear on logging level nlog.LOG_WARN and higher. Handles arguments in the same manner
// as log.Printf.
func (n *NLog) Warnf(format string, v ...interface{}) {
	if n.logLevel >= LOG_WARN {
		n.WarningLogger.Printf(format, v...)
	}
}

// Warnln prints warning messages. They will appear on logging level nlog.LOG_WARN and higher. Handles arguments in the same manner
// as log.Println
func (n *NLog) Warnln(v ...interface{}) {
	if n.logLevel >= LOG_WARN {
		n.WarningLogger.Println(v...)
	}
}

// Notice prints notification messages. They will appear on logging level nlog.LOG_NOTICE and higher. Handles arguments in the same
// manner as log.Print
func (n *NLog) Notice(v ...interface{}) {
	if n.logLevel >= LOG_NOTICE {
		n.NoticeLogger.Print(v...)
	}
}

// Noticef prints notification messages. They will appear on logging level nlog.LOG_NOTICE and higher. Handles arguments in the same
// manner as log.Printf
func (n *NLog) Noticef(format string, v ...interface{}) {
	if n.logLevel >= LOG_NOTICE {
		n.NoticeLogger.Printf(format, v...)
	}
}

// Noticeln prints notification messages. They will appear on logging level nlog.LOG_NOTICE and higher. Handles arguments in the same
// manner as log.Println
func (n *NLog) Noticeln(v ...interface{}) {
	if n.logLevel >= LOG_NOTICE {
		n.NoticeLogger.Println(v...)
	}
}

// Info prints informational messages. They will appear on logging level nlog.LOG_INFO and higher. Handles arguments in the same
// manner as log.Print
func (n *NLog) Info(v ...interface{}) {
	if n.logLevel >= LOG_INFO {
		n.InfoLogger.Print(v...)
	}
}

// Infof prints informational messages. They will appear on logging level nlog.LOG_INFO and higher. Handles arguments in the same
// manner as log.Printf
func (n *NLog) Infof(format string, v ...interface{}) {
	if n.logLevel >= LOG_INFO {
		n.InfoLogger.Printf(format, v...)
	}
}

// Infoln prints informational messages. They will appear on logging level nlog.LOG_INFO and higher. Handles arguments in the same
// manner as log.Println
func (n *NLog) Infoln(v ...interface{}) {
	if n.logLevel >= LOG_INFO {
		n.InfoLogger.Println(v...)
	}
}

// Debug prints debugging messages. They will appear only on logging level nlog.LOG_DEBUG. Handles arguments in the same manner
// as log.Print
func (n *NLog) Debug(v ...interface{}) {
	if n.logLevel >= LOG_DEBUG {
		n.DebugLogger.Print(v...)
	}
}

// Debugf prints debugging messages. They will appear only on logging level nlog.LOG_DEBUG. Handles arguments in the same manner
// as log.Printf
func (n *NLog) Debugf(format string, v ...interface{}) {
	if n.logLevel >= LOG_DEBUG {
		n.DebugLogger.Printf(format, v...)
	}
}

// Debugln prints debugging messages. They will appear only on logging level nlog.LOG_DEBUG. Handles arguments in the same manner
// as log.Println
func (n *NLog) Debugln(v ...interface{}) {
	if n.logLevel >= LOG_DEBUG {
		n.DebugLogger.Println(v...)
	}
}
