package main

import (
	"fmt"
	"os/exec"
	"time"
)

type RDPCredentials struct {
	IP       string
	Username string
	Password string
}

func main() {
	rdpCredentials := []RDPCredentials{
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
		{"Host", "Username", "Password"},
	}
	for i, rdpInfo := range rdpCredentials {
		cmd := exec.Command("cmdkey", "/generic:"+rdpInfo.IP, "/user:"+rdpInfo.Username, "/pass:"+rdpInfo.Password)
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
		rdpArgs := []string{
			"cmd", "/c", "start", "mstsc.exe",
			"/v:" + rdpInfo.IP,
		}
		cmd = exec.Command(rdpArgs[0], rdpArgs[1:]...)
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
		exec.Command("cmdkey /delete:TERMSRV/" + fmt.Sprintf("server_%d", i))
	}
}
