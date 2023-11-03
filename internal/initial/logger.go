package initial

import (
	"github.com/TakasakiApps/flea-trove-go/internal/consts"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/hanakogo/digine"
	"github.com/valyala/bytebufferpool"
)

var loggerFormat = []string{
	"[", "{{app}}", "-", "{{build_type}}", "] ",
	"[", "{{level}}", "] ",
	"[", "{{datetime}}", "] ",
	"{{message}}", "\n",
}

type LoggerFormatter struct {
	slog.TextFormatter
}

func NewLoggerFormatter() *LoggerFormatter {
	return &LoggerFormatter{
		TextFormatter: slog.TextFormatter{
			TimeFormat: "2006/01/02 - 15:04:05",
			// EnableColor: true,
			ColorTheme: slog.ColorTheme,
			// FullDisplay: false,
			EncodeFunc: slog.EncodeToString,
		},
	}
}

// from gookit/slog/formatter_test.go
var textPool bytebufferpool.Pool

func (l *LoggerFormatter) renderColorByLevel(text string, level slog.Level) string {
	if theme, ok := l.ColorTheme[level]; ok {
		return theme.Render(text)
	}
	return text
}

//goland:noinspection GoUnhandledErrorResult
func (l *LoggerFormatter) Format(r *slog.Record) ([]byte, error) {
	buf := textPool.Get()
	defer textPool.Put(buf)

	for _, field := range loggerFormat {
		if !(strutil.IsStartOf(field, "{{") && strutil.IsEndOf(field, "}}")) {
			buf.WriteString(field)
			continue
		}

		switch {
		case field == "{{app}}":
			buf.WriteString(consts.AppName)
		case field == "{{build_type}}":
			buf.WriteString(consts.BuildType)
		case field == "{{datetime}}":
			buf.B = r.Time.AppendFormat(buf.B, l.TimeFormat)
		case field == "{{level}}":
			buf.WriteString(l.renderColorByLevel(r.LevelName(), r.Level))
		case field == "{{message}}":
			buf.WriteString(l.renderColorByLevel(r.Message, r.Level))
		}
	}

	return buf.B, nil
}

func initLogger() {
	loggerHandler := handler.NewConsoleHandler(slog.AllLevels)
	loggerHandler.SetFormatter(NewLoggerFormatter())
	logger := slog.NewWithHandlers(loggerHandler)

	digine.Bind[slog.Logger](logger, digine.NilLabel)
}
