package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run image <cmd> <parameters>

// go run container.go run <cmd> <parameters>

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	fmt.Printf("Running %v as %d \n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	must(cmd.Run())

}

func child() {
	fmt.Printf("Running %v %d \n", os.Args[2:], os.Getpid())
	syscall.Sethostname([]byte("container"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := syscall.Chroot("/home/franciscod/root/rootfs"); err != nil {
		fmt.Println("Failed chroot")
	}
	must(syscall.Chdir("/"))
	fmt.Println(os.Args)
	must(cmd.Run())

}

func main() {

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()

	default:
		panic("bad command")
	}
}
