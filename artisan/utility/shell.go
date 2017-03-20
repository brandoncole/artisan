package utility

import (
    "os/exec"
    "os"
    "fmt"
)

func Shell(cmd string, args ...string) error {

    fmt.Printf("Shell: %s %v\n", cmd, args)

    c := exec.Command(cmd, args...)
    c.Stdin = os.Stdin
    c.Stdout = os.Stdout
    c.Stderr = os.Stderr
    return c.Run()

}

func ShellCapture(cmd string, args ...string) (result string, err error) {

    fmt.Printf("ShellCapture: %s %v\n", cmd, args)

    c := exec.Command(cmd, args...)
    c.Stdin = os.Stdin
    c.Stdout = nil
    data, err := c.Output()
    if nil != err {
        return
    }
    return string(data), nil

}