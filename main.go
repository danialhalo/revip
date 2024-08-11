package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetchAndProcessURL(ip string) {
	fmt.Println("Reverse IP Lookup on", ip)
	url := fmt.Sprintf("https://rapiddns.io/sameip/%s?full=1", ip)
	const maxRetries = 5

	var found bool
	var errorMsg string

	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(url)
		if err != nil {
			errorMsg = fmt.Sprintf("Error fetching the URL: %v", err)
			fmt.Println(errorMsg)
			time.Sleep(2 * time.Second)
			continue
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "<td>") && strings.Contains(line, ".") && !strings.Contains(line, "same ip website") {
				// Find the content between <td> and </td>
				start := strings.Index(line, "<td>")
				end := strings.Index(line, "</td>")
				if start != -1 && end != -1 {
					value := strings.TrimSpace(line[start+4 : end])
					fmt.Println(value)
					found = true
				}
			}
		}

		if err := scanner.Err(); err != nil {
			errorMsg = fmt.Sprintf("Error reading the response body: %v", err)
			fmt.Println(errorMsg)
		}

		if found {
			return
		}


		time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
	}

	if !found {
		errorMsg = "Failed to retrieve data after multiple attempts. Check your internet connection or the rapiddns site."
		fmt.Println(errorMsg)
	}
}

func isValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	privateIP := false
	if parsedIP.IsPrivate() {
		privateIP = true
	}

	return !privateIP
}

func printBanner() {
	banner := `
   ___  _____   _________    __   ____  ____  __ ____  _____
  / _ \/ __/ | / /  _/ _ \  / /  / __ \/ __ \/ //_/ / / / _ \
 / , _/ _/ | |/ // // ___/ / /__/ /_/ / /_/ / ,< / /_/ / ___/
/_/|_/___/ |___/___/_/    /____/\____/\____/_/|_|\____/_/  -: By Muhammad Danial :-

`
	fmt.Println(banner)
}

func main() {


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ip := scanner.Text()
		ip = strings.TrimSpace(ip)
		if ip != "" {
			if isValidIP(ip) {
				fetchAndProcessURL(ip)
			} else {
				fmt.Printf("Error: Provided IP address %s is not a valid public IP. Please enter a valid public IP address.\n", ip)
			}
		} else {
			printBanner()
			fmt.Println("example: echo x.x.x.x | revip")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from standard input: %v\n", err)
	}
}
