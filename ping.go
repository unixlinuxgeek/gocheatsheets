package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
    ip := os.Args[1]
    arg := "-c 3"

    app := "ping"

    cmd := exec.Command(app, ip, arg)
    stdout, err := cmd.Output()

    if err != nil {
      fmt.Println(err.Error())
      return
    }
    fmt.Println(string(stdout))
}
