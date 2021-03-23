package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
	colorBlue   string = "\033[34m"
	colorPurple string = "\033[35m"
	colorCyan   string = "\033[36m"
	colorWhite  string = "\033[37m"
)

func checkTool(toolname string, url string) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transCfg,
	}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed), toolname+" ---> "+"[Offline]", string(colorReset))
	} else if resp.StatusCode == 200 {
		fmt.Println(string(colorGreen), toolname+" ---> "+"[ACTIVE]", string(colorReset))
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		checkTool(split[0], split[1])
	}
}

func main() {
	fmt.Println(string(colorPurple), "BEGIN ........ Solo Monitor ........... BEGIN", string(colorReset))
	fmt.Println(string(colorYellow), " ---------- ST -----------", string(colorReset))
	readFile("st_tools.txt")
	fmt.Println(string(colorYellow), " ---------- ET -----------", string(colorReset))
	readFile("et_tools.txt")
	fmt.Println(string(colorYellow), " ---------- PR -----------", string(colorReset))
	readFile("pr_tools.txt")
	fmt.Println(string(colorPurple), "END ........ Solo Monitor ........... END ", string(colorReset))
}
