package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// 设置日志信息的输出目的地，配置项Out取值
const (
	Discard = 1 << iota
	Stdout
	Stderr
	StdFile
	StdFlag = Stdout | Stderr
	StdBoth = StdFlag | StdFile
)

// 设置日志信息的输出细节，配置项Flag取值
const (
	Date         = 1 << iota
	Time
	Microseconds
	LongFile
	ShortFile
	StdFlags     = Date | Time
)

func caller(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)

	if !ok {
		return ""
	}

	tmp := file + ":" + strconv.Itoa(line)

	//tmp = tmp + runtime.FuncForPC(pc).Name()
	_ = pc

	return tmp
}

func logger(out io.Writer, prefix string, flag int) *log.Logger {
	prefix = strings.Join([]string{
		"[",
		prefix,
		"] ",
	}, "")

	return log.New(out, prefix, flag)
}

func writer(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		stdLog := logger(os.Stderr, "Fatal", log.LstdFlags|log.Lshortfile)
		stdLog.Fatalln(path.Base(caller(2)) + ":", "Failed to open log file:", err)
	}

	return file
}

type Config struct {
	Out       int
	Flag      int
	depth     int
	Traceback bool
	Prefix    string
	Filename  string
}

type Logger struct {
	out       io.Writer
	error     io.Writer
	filename  string
	format    string
	prefix    string
	flag      int
	status    int
	depth     int
	traceback bool
}

func New(cfg *Config) *Logger {
	lgr := &Logger{
		out:       os.Stdout,
		error:     os.Stderr,
		flag:      log.LstdFlags | log.Lshortfile,
		status:    StdFlag,
		depth:     5,
	}

	lgr.Configure(cfg)

	return lgr
}

func (lgr *Logger) Configure(cfg *Config) *Logger {
	lgr.traceback = cfg.Traceback

	if cfg.Prefix != "" {
		lgr.prefix = cfg.Prefix
	}

	if cfg.Filename != "" {
		lgr.filename = cfg.Filename
	} else {
		lgr.filename = "./runtime.log"
	}

	var file *os.File
	if cfg.Out == 8 || cfg.Out == 10 || cfg.Out == 12 || cfg.Out == 14 {
		file = writer(lgr.filename)
	}

	if cfg.Out > 0 {
		lgr.out    = ioutil.Discard
		lgr.status = cfg.Out

		switch cfg.Out {
		case 2:
			lgr.out = os.Stdout
		case 4:
			lgr.out = os.Stderr
		case 8:
			lgr.out = file
		case 10:
			lgr.out = io.MultiWriter(file, os.Stdout)
		case 12:
			lgr.out = io.MultiWriter(file, os.Stderr)
		case 14:
			lgr.out   = io.MultiWriter(file, os.Stdout)
			lgr.error = io.MultiWriter(file, os.Stderr)
		default:
			lgr.status = Discard
		}

		if cfg.Out == 6 {
			lgr.out   = os.Stdout
			lgr.error = os.Stderr
		}

		if cfg.Out != 6 && cfg.Out != 14 {
			lgr.error = lgr.out
		}
	}

	if cfg.Flag > 0 {
		lgr.flag = cfg.Flag
	}

	if cfg.depth > 0 {
		lgr.depth = cfg.depth
	}

	return lgr
}

var std = New(&Config{depth: 6})

func (lgr *Logger) Format(s string) *Logger {
	lgr.format = s

	return lgr
}

func (lgr *Logger) caller(skip int) string {
	tmp := caller(skip)

	if lgr.flag >= 8 && lgr.flag <= 15 {
		return tmp
	} else if lgr.flag >= 16 && lgr.flag <= 25 {
		return path.Base(tmp)
	}

	return ""
}

func (lgr *Logger) echo(i ...interface{}) {
	if lgr.status == Discard {
		return
	}

	if lgr.prefix != "" {
		lgr.prefix = strings.Join([]string{
			"[",
			lgr.prefix,
			"] ",
		}, "")
	}

	var (
		loc string
		arr []interface{}
	)

	if lgr.traceback {
		for i := 3; i <= lgr.depth; i++ {
			loc = loc + lgr.caller(i) + ": "
		}

		arr = append(arr, strings.TrimRight(loc, " "))
	}

	//arr = append(arr, lgr.prefix)
	arr = append(arr, i...)

	stdLog := log.New(lgr.out, lgr.prefix, lgr.flag)

	if lgr.format != "" {
		stdLog.Printf(loc + lgr.format, i...)
		//stdLog.Println()
		lgr.format = ""
	} else {
		stdLog.Println(arr...)
	}
}

func (lgr *Logger) run(i ...interface{}) {
	tmp := lgr.prefix

	if func(prefix string, args []string) (ok bool) {
		for i := 0; i < len(args); i++ {
			if strings.HasPrefix(strings.ToLower(prefix), args[i]) {
				return ok
			}
		}

		return ok
	}(lgr.prefix, []string{"fatal", "err", "warn"}) {
		lgr.out, lgr.error = lgr.error, lgr.out
		lgr.echo(i...)
		lgr.error, lgr.out = lgr.out, lgr.error
	} else {
		lgr.echo(i...)
	}

	lgr.prefix = tmp
}

func (lgr *Logger) Print(i ...interface{}) {
	lgr.run(i...)
}

func (lgr *Logger) Trace(i ...interface{}) {
	lgr.prefix = "Trace"
	lgr.run(i...)
}

func (lgr *Logger) Debug(i ...interface{}) {
	lgr.prefix = "Debug"
	lgr.run(i...)
}

func (lgr *Logger) Info(i ...interface{}) {
	lgr.prefix = "Info"
	lgr.run(i...)
}

func (lgr *Logger) Notice(i ...interface{}) {
	lgr.prefix = "Notice"
	lgr.run(i...)
}

func (lgr *Logger) Warning(i ...interface{}) {
	lgr.prefix = "Warning"
	lgr.run(i...)
}

func (lgr *Logger) Error(i ...interface{}) {
	lgr.prefix = "Error"
	lgr.run(i...)
}

func (lgr *Logger) Fatal(i ...interface{}) {
	lgr.prefix = "Fatal"
	lgr.run(i...)
	os.Exit(1)
}

func (lgr *Logger) Panic(i ...interface{}) {
	lgr.prefix = "Panic"
	lgr.run(i...)
	panic(fmt.Sprint(i...))
}

func Configure(cfg *Config) *Logger {
	return std.Configure(cfg)
}

func Format(s string) *Logger {
	std.format = s

	return std
}

func Print(i ...interface{}) {
	std.Print(i...)
}

func Trace(i ...interface{}) {
	std.Trace(i...)
}

func Debug(i ...interface{}) {
	std.Debug(i...)
}

func Info(i ...interface{}) {
	std.Info(i...)
}

func Notice(i ...interface{}) {
	std.Notice(i...)
}

func Warning(i ...interface{}) {
	std.Warning(i...)
}

func Error(i ...interface{}) {
	std.Error(i...)
}

func Fatal(i ...interface{}) {
	std.Fatal(i...)
}

func Panic(i ...interface{}) {
	std.Panic(i...)
}