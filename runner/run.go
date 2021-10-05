package runner

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func Run(info []string) {
	projName := info[0]
	projPath := path.Join(projName)
	_, err := os.Stat(projPath)

	if err != nil {
		panic("Project does not exists")
	}

	// cmd := exec.Command("cd", path.Join(projPath, "app"))
	cmd := exec.Command("air")
	cmd.Dir = path.Join(projPath, "app")

	pipe, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		panic("Was not possible to start CMD")
	}

	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')

	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}

}
