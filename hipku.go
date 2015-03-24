package main

import (
	"fmt"
	"regexp"
	"os"
)

func get_ip_type(ip string) int {
	ip_regexs := []string{"^(?:[0-9]{1,3}.){3}[0-9]{1,3}$", "^(?:[0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$"}
	for i := range ip_regexs {
		ip_regex := ip_regexs[i]
		matched, error := regexp.MatchString(ip_regex, ip)
		if matched == true && error == nil {
			return i
		}
	}
	return -1
}

func main() {
	ip := "192.168.0.1"
	ip_type := get_ip_type(ip)
	fmt.Println(ip_type)
	if ip_type == -1 {
		fmt.Println("Is not a ip address")
		os.Exit(255)
	}
}