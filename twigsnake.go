package twigsnake

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

// Logger is the logging object itself. Under the hood it has separate log.Logger instance for every severity level. All of them are
// exported, so you can fine-tune them individually (set custom prefix, output and whatever log.Logger allows to to with it).
type Logger struct {
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

// New creates new Logger instance with specified logging level and output; by default messages of every severity level
// will have its own prefix and output flags of underlying log.Logger set to log.Ldate|log.Ltime|log.Lmsgprefix. Prefixes are:
//	Emergency level 	- [EMERG]
//	Alert level		- [ALERT]
//	Critical level		- [CRIT]
//	Error level		- [ERROR]
//	Warning level		- [WARN]
//	Notification level	- [NOTICE]
//	Informational level	- [INFO]
//	Debug level		- [DEBUG]
func New(lvl int, dest io.Writer) (*Logger, error) {
	if err := checkLogLevel(lvl); err != nil {
		return nil, err

	}

	return &Logger{
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
func (l *Logger) LogLevel() int {
	return l.logLevel
}

// SetLogLevel sets logging level. Returns error if specified level is incorrect.
func (l *Logger) SetLogLevel(lvl int) error {
	if err := checkLogLevel(lvl); err != nil {
		return err
	}
	l.logLevel = lvl
	return nil
}

// Emerg prints emergency messages. They will appear on any logging level. Handles arguments in the same manner as log.Print.
func (l *Logger) Emerg(v ...interface{}) {
	l.EmergLogger.Print(v...)
}

// Emergf prints emergency messages. They will appear on any logging level. Handles arguments in the same manner as log.Printf.
func (l *Logger) Emergf(format string, v ...interface{}) {
	l.EmergLogger.Printf(format, v...)
}

// Emergln prints emergency messages. They will appear on any logging level. Handles arguments in the same manner as log.Println.
func (l *Logger) Emergln(v ...interface{}) {
	l.EmergLogger.Println(v...)
}

// Alert prints alert messages. They will appear on logging level twigsnake.LOG_ALERT and higher. Handles arguments in the same
// manner as log.Print.
func (l *Logger) Alert(v ...interface{}) {
	if l.logLevel >= LOG_ALERT {
		l.AlertLogger.Print(v...)
	}
}

// Alertf prints alert messages. They will appear on logging level twigsnake.LOG_ALERT and higher. Handles arguments in the same
// manner as log.Printf.
func (l *Logger) Alertf(format string, v ...interface{}) {
	if l.logLevel >= LOG_ALERT {
		l.AlertLogger.Printf(format, v...)
	}
}

// Alertln prints alert messages. They will appear on logging level twigsnake.LOG_ALERT and higher. Handles arguments in the same
// manner as log.Println.
func (l *Logger) Alertln(v ...interface{}) {
	if l.logLevel >= LOG_ALERT {
		l.AlertLogger.Println(v...)
	}
}

// Crit prints critical messages. They will appear on logging level twigsnake.LOG_CRIT and higher. Handles arguments in the same
// manner as log.Print.
func (l *Logger) Crit(v ...interface{}) {
	if l.logLevel >= LOG_CRIT {
		l.CritLogger.Print(v...)
	}
}

// Critf prints critical messages. They will appear on logging level twigsnake.LOG_CRIT and higher. Handles arguments in the same
// manner as log.Printf.
func (l *Logger) Critf(format string, v ...interface{}) {
	if l.logLevel >= LOG_CRIT {
		l.CritLogger.Printf(format, v...)
	}
}

// Critln prints critical messages. They will appear on logging level twigsnake.LOG_CRIT and higher. Handles arguments in the same
// manner as log.Println.
func (l *Logger) Critln(v ...interface{}) {
	if l.logLevel >= LOG_CRIT {
		l.CritLogger.Println(v...)
	}
}

// Error prints error messages. They will appear on logging level twigsnake.LOG_ERROR and higher. Handles arguments in the same
// manner as log.Print.
func (l *Logger) Error(v ...interface{}) {
	if l.logLevel >= LOG_ERROR {
		l.ErrorLogger.Print(v...)
	}
}

// Errorf prints error messages. They will appear on logging level twigsnake.LOG_ERROR and higher. Handles arguments in the same
// manner as log.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.logLevel >= LOG_ERROR {
		l.ErrorLogger.Printf(format, v...)
	}
}

// Errorln prints error messages. They will appear on logging level twigsnake.LOG_ERROR and higher. Handles arguments in the same
// manner as log.Println.
func (l *Logger) Errorln(v ...interface{}) {
	if l.logLevel >= LOG_ERROR {
		l.ErrorLogger.Println(v...)
	}
}

// Warn prints warning messages. They will appear on logging level twigsnake.LOG_WARN and higher. Handles arguments in the same
// manner as log.Print.
func (l *Logger) Warn(v ...interface{}) {
	if l.logLevel >= LOG_WARN {
		l.WarningLogger.Print(v...)
	}
}

// Warnf prints warning messages. They will appear on logging level twigsnake.LOG_WARN and higher. Handles arguments in the same
// manner as log.Printf.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.logLevel >= LOG_WARN {
		l.WarningLogger.Printf(format, v...)
	}
}

