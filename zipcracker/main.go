package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cheggaaa/pb"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// func clear() {
// 	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

func extractZipWithPassword(password string, targetzip string) (found bool, pass string) {
	// clear()
	var stderr bytes.Buffer
	zip_path := targetzip
	commandString := "7z e " + zip_path + " -p" + password + " -aoa"
	commandSlice := strings.Fields(commandString)
	cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
	cmd.Stderr = &stderr
	e := cmd.Run()
	if e != nil {
		return false, ""
	}
	fmt.Println("Password is ", password)
	return true, password

}

func main() {
	// Take user Inputs, wordlist name, zip file path.
	wordlistName := flag.String("wordlist", "wordlist.txt", "Wordlist for bruteforcing")
	targetzip := flag.String("zip", "zippath", "Path of Zip file.")
	//
	flag.Parse()
	file, err := os.Open(*wordlistName)
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
	fmt.Println("[+] Bruteforcing ", *targetzip, "with wordlist ", *wordlistName)
	length := len(txtlines)
	count := length
	bar := pb.StartNew(count)
	// Bar customization
	// tmpl := `{{ red "With funcs:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`
	// start bar based on our template
	bar.Format("<.->")

	// Bar custom ends

	j := length - 1

	// Bruteforce starts here
	for i := 0; i <= length-1; i++ {
		bar.Increment()
		if i == length/2 && j == (length/2)+1 {
			break
		}
		prefix := "[" + txtlines[i] + "] #>"
		bar.Prefix(prefix)
		res, pass := extractZipWithPassword(txtlines[i], *targetzip)
		if res == true {
			fmt.Println(" Password is ", pass)
			bar.Finish()
			log.Fatal("Password Found")
			os.Exit(0)
		}
		prefix = "[" + txtlines[j] + "] #>"
		bar.Prefix(prefix)
		res, pass = extractZipWithPassword(txtlines[j], *targetzip)
		if res == true {
			fmt.Println("Password is ", pass)
			bar.Finish()
			log.Fatal("Password Found")
			os.Exit(0)
		}
		j--

	}
	print("Password Not found, Try different wordlist.")
}
