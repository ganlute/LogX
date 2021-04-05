package logrus

func Trace(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Trace(args...)
}
func Debug(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Debug(args...)
}
func Print(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Print(args...)
}
func Info(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Info(args...)
}
func Warn(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Warn(args...)
}

func Warning(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Warning(args...)
}
func Error(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Error(args...)
}
func Fatal(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Fatal(args...)
}
func Panic(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Panic(args...)
}
func Tracef(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Tracef(format, args...)
}
func Debugf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Infof(format, args...)
}
func Printf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Printf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Warnf(format, args...)
}
func Warningf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Warningf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Fatalf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Panicf(format, args...)
}

func Traceln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Traceln(args...)
}
func Debugln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Debugln(args...)
}
func Infoln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Infoln(args...)
}
func Println(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Println(args...)
}
func Warnln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Warnln(args...)
}
func Warningln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Warningln(args...)
}
func Errorln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Errorln(args...)
}
func Fatalln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Fatalln(args...)
}
func Panicln(args ...interface{}) {
	logEntry := getLogEntry(1)
	logEntry.Panicln(args...)
}
