package speech

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
)

var (
	speechApi = "https://api.wit.ai/speech?v=20160526"
	cmd       = "rec"
	args      = []string{
		"-t", "wav",
		"-c", "1",
		"-",
		"rate", "16k",
		"silence", "1", "0.05", "1.5%", "2", "1.0", "2%",
	}
)

func Ai() {
	httpClient := &http.Client{}

	for {
		startAi(httpClient)
	}
}

func startAi(httpClient *http.Client) {
	pr, pw := io.Pipe()

	cmdExec := exec.Command(cmd, args...)
	stdout, err := cmdExec.StdoutPipe()

	if err != nil {
		panic("failed opening pipes")
	}

	go func() {
		io.Copy(pw, stdout)
		pw.Close()
	}()

	err = cmdExec.Start()

	if err != nil {
		panic("failed executing listener")
	}

	req, err := http.NewRequest("POST", speechApi, pr)

	if err != nil {
		panic("failed creating request")
	}

	req.Header.Set("Authorization", "Bearer "+witKey)
	req.Header.Set("Content-Type", "audio/wav")

	// Wait for first byte
	pr.Read([]byte{})

	// Send request now!
	res, err := httpClient.Do(req)

	if err != nil {
		panic("request failed")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic("failed reading response body")
	}

	fmt.Println(body)
}
