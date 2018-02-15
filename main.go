package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

// https://github.com/golang/go/wiki/WindowsDLLs

var config struct {
	ttl         time.Duration
	action      string
	blankscreen bool
	mutesound   bool
}

var user32dll = syscall.NewLazyDLL("user32.dll")
var sysSendMsg = user32dll.NewProc("SendMessageW")
var sysFindWindow = user32dll.NewProc("FindWindowW")

func main() {
	flag.DurationVar(&config.ttl, "ttl", 0, "sets the amount of time before sleeptimer executes whatever it's doing")
	flag.StringVar(&config.action, "action", "", "sets if the machine should shutdown, restart, lock or neither (default is neither)")
	flag.BoolVar(&config.blankscreen, "blank", false, "if true, blanks the screen on execution")
	flag.BoolVar(&config.mutesound, "mute", false, "if true, mutes the sound on execution")

	flag.Parse()

	if config.ttl <= time.Duration(0) {
		fmt.Fprintln(os.Stderr, "ttl must be a duration > 0")
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "sleeping for", config.ttl)
	time.Sleep(config.ttl)

	switch config.action {
	case "shutdown":
		shutdown()
		os.Exit(0)
	case "restart":
		restart()
		os.Exit(0)
	}

	if config.blankscreen {
		blankScreen()
	}

	if config.mutesound {
		muteSound()
	}
}

func shutdown() error {
	return nil
}

func restart() error {
	return nil
}

func blankScreen() error {
	ret, _, _ := sysSendMsg.Call(0xffff, 0x0112, 0xF170, 2)
	if ret != 0 {
		return errors.Errorf("non-zero return: %d", ret)
	}
	return nil
}

func muteSound() error {
	h, _, err := sysFindWindow.Call()
	if err != nil {
		return errors.Wrap(err, "get handle")
	}

	ret, _, _ := sysSendMsg.Call(uintptr(h), 0x319, uintptr(h), 0x88000)
	if ret != 0 {
		return errors.Errorf("non-zero return: %d", ret)
	}

	return nil
}
