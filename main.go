package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// docker             run <cmd> <params>
// go run main.go     run <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Bad command")
	}
}

func run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"},
		os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cgrp()

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/")
	syscall.Chdir("/")
	syscall.Mount("proc", "/proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	must(syscall.Unmount("/proc", 0))
}

func cgrp() {
	cgroups := "/sys/fs/cgroup/"
	lizGroup := filepath.Join(cgroups, "liz")

	// Check if the directory exists
	if _, err := os.Stat(lizGroup); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		must(os.MkdirAll(lizGroup, 0755))
	}

	// Set the PID limit (20 processes)
	must(os.WriteFile(filepath.Join(lizGroup, "pids.max"), []byte("20"), 0700))

	// Add the current process to the cgroup
	must(os.WriteFile(filepath.Join(lizGroup, "cgroup.procs"),
		[]byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
