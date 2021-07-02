// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

// Example service program
//
// The program demonstrates how to create Windows service and
// install / remove it on a computer. It also shows how to
// stop / start / pause / continue any service, and how to
// write to event log. It also shows how to use debug
// facilities available in debug package.
//
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/svc"
	"luoxy.xyz/winmgr/common"
)

var cfg = common.GetCfg()
var workDir = common.GetWorkDir()

func usage(errs string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errs, os.Args[0])
	os.Exit(2)
}

func initLogger() {
	//log.SetPrefix("")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(common.NewRotatorWriter("luoxy"))
}

func main() {
	initLogger()

	svcName := cfg.Service.Name
	svcNameLong := cfg.Service.Desc

	log.Printf("服务名: %v, WorkDir: %v\n", svcName, workDir)

	isService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("failed to determine if we are running in service: %v", err)
	}
	if isService {
		// 通过services.msc进行启动
		log.Println("服务启动")
		runService(svcName, false)
		return
	}

	// 从命令行执行的操作都会走到这里
	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(svcName, true)
		return
	case "install":
		log.Println("install .....")
		err = installService(svcName, svcNameLong)
	case "remove":
		log.Println("remove .....")
		err = removeService(svcName)
	case "start":
		log.Println("Start .....")
		err = startService(svcName)
		log.Println("Start ..... End")
	case "stop":
		log.Println("Stop .....")
		err = controlService(svcName, svc.Stop, svc.Stopped)
		log.Println("Stop ..... End")
	case "pause":
		log.Println("Pause .....")
		err = controlService(svcName, svc.Pause, svc.Paused)
	case "continue":
		log.Println("Continue .....")
		err = controlService(svcName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
}
