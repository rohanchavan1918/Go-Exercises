package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func extractZipWithPassword(password string, start time.Time) {
	var stderr bytes.Buffer
	zip_path := "secret.zip"
	fmt.Println("[!] Trying > ", password)
	commandString := "7z e " + zip_path + " -p" + password + " -aoa"
	commandSlice := strings.Fields(commandString)
	cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
	cmd.Stderr = &stderr
	e := cmd.Run()
	if e != nil {
		return
	}
	fmt.Println("Password is ", password)
	fmt.Println(time.Since(start))
	os.Exit(0)

}

func main() {
	start := time.Now()
	file, err := os.Open("wordlist.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	length := len(txtlines)
	j := length - 1
	for i := 0; i <= length-1; i++ {
		if i == length/2 && j == (length/2)+1 {
			break
		}
		extractZipWithPassword(txtlines[i], start)
		extractZipWithPassword(txtlines[j], start)
		j--
	}
}
