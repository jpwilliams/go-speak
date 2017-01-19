package speech

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func convArgs(strArray []string) string {
	res := ""
	for x := 0; x < len(strArray); x++ {
		res += strArray[x]
	}
	return res
}

func ContinuousRecognition() {
	for {
		start()
	}
}

func start() {
	cmd2 := "rec"
	arg2 := []string{
		"-t", "wav",
		"-",
		"rate", "24k",
		"silence", "1", "0.05", "1.5%", "2", "1.0", "2%"}

	var byteArr []byte
	buf := bytes.NewBuffer(byteArr)
	fmt.Println("Executing recording...")
	cmdExec := exec.Command(cmd2, arg2...)
	fmt.Println("set up command")
	stdout, err := cmdExec.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("set up stoutpipe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("started")
	err = cmdExec.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("reading from buf")
	buf.ReadFrom(stdout)
	fmt.Println("Sending...")
	fmt.Println(SendWitBuff(buf))
}
