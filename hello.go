package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const MONITORING = 2
const DELAY = 5

func main() {

	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs")
			showLogs()
		case 0:
			fmt.Println("Disconnecting")
			os.Exit(0)
		default:
			fmt.Println("Command not allowed")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	name := "Eduardo"
	version := 1.2
	fmt.Println("Hello", name)
	fmt.Println("The app version is", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var commandRead int
	_, err := fmt.Scan(&commandRead)
	if err != nil {
		return 0
	}
	fmt.Println("")
	fmt.Println("Chosen command:", commandRead)
	fmt.Println("")
	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitoring")

	//urls := []string{"http://wookye.com.br", "http://linkees.com.br", "https://displaynone.com.br"}

	urls := getUrls()

	for i := 0; i < MONITORING; i++ {
		for index, url := range urls {
			fmt.Println("Checking Website:", index, ":", url)
			urlStatus(url)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func urlStatus(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error occur", err)
	}

	if response.StatusCode == 200 {
		logRegister(url, true, response.StatusCode)
		fmt.Println("Website:", url, "was loaded successfully! Status Code", response.StatusCode)
	} else {
		logRegister(url, false, response.StatusCode)
		fmt.Println("Website:", url, "it has an issue. Status Code", response.StatusCode)
	}
}

func getUrls() []string {

	var urls []string

	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Error occur", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	reader := bufio.NewReader(file)

	for {
		url, err := reader.ReadString('\n')
		url = strings.TrimSpace(url)

		urls = append(urls, url)

		if err == io.EOF {
			break
		}
	}

	return urls

}

func logRegister(url string, status bool, code int) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	_, err = file.WriteString(time.Now().Format("02-01-2006 15:04:05 - ") + url + " - online: " + strconv.FormatBool(status) + " - code: " + strconv.Itoa(code) + "\n")
	if err != nil {
		fmt.Println(err)
	}

}

func showLogs() {

	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
