package main

import (
	"./JNProgrammer"
	"os"
	"log"
	"flag"
)

var logfile = "JNProgrammer.log"

func main() {
	var fpLog *os.File
	JNProgrammer.DEBUG = *flag.Bool("x", false, "Program debug")
	logfile = *flag.String("l", "", "log")
	JNProgrammer.LogLevel = JNProgrammer.VERVOSITY(*flag.Int("v", 0, "Log level. 0 - 4. higher is vervose"))
	if logfile == "" {
		fpLog = os.Stdout
	}
	flag.Parse()

	if fpLog == nil {
		fpLog, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer fpLog.Close()
	}
	JNProgrammer.Logger = log.New(fpLog, "", log.Ldate|log.Ltime|log.Lshortfile)

	var fwInfo *JNProgrammer.FirmwareInfo
	fwInfo = &JNProgrammer.FirmwareInfo{}
	JNProgrammer.FwOpen(fwInfo, "./firmware.bin")
}
