package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var fileHash map[string]string

func watch(path string) {
	buf, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	h := sha256.New()
	h.Write(buf)
	bs := h.Sum(nil)

	if fileHash[path] != "" {
		if fileHash[path] != string(bs) {
			fmt.Println("GOMON::Reloading........")
			fileHash[path] = string(bs)

			cmd := exec.Command("go", "build", "-o", "temp", path)
			err = cmd.Run()
			if err != nil {
				panic(err)
			}
			var cmdExec *exec.Cmd
			cmdExec = exec.Command("./temp")

			out, err := cmdExec.Output()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Print(string(out))
		}
	} else {
		fmt.Println("GOMON::Reloading........")
		fileHash[path] = string(bs)

		cmd := exec.Command("go", "build", "-o", "temp", path)
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
		var cmdExec *exec.Cmd
		cmdExec = exec.Command("./temp")

		out, err := cmdExec.Output()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print(string(out))

	}
}

func main() {
	filesToWatch := make([]string, 0)
	fileHash = make(map[string]string)

	if cd, err := os.Getwd(); err != nil {
		panic("Error in getting cwd")
	} else {
		filesToWatch = append(filesToWatch, cd+"/main.go")
	}

	// args := os.Args[1:]

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)

		defer ticker.Stop()

		for range ticker.C {
			for _, path := range filesToWatch {
				go watch(path)
			}
		}
	}()

	for {
	}
}
