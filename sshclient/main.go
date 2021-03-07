package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

type Connection struct {
	*ssh.Client
	password string
}

type SshHost struct {
	IpPort   string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type HostList struct {
	Hosts []SshHost `json:"hosts"`
}

func Connect(addr, user, password string) (*Connection, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	return &Connection{conn, password}, nil

}

func (conn *Connection) SendCommands(cmds ...string) ([]byte, error) {
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	in, err := session.StdinPipe()
	if err != nil {
		fmt.Println("err in stdinPipe", err.Error())

		// log.Fatal(err)
	}

	out, err := session.StdoutPipe()
	if err != nil {
		fmt.Println("err in stdoutPipe", err.Error())
		// log.Fatal(err)
	}

	var output []byte
	_, _ = in, out

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
				_, err = in.Write([]byte(conn.password + "\n"))
				if err != nil {
					// break
					fmt.Println("ERR while writing sudo pwd - ", err.Error())
				}
			}
		}
	}(in, out, &output)

	cmd := strings.Join(cmds, "; ")
	_, err = session.Output(cmd)
	if err != nil {
		return []byte{}, err
	}

	return output, nil

}

func executeCommand(host string, username string, password string, cmd string) {
	conn, err := Connect("host", "uname", "pwd")
	// conn, err := Connect(ip, username, password)
	if err != nil {
		log.Fatal("[! Error connecting to host")
		log.Fatal(err)
	}
	// hostclr := color.New(color.FgGreen)
	red := color.New(color.FgHiWhite)
	whiteBackground := red.Add(color.BgHiBlue)
	whiteBackground.Printf("[ %s ]>", host)
	output, err := conn.SendCommands(cmd)
	if err != nil {
		fmt.Println("err in stdout", err)
	}
	fmt.Println(string(output))

}

func executeBatchCommands(cmd string, targets []SshHost) {
	for _, target := range targets {
		executeCommand(target.IpPort, target.Username, target.Password, cmd)
	}
	color.Yellow("[!] Batch Completed")
	color.Yellow("--------------------------------")
}

func getTargets(file string) []SshHost {
	var hostlist HostList
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		log.Fatal("[!] ERROR > ", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &hostlist)
	return hostlist.Hosts

}

func main() {

	file := flag.String("file", "hosts.json", "This script accepts input from a json file, please follow the structure from the temp file.")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	targets := getTargets(*file)
	cmdIp := color.New(color.FgCyan)
	for {
		cmdIp.Print("[cmd]> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "quit" {
			break
		}
		executeBatchCommands(cmd, targets)
	}

}
