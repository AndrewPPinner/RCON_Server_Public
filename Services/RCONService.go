package services

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func SSHReboot() {
	config := &ssh.ClientConfig{
		User:            "server",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SUDO_Password")),
		},
	}
	client, err := ssh.Dial("tcp", net.JoinHostPort("192.168.50.15", "469"), config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer client.Close()

	runCommand(client, "sudo reboot")
}

func runCommand(client *ssh.Client, cmd string) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	in, err := session.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	out, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var output []byte

	go func(in io.WriteCloser, out io.Reader, output *[]byte) {
		var (
			line string
			r    = bufio.NewReader(out)
		)
		for {
			b, err := r.ReadByte()
			if err != nil {
				break
			}

			*output = append(*output, b)

			if b == byte('\n') {
				line = ""
				continue
			}

			line += string(b)

			if strings.HasPrefix(line, "[sudo] password for ") && strings.HasSuffix(line, ": ") {
				_, err = in.Write([]byte(os.Getenv("SUDO_Password") + "\n"))
				if err != nil {
					break
				}
			}
		}
	}(in, out, &output)

	err = session.Run(cmd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
