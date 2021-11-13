package log

import (
	"fmt"
	"testing"
)

var lgr = New(&Config{})

func TestLogger_Configure(t *testing.T) {
	fmt.Println(lgr)
	lgr.Configure(&Config{Out: StdBoth, Traceback: true})
	fmt.Println(lgr)
}

func TestLogger_Print(t *testing.T) {
	lgr.Print("Logger Print")
	lgr.Format("%s").Print("Logger Format Print")
	fmt.Println(lgr)
	lgr.Configure(&Config{Traceback: true})
	lgr.Print("Logger Traceback Print")
	lgr.Format("%s").Print("Logger Traceback Format Print")
	fmt.Println(lgr)
}

func TestLogger_Trace(t *testing.T) {
	lgr.Trace("Logger Trace")
	lgr.Format("%s").Trace("Logger Format Trace")
}

func TestLogger_Debug(t *testing.T) {
	lgr.Debug("Logger Debug")
	lgr.Format("%s").Debug("Logger Format Debug")
}

func TestLogger_Info(t *testing.T) {
	lgr.Info("Logger Info")
	lgr.Format("%s").Info("Logger Format Info")
}

func TestLogger_Notice(t *testing.T) {
	lgr.Notice("Logger Notice")
	lgr.Format("%s").Notice("Logger Format Notice")
}

func TestLogger_Warning(t *testing.T) {
	lgr.Warning("Logger Warning")
	lgr.Format("%s").Warning("Logger Format Warning")
}

func TestLogger_Error(t *testing.T) {
	lgr.Error("Logger Error")
	lgr.Format("%s").Error("Logger Format Error")
}

func TestLogger_Fatal(t *testing.T) {
	lgr.Fatal("Logger Fatal")
}

func TestLogger_Fatal2(t *testing.T) {
	lgr.Format("%s").Fatal("Logger Format Fatal")
}

func TestLogger_Panic(t *testing.T) {
	lgr.Panic("Logger Panic")
}

func TestLogger_Panic2(t *testing.T) {
	lgr.Format("%s").Panic("Logger Format Panic")
}

func TestConfigure(t *testing.T) {
	fmt.Println(std)
	Configure(&Config{Out: StdBoth, Traceback: true})
	fmt.Println(std)
}

func TestPrint(t *testing.T) {
	Print("Log Print")
	Format("%s").Print("Log Format Print")
	fmt.Println(std)
	Configure(&Config{Traceback: true})
	Print("Log Traceback Print")
	Format("%s").Print("Log Traceback Format Print")
	fmt.Println(std)
}

func TestTrace(t *testing.T) {
	Trace("Log Trace")
	Format("%s").Trace("Log Format Trace")
}

func TestDebug(t *testing.T) {
	Debug("Log Debug")
	Format("%s").Debug("Log Format Debug")
}

func TestInfo(t *testing.T) {
	Info("Log Info")
	Format("%s").Info("Log Format Info")
}

func TestNotice(t *testing.T) {
	Notice("Log Notice")
	Format("%s").Notice("Log Format Notice")
}

func TestWarning(t *testing.T) {
	Warning("Log Warning")
	Format("%s").Warning("Log Format Warning")
}

func TestError(t *testing.T) {
	Error("Log Error")
	Format("%s").Error("Log Format Error")
}

func TestFatal(t *testing.T) {
	Fatal("Log Fatal")
}

func TestFatal2(t *testing.T) {
	Format("%s").Fatal("Log Format Fatal")
}

func TestPanic(t *testing.T) {
	Panic("Log Panic")
}

func TestPanic2(t *testing.T) {
	Format("%s").Panic("Log Format Panic")
}