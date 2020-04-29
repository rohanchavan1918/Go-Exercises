package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Antivirus implementation in Go using Virus total APi
func clearScreen() {
	// CLear the screen
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func banner() {
	fmt.Println("[+] ANTIVIRUS [+]")
	fmt.Println("USAGE :- av.exe -T dir/file/full ")
}

// Md5Hash returns the md5 of the given file
func Md5Hash(filePath string) (string, error) {
	fmt.Println("calculating MD5 for ->", filePath)
	var returnMD5String string
	file, _ := os.Open(filePath)
	// if err != nil {
	// 	return returnMD5String, err
	// }
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// SingleFileMode is for single file
func SingleFileMode(upath string) bool {
	fmt.Println(upath)
	mdhash, _ := Md5Hash(upath)
	Api(mdhash)
	return true
}

// Api sends req to the path
func Api(md5hash string) {
	const base_url string = "https://www.virustotal.com/api/v3/files/"
	const APIKey string = ""
	finalurl := base_url + md5hash
	// fmt.Println(finalurl)

	// Create a HTTP CLient
	client := &http.Client{}
	req, err := http.NewRequest("GET", finalurl, nil)
	req.Header.Set("x-apikey", APIKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

}

// VirusTotal Hit

func main() {
	clearScreen()
	// Get the type of scan the user wants to do ...
	// entire dir, single file, or entire computer
	banner()
	scantype := flag.String("type", "dir", "Enter the type of scan")
	upath := flag.String("path", "/", "Enter the file path")
	flag.Parse()
	switch *scantype {
	case "dir":
		fmt.Println("dir")
	case "file":
		fmt.Println("file")
		res := SingleFileMode(*upath)
		fmt.Println(res)
	}
}
