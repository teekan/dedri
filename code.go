package main

import (
	"log"
	"os/exec"
	"time"
)

func main(){
	chromeapp := "chromium-browser"
	chromeappArg := []string{"--headless","--proxy-server='socks5://localhost:9050'","--host-resolver-rules='MAP * ~NOTFOUND , EXCLUDE localhost'","--hide-scrollbars","--remote-debugging-port=9222","--disable-gpu","--alow-insecure-localhost"}
	cmd := exec.Command(chromeapp,chromeappArg...)
	err := cmd.Start()
	if err != nil {
		log.Println("cannot start browser", err)
	}
	log.Println(chromeappArg[1])

	time.Sleep(30 * time.Second)
	killapp := "kill"
	killappArg := []string{"$(lsof -t -i:9222)"}
	cmd = exec.Command(killapp, killappArg...)
	err = cmd.Start()
	if err !=nil {
		log.Println("cannot kill processes", err)
	}
}
