package log

var LogConfig *Log

type Log struct {
	Level     string
	FilePath  string
	MaxSize   int
	MaxAge    int
	MaxBackup int
}

func (l *Log) GetSettingValues() (path, level string, size, backup, age int) {
	path = l.FilePath
	level = l.Level
	size = l.MaxSize
	backup = l.MaxBackup
	age = l.MaxAge
	return
}
