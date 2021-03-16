// Package twigsnake  provides a minimalistic wrapper for standard logger, allowing you to implement leveled logging with
// minimum overhead.
//
// Twigsnake produces output to io.Writer and supports eight logging levels: emergency, alert, critical, error, warning, notice,
// informational and debug. Each package logging object instance, twigsnake.Logger, has its own logging level and will print only
// messages with equal or higher severity.
// Package provides convinient wrappers around standard log.Logger's Print, Println and Printf methods for each severity level.
//
// Example - basic usage
//
// In this example we will stick to defaults: create twigsnake.Logger without any customization and log some stuff with it.
//
//	package main
//
//	import (
//		"github.com/lereinth/twigsnake"
//		"os"
//	)
//
//	func main() {
//		// Create twigsnake.Logger instance with informational logging level (i.e. log everything except debug messages) and
//		// with output directed to console
//		logger, err := twigsnake.New(twigsnake.LOG_INFO, os.Stdout)
//		if err != nil {
//			panic(err)
//		}
//
//		// Log some stuff
//		logger.Infoln("This is informational message")
//		logger.Noticef("Formatted notice message: here are some random numbers and strings: %d, %s", 10, "random string")
//		logger.Debugln("This debug message won't appear")
//
//		// Change current severity level
//		err = logger.SetLogLevel(twigsnake.LOG_DEBUG)
//		if err != nil {
//			logger.Errorln("Failed to change logging level:", err)
//		}
//
//		// NOW this message will appear
//		logger.Debugln("Behold! Debug message")
//	}
//
// This code will print following lines to the console:
//	2021/03/05 16:21:32 [INFO] This is informational message
//	2021/03/05 16:21:32 [NOTICE] Formatted notice message: here are some random numbers and strings: 10, random string
//	2021/03/05 16:21:32 [DEBUG] Behold! Debug message
//
// Example - advanced usage
//
// Now let's make our task a little bit more intricate. Suppose we want to make timestamps of debug messages more precise and send
// those messages to the separate destination, say, text file. Messages of all other severities will go to the console, just as before.
//
//	package main
//
//	import (
//		"github.com/lereinth/twigsnake"
//		"log"
//		"os"
//	)
//
//	func main() {
//		// Create twigsnake.Logger instance, just like in the previous example
//		logger, err := twigsnake.New(twigsnake.LOG_DEBUG, os.Stdout)
//		if err != nil {
//			panic(err)
//		}
//
//		// Set up a file for collecting debug messages
//		debugLog, err := os.OpenFile("debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
//		if err != nil {
//			panic(err)
//		}
//		defer debugLog.Close()
//
//		// It's customization time! First, make message timestamp more precise
//		logger.DebugLogger.SetFlags(log.Ldate | log.Lmicroseconds)
//
//		// Get rid of prefix: we don't need it anyway since there will be only debug messages in our file
//		logger.DebugLogger.SetPrefix("")
//
//		// Finally, set our file as destination for debug messages
//		logger.DebugLogger.SetOutput(debugLog)
//
//		// Log some stuff
//		logger.Infoln("Regular message")
//		logger.Debugln("Debug message")
//	}
//
// Informational message will appear at the console as usual:
//	2021/03/05 16:21:32 [INFO] Regular message
//
// Meanwhile in debug.log the next entry will appear:
//	2021/03/05 16:21:32.100005 Debug message
package twigsnake