// Warnln prints warning messages. They will appear on logging level twigsnake.LOG_WARN and higher. Handles arguments in the same
// manner as log.Println.
func (l *Logger) Warnln(v ...interface{}) {
	if l.logLevel >= LOG_WARN {
		l.WarningLogger.Println(v...)
	}
}

// Notice prints notification messages. They will appear on logging level twigsnake.LOG_NOTICE and higher. Handles arguments in the
// same manner as log.Print.
func (l *Logger) Notice(v ...interface{}) {
	if l.logLevel >= LOG_NOTICE {
		l.NoticeLogger.Print(v...)
	}
}

// Noticef prints notification messages. They will appear on logging level twigsnake.LOG_NOTICE and higher. Handles arguments in the
// same manner as log.Printf.
func (l *Logger) Noticef(format string, v ...interface{}) {
	if l.logLevel >= LOG_NOTICE {
		l.NoticeLogger.Printf(format, v...)
	}
}

// Noticeln prints notification messages. They will appear on logging level twigsnake.LOG_NOTICE and higher. Handles arguments in the
// same manner as log.Println.
func (l *Logger) Noticeln(v ...interface{}) {
	if l.logLevel >= LOG_NOTICE {
		l.NoticeLogger.Println(v...)
	}
}

// Info prints informational messages. They will appear on logging level twigsnake.LOG_INFO and higher. Handles arguments in the same
// manner as log.Print.
func (l *Logger) Info(v ...interface{}) {
	if l.logLevel >= LOG_INFO {
		l.InfoLogger.Print(v...)
	}
}

// Infof prints informational messages. They will appear on logging level twigsnake.LOG_INFO and higher. Handles arguments in the same
// manner as log.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.logLevel >= LOG_INFO {
		l.InfoLogger.Printf(format, v...)
	}
}

// Infoln prints informational messages. They will appear on logging level twigsnake.LOG_INFO and higher. Handles arguments in the same
// manner as log.Println.
func (l *Logger) Infoln(v ...interface{}) {
	if l.logLevel >= LOG_INFO {
		l.InfoLogger.Println(v...)
	}
}

// Debug prints debugging messages. They will appear only on logging level twigsnake.LOG_DEBUG. Handles arguments in the same manner
// as log.Print.
func (l *Logger) Debug(v ...interface{}) {
	if l.logLevel >= LOG_DEBUG {
		l.DebugLogger.Print(v...)
	}
}

// Debugf prints debugging messages. They will appear only on logging level twigsnake.LOG_DEBUG. Handles arguments in the same manner
// as log.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.logLevel >= LOG_DEBUG {
		l.DebugLogger.Printf(format, v...)
	}
}

// Debugln prints debugging messages. They will appear only on logging level twigsnake.LOG_DEBUG. Handles arguments in the same manner
// as log.Println.
func (l *Logger) Debugln(v ...interface{}) {
	if l.logLevel >= LOG_DEBUG {
		l.DebugLogger.Println(v...)
	}
}
