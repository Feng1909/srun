package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	log2 "log"
	"net/http"
	"os"
	"time"
)

const (
	Version = "v0.1.23"
	Timeout = 3 * time.Second
)

var LogLevel = log.InfoLevel

var CommandMap = map[string]Func{
	"config": DefaultClient.SetAccount,
	"login":  DefaultClient.Login,
	"logout": DefaultClient.Logout,
	"info":   DefaultClient.GetInfo,
	"update": DefaultClient.Update,

	"help": DefaultClient.CmdHelp,
}

var optionDocs = map[string]string{
	"-d": "Show debug message",
	"-v": "Print version information and quit",
	"-h": "Show help",
}

var cmdDocs = map[string][]string{
	"config": {"srun config", "Set Username and Password"},
	"login":  {"srun [login]", "Login Srun"},
	"logout": {"srun logout", "Logout Srun"},
	"info":   {"srun info", "Get Srun Info"},
	"update": {"srun update", "Update srun"},
}

func main() {
	var debugMode bool
	var helpMode bool
	var versionMode bool

	flag.BoolVar(&debugMode, "d", false, "debug mode")
	flag.BoolVar(&helpMode, "h", false, "show help")
	flag.BoolVar(&versionMode, "v", false, "show version")

	flag.Parse()

	var cmd string
	var params []string

	args := flag.Args()
	if len(args) > 0 {
		cmd = args[0]
		params = args[1:]
	} else {
		cmd = "login"
	}

	switch {
	case helpMode:
		DefaultClient.CmdHelp(cmd, args...)
		return
	case versionMode:
		DefaultClient.ShowVersion()
		return
	case debugMode:
		LogLevel = log.DebugLevel
	}

	// config
	http.DefaultClient.Timeout = Timeout
	log2.SetOutput(nil)
	log.SetOutput(os.Stdout)
	log.SetLevel(LogLevel)
	log.SetFormatter(&log.TextFormatter{
		//DisableTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	if handle, ok := CommandMap[cmd]; ok {
		handle(cmd, params...)
	} else {
		DefaultClient.CmdHelp(cmd, params...)
	}

}
