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
    "regexp"
    "syscall"
    "bufio"
)

func main() {

    // Read layout
    config, err := readTextFile("/home/master/.config/nfetch/nfetch.conf")
    if err != nil {
        return
    }

    // Define replacements
    replacements := map[string]string{
        "<os>": getOS(),
        "<kernel>": getKernel(),
        "<distro>": getDistro(),
        "<user>": getUsername(),
        "<host>": getHostname(),
        "<cpu>": getCPU(),
        "<uptime>": getUptime(),
        "<memory>": getMemoryUsage(),
        "<packages>": getPackageCount(),
        "<flatpaks>": getFlatpakCount(),
        "<disk>": getDiskUsage("/"),
        "<time>": getCurrentTime(),
        "<shell>": getCurrentShell(),
        "<wm>": getDesktopEnvironment(),

        // foreground colors
        "<fg-black>": "\033[30m",
        "<fg-red>": "\033[31m",
        "<fg-green>": "\033[32m",
        "<fg-yellow>": "\033[33m",
        "<fg-blue>": "\033[34m",
        "<fg-magenta>": "\033[35m",
        "<fg-cyan>": "\033[36m",
        "<fg-white>": "\033[37m",

        // background colors
        "<bg-black>": "\033[40m",
        "<bg-red>": "\033[41m",
        "<bg-green>": "\033[42m",
        "<bg-yellow>": "\033[43m",
        "<bg-blue>": "\033[44m",
        "<bg-magenta>": "\033[45m",
        "<bg-cyan>": "\033[46m",
        "<bg-white>": "\033[47m",

        // text styles
        "<bold>": "\033[1m",
        "<underline>": "\033[4m",
        "<blink>": "\033[5m",
        "<inverse>": "\033[7m",
        "<reset>": "\033[0m",
    }

    // Define output
    output := config

    // Replace tags with infos
    for tag, value := range replacements {
        re := regexp.MustCompile(tag)
        output = re.ReplaceAllString(output, value)
    }

    // Print output
    fmt.Println(output)

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

func getDistro() string {
	data, err := readTextFile("/etc/os-release")
	if err != nil {
		return "Unknown Distro"
	}

	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(line[len("PRETTY_NAME="):], `"'`)
		}
	}

	return "Unknown Distro"
}

func getKernel() string {
	data, err := readTextFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "Unknown Kernel"
	}

	return strings.TrimSpace(data)
}

func getUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		return "Unknown User"
	}

	return currentUser.Username
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown Hostname"
	}

	return hostname
}

func getCPU() string {
	data, err := readTextFile("/proc/cpuinfo")
	if err != nil {
		return "Unknown CPU"
	}

	lines := strings.Split(data, "\n")
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
		return "Unknown CPU"
	}

	return fmt.Sprintf("%s (%s)", name, cores)
}

func getUptime() string {
	data, err := readTextFile("/proc/uptime")
	if err != nil {
		return "Unknown Uptime"
	}

	fields := strings.Fields(data)
	if len(fields) < 2 {
		return "Unknown Uptime"
	}

	uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return "Unknown Uptime"
	}

	uptimeDuration := time.Duration(int(uptimeSeconds)) * time.Second
	uptime := fmt.Sprintf("%dh %dm", int(uptimeDuration.Hours()), int(uptimeDuration.Minutes())%60)

	return uptime
}

func getMemoryUsage() string {
	data, err := readTextFile("/proc/meminfo")
	if err != nil {
		return "Unknown Memory Usage"
	}

	lines := strings.Split(data, "\n")
	var total, available uint64

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return "Unknown Memory Usage"
		}

		switch fields[0] {
		case "MemTotal:":
			total = value
		case "MemAvailable:":
			available = value
		}
	}

	if total == 0 {
		return "Unknown Memory Usage"
	}

	used := total - available
	usedPercentage := int((float64(used) / float64(total)) * 100)

	return fmt.Sprintf("%d%% (%dmb/%dmb)", usedPercentage, used/1024, total/1024)
}

func getPackageCount() string {
	dir := "/usr/bin/"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "Error: " + err.Error()
	}

	count := len(files)
	return fmt.Sprintf("%d", count)
}

func getFlatpakCount() string {
	dir := "/var/lib/flatpak/exports/share/"

	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return "Error: " + err.Error()
	}

	count := 0
	for _, d := range dirs {
		if d.IsDir() {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func getDiskUsage(path string) string {
	var stat syscall.Statfs_t
	syscall.Statfs(path, &stat)

	totalBytes := stat.Blocks * uint64(stat.Bsize)
	usedBytes := (stat.Blocks - stat.Bfree) * uint64(stat.Bsize)
	usedPercentage := int(float64(usedBytes) / float64(totalBytes) * 100)

	totalGB := totalBytes / (1024 * 1024 * 1024)
	usedGB := usedBytes / (1024 * 1024 * 1024)

	return fmt.Sprintf("%d%% (%dGB/%dGB)", usedPercentage, usedGB, totalGB)
}

func getCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("03:04 pm")
}

func getCurrentShell() string {
	currentUser, err := user.Current()
	if err != nil {
		return "Unknown Shell"
	}

	passwdData, err := readTextFile("/etc/passwd")
	if err != nil {
		return "Unknown Shell"
	}

	scanner := bufio.NewScanner(strings.NewReader(passwdData))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")
		if len(fields) >= 7 && fields[0] == currentUser.Username {
			shell := fields[6]
			splitPath := strings.Split(shell, "/")
			return splitPath[len(splitPath)-1]
		}
	}

	if err := scanner.Err(); err != nil {
		return "Unknown Shell"
	}

	return "Unknown Shell"
}

func getDesktopEnvironment() string {
	pid := os.Getpid()
	environFile := fmt.Sprintf("/proc/%d/environ", pid)

	data, err := ioutil.ReadFile(environFile)
	if err != nil {
		return "Unknown Desktop Environment"
	}

	environStr := string(data)
	environ := strings.Split(environStr, "\x00")

	var desktopEnv string
	for _, env := range environ {
		if strings.HasPrefix(env, "DESKTOP_SESSION=") {
			desktopEnv = strings.TrimPrefix(env, "DESKTOP_SESSION=")
			break
		}
	}

	if desktopEnv == "" {
		return "Unknown Desktop Environment"
	}

	return desktopEnv
}
