package main

import (
    "fmt"
    "strings"
    "strconv"
    "os/user"
    "os"
    "runtime"
    "time"
    "io/ioutil"
)


func main() {

    // Chooose an accent color
    primary := "\033[34m"
    secondary := "\033[36m"
    tertiary := "\033[37m"
    reset  := "\033[0m"

    // Choose a seperator and it's color
    sep := tertiary + " ~ "

    // Get distro name
    distro, err := getDistroName()
	if err != nil {
		return
	}

    // Get kernel
    kernel, err := getKernelName()
    if err != nil {
        return
    }

    // Get username
    user, err := getUsername()
    if err != nil {
        return
    }

    // Get hostname
    hostname, err := getHostname()
    if err != nil {
        return
    }

    // Get cpu
    cpu, err := getCPUInfo()
    if err != nil {
        return
    }

    // Get uptime
    uptime, err := getUptime()
    if err != nil {
        return
    }

    // Get memory
    memory, err := getMemoryUsage()
    if err != nil {
        return
    }
    
    // Get art
    artPath := "/home/" + user + "/.config/nfetch/art.txt"
    art, err := readTextFile(artPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
    }

    // Split art into its lines
    artLines := strings.Split(art, "\n")

    // Add color reset and spacing
	for i, line := range artLines {
        // You can replace reset with a color if you like
		artLines[i] = reset + line + "  "
	}

    // User@Host
    println(artLines[0] + primary + user + tertiary + "@" + secondary + hostname)

    // Os
    println(artLines[1] + primary + "os     " + sep + secondary + distro + reset)

    // Kernel
    println(artLines[2] + primary + "kernel " + sep + secondary + kernel + reset)

    // Cpu
    println(artLines[3] + primary + "cpu    " + sep + secondary + cpu + reset)

    // Uptime
    println(artLines[4] + primary + "uptime " + sep + secondary + uptime + reset)

    // Memory
    println(artLines[5] + primary + "memory " + sep + secondary + memory + reset)

    // Print the remaining lines
	for i := 6; i < len(artLines); i++ {
		fmt.Println(artLines[i])
	}
}

func readTextFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getOS() string {
    return runtime.GOOS
}

func getDistroName() (string, error) {
	data, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(line[len("PRETTY_NAME="):], `"'`), nil
		}
	}

	return "", fmt.Errorf("Distro name not found in /etc/os-release")
}

func getKernelName() (string, error) {
	data, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func getUsername() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.Username, nil
}

func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return hostname, nil
}

func getCPUInfo() (string, error) {
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var name, cores string

	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) == 2 {
			key := strings.TrimSpace(fields[0])
			value := strings.TrimSpace(fields[1])
			switch key {
			case "model name":
				name = value
			case "cpu cores":
				cores = value
			}
		}
	}

	if name == "" || cores == "" {
		return "", fmt.Errorf("CPU info not found in /proc/cpuinfo")
	}

	return fmt.Sprintf("%s (%s)", name, cores), nil
}

func getUptime() (string, error) {
	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 2 {
		return "", fmt.Errorf("Invalid format in /proc/uptime")
	}

	uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return "", err
	}

	uptimeDuration := time.Duration(int(uptimeSeconds)) * time.Second
	uptime := fmt.Sprintf("%dh %dm", int(uptimeDuration.Hours()), int(uptimeDuration.Minutes())%60)

	return uptime, nil
}

func getMemoryUsage() (string, error) {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var total, available uint64

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return "", err
		}

		switch fields[0] {
		case "MemTotal:":
			total = value
		case "MemAvailable:":
			available = value
		}
	}

	if total == 0 {
		return "", fmt.Errorf("Failed to read total memory from /proc/meminfo")
	}

	used := total - available
	usedPercentage := int((float64(used) / float64(total)) * 100)

	return fmt.Sprintf("%d%% (%dmb/%dmb)", usedPercentage, used/1024, total/1024), nil
}

