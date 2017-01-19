package speech

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func ContinuousRecognition() {
	for {
		start()
	}
}

func start() {
	cmd := "rec"

	args := []string{
		"-t", "wav",
		"-c", "1",
		"-",
		"rate", "16k",
		"silence", "1", "0.05", "1.5%", "2", "1.0", "2%",
	}

	var byteArr []byte
	buf := bytes.NewBuffer(byteArr)

	cmdExec := exec.Command(cmd, args...)
	stdout, err := cmdExec.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	err = cmdExec.Start()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("listening")
	buf.ReadFrom(stdout)

	fmt.Println("analysing sample")
	fmt.Println(SendWitBuff(buf))
}
