package concurrent

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func Command() {
	command := exec.Command("echo", "-n", "my first command")
	data, err := command.StdoutPipe()
	if err != nil {
		return
	}
	defer func() {
		if err := data.Close(); err != nil {
			fmt.Println("close err", err)
			return
		}
	}()
	var bf bytes.Buffer
	for {
		buffer := make([]byte, 50)
		n, err := data.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read task done")
				break
			} else {
				fmt.Println("read err happen:", err)
				return
			}
		}
		if n > 0 {
			bf.Write(buffer)
		}
	}
	fmt.Printf("content", bf.String())
}
