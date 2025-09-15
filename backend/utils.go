package backend

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

var AllLogLevel = []struct {
	Value   LogLevel
	TsValue string
}{
	{DEBUG, "DEBUG"},
	{INFO, "INFO"},
	{WARN, "WARN"},
	{ERROR, "ERROR"},
}

type LogMsg struct {
	Time  int64
	Level LogLevel
	Msg   string
}

func (app *App) Log(msg string) LogMsg {
	return LogMsg{
		Time: time.Now().UnixMilli(),
		Msg:  msg,
	}
}

// Log 打印日志并发送事件
func Log(level LogLevel, msg string) {
	ts := time.Now().UnixMilli()
	fullMsg := fmt.Sprintf("%d [%s] %s", ts, level, msg)
	// 打印到控制台，带颜色
	switch level {
	case INFO:
		color.Cyan(fullMsg)
	case DEBUG:
		color.Blue(fullMsg)
	case WARN:
		color.Yellow(fullMsg)
	case ERROR:
		color.Red(fullMsg)
	}

	// 推送到前端
	go runtime.EventsEmit(app.ctx, "logMsg", LogMsg{
		Time:  ts,
		Level: level,
		Msg:   msg,
	})
}

func LogDebug(msg string) { Log(DEBUG, msg) }
func LogInfo(msg string)  { Log(INFO, msg) }
func LogWarn(msg string)  { Log(WARN, msg) }
func LogError(msg string) { Log(ERROR, msg) }
