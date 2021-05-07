// Copyright (C) 2017 Next Thing Co. <software@nextthing.co>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>

package hostapd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	PID_FILE = "/var/run/hostapd.pid"
	CFG_FILE = "/etc/hostapd.conf"
)

func IsRunning() bool {
	pid := GetPID()
	if pid < 0 {
		return false
	}

	_, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	return true
}

func GetPID() int {
	data, err := ioutil.ReadFile(PID_FILE)
	if err != nil {
		fmt.Println("GetPID: cannot read from", PID_FILE)
		return -1
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Printf("GetPID: cannot convert %s to an integer", string(data))
		return -1
	}
	return pid
}

func Start() bool {
	fmt.Println("HostAPD: starting")
	cmd := exec.Command("/usr/sbin/hostapd", "-B", "-P", PID_FILE, CFG_FILE)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HostAPD Start: Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("HostAPD Start: Command finished with error: %v", err)
	return true
}

func Stop() bool {
	fmt.Println("HostAPD Stop: HostAPD: stopping")
	pid := GetPID()
	fmt.Println("HostAPD Stop: HostAPD's pid=", pid)
	if pid < 0 {
		return false
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(os.Interrupt)
	if err != nil {
		return false
	}
	process.Wait()
	return true
}
