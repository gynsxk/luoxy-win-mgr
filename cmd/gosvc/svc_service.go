// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"fmt"
	"log"
	"time"

	"luoxy.xyz/winmgr/app"

	"github.com/pkg/errors"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var elog debug.Log

type AppService struct{}

func (m *AppService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	fasttick := time.Tick(500 * time.Millisecond)
	slowtick := time.Tick(2 * time.Second)
	tick := slowtick
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	err := app.Run()
	if err != nil {
		elog.Error(1, errors.Wrap(err, "app.Run").Error())
		return
	}

loop:
	for {
		select {
		case <-tick:
			//

		case c := <-r:
			log.Printf("收到请求指令: %d\n", c.Cmd)
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				break loop
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
				tick = slowtick
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
				tick = fasttick
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func runService(name string, isDebug bool) {
	var err error
	if isDebug {
		log.Println("从命令行执行任务")
		elog = debug.New(name)
	} else {
		log.Println("从Windows服务执行任务")
		elog, err = eventlog.Open(name)
		if err != nil {
			return
		}
	}
	defer elog.Close()

	log.Printf("%s: starting\n", name)
	elog.Info(1, fmt.Sprintf("%s: starting", name))
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &AppService{})
	if err != nil {
		log.Printf("%s: service failed %v\n", name, err)
		elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
		return
	}
	log.Printf("%s: stopped\n", name)
	elog.Info(1, fmt.Sprintf("%s: stopped", name))
}
