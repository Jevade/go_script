package myexec

import (
	"fmt"
	"os"
	"os/exec"
)

//MyStartProcessps is My StartProcess
func MyStartProcessps() {
	env := os.Environ()
	protAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, protAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("***********ps process is %v\n", pid)
}

//MyStartProcessls is My StartProcess
func MyStartProcessls() {
	env := os.Environ()
	protAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, protAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("************ls process is %v\n", pid)
}

//MyStartProcessCMD is My StartProcess
func MyStartProcessCMD() {
	cmd := exec.Command("ls", "-l")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("**********cmd is %v\n", cmd)
}
