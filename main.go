//JNProgrammer;JN516x firmware programmer written in go
//Copyright (C) HanWool Lee <kudnya@gmail.com>
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
