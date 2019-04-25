package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func scheduleThis(x string, y time.Duration) {
	time.Sleep(y * time.Second)
	openBrowser(x)
}

func x() {
	fmt.Println("this is function x. Not sure why it's here tbh, it's not being used.")
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("This is a simple website timer. Please type in a time in seconds.")
	var i time.Duration
	var web string
	fmt.Scan(&i)
	fmt.Println("Now type in a website (hint, start with https://)")
	fmt.Scan(&web)
	scheduleThis(web, i)
}
