package cmd

import (
	"bufio"
	"fmt"
	_ "github.com/spf13/cobra"
	"os"
	"os/exec"
	_ "os/exec"
	"runtime"
	"sort"
	"strings"
)

var keyPath = os.Getenv("MIRINAE_PATH") + "byul-jenkins.pem"

const jenkinsEc2 = "15.165.251.109"

var workerNodeIP = "10.0.101.205"
var UserOs OS

type CmdType int

const (
	ROOT CmdType = iota
	REQUIRE
	TUNNEL
	SHOW
	LOG
)

type OS int

const (
	WINDOWS OS = iota
	MAC
	LINUX
)

func checkRuntimeOS() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("🖼 Runtime Os > Windows")
		UserOs = 0
	case "darwin":
		fmt.Println("🍏 Runtime Os > MAC operating system")
		UserOs = 1
	case "linux":
		fmt.Println("🐧 Runtime Os > Linux")
		UserOs = 2
	default:
		fmt.Printf("%s.\n", os)
		fmt.Println("CLI Do Not Support your Operation System")
	}
}

func checkArrParameter(cmdType CmdType, paramArr []string) (bool, []string) {
	paramArr = strings.Split(strings.Join(paramArr, ","), ",")
	fmt.Printf("check > %s\n", paramArr)

	var checkArr []string
	switch cmdType {
	case ROOT:
		return false, paramArr

	case REQUIRE:
		return false, paramArr

	case TUNNEL:
		checkArr = []string{"wok1", "wok2", "postgre", "mongo", "redis"}

	case SHOW, LOG:
		checkArr = []string{"po", "deploy", "svc", "rs", "all"}

	}
	sort.Strings(checkArr)
	sort.Strings(paramArr)
	for i := 0; i < len(paramArr); i++ {
		result := contains(checkArr, paramArr[i])
		if result == false {
			return false, paramArr
		}
	}
	return true, paramArr
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func sendCommandToJenkins(cmd string, syncCmd bool) string {

	fmt.Println("\n⚓️ 개발 서버 연결 ⚓️")
	fmt.Printf("Call sendCommandToJenkins : Target > %s , KeyPath > %s, Cmd > %s\n\n", jenkinsEc2, keyPath, cmd)
	//ssh -i ${MIRINAE_PATH}/byul-jenkins.pem ec2-user@15.165.251.109 -oStrictHostKeyChecking=no -t "hostname"
	command := exec.Command("ssh", "-i", keyPath, "ec2-user@"+jenkinsEc2, "-oStrictHostKeyChecking=no", "-t", cmd)

	if syncCmd {
		fmt.Println("Test")
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			fmt.Println("Connection Fail -> [" + err.Error() + "]")
			os.Exit(-1)
		}
		os.Exit(1)
	}

	result, err := command.Output()
	if err != nil {
		fmt.Println("Connection Fail -> [" + err.Error() + "]")
		os.Exit(-1)
	} else {
		fmt.Printf("🏖 Connection Success -> Print result \n %s\n-End 🏖\n", strings.Trim(string(result), "\n"))
	}
	return string(result)
}

func tunnelingOneService(serviceNames []string, localPort string) {
	fmt.Println("\n⚓️ 개발 서버 연결 및 터널링 ⚓️")
	fmt.Println("⛵️ Target : " + strings.Join(serviceNames, ",") + " -> Binding to Localhost : " + localPort + " with " + keyPath + " ⛵️\n")

	localhostForwardFromTunnelingService(serviceNames)
}

func localhostForwardFromTunnelingService(serviceNames []string) {
	sort.Strings(serviceNames)

	for i := 0; i < len(serviceNames); i++ {
		service := serviceNames[i]
		if service == "wok2" {
			workerNodeIP = "10.0.102.169"
		}
		portNum := connectionKubeApiServerAndGetServiceNodePort(service)
		readClientLine(service, portNum, workerNodeIP)
	}
}

func readClientLine(service string, portNum string, workerNodeIP string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Service Target : " + service + " -> " + workerNodeIP + " : " + portNum)
	fmt.Print("What local port do you want?: ")
	localport, _ := reader.ReadString('\n')
	fmt.Println("select localPort : " + localport)
	localPortForwarding(strings.Trim(localport, "\n"), workerNodeIP, portNum)

}

func localPortForwarding(localport string, workerNodeIP string, portNum string) {
	//ssh -L 2000:10.0.101.205:22 ec2-user@15.165.251.109 -i ~/Documents/AWS/byul/byulsoft-pemkey/byul-jenkins.pem -fNT
	//ssh -O exit -L 2000:10.0.101.205:22 ec2-user@15.165.251.109 -i ~/Documents/AWS/byul/byulsoft-pemkey/byul-jenkins.pem -fN
	//ps aux | grep ssh
	//fmt.Println(localport+" / "+workerNodeIP+" / "+portNum+" / "+jenkinsEc2+" / "+keyPath)
	command := exec.Command("ssh", "-L", localport+":"+workerNodeIP+":"+portNum, "ec2-user@"+jenkinsEc2, "-i", keyPath, "-fNT")
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		fmt.Println("chmod 400 -> Fail [" + err.Error() + "]")
		fmt.Println("🚧 Fail SSL Port Forward, Stop Process 🚧")
		os.Exit(1)
	}
	fmt.Println("\nSSL Port Forward -> OK.\n✈️ Connected Your Local Machine...✈️")
	fmt.Println("See SSL Information, run ps aux | grep ssh Command\n")
}

func connectionKubeApiServerAndGetServiceNodePort(service string) string {
	cmd := ""
	serviceNodePort := ""
	if service == "wok1" || service == "wok2" {
		return "22"
	} else {
		cmd = "sudo kubectl get svc name -o jsonpath='{range .items[*]}{.spec.ports[].nodePort}'"
	}

	switch service {
	case "postgre":
		{
			cmd = strings.Replace(cmd, "name", "postgres", len(cmd))
			fmt.Println(cmd)
			serviceNodePort = sendCommandToJenkins(cmd, false)
			break
		}
	case "mongo":
		{
			cmd = strings.Replace(cmd, "name", "mongodb", len(cmd))
			fmt.Println(cmd)
			serviceNodePort = sendCommandToJenkins(cmd, false)
			break
		}
	case "redis":
		{
			cmd = strings.Replace(cmd, "name", "redis-leader", len(cmd))
			fmt.Println(cmd)
			serviceNodePort = sendCommandToJenkins(cmd, false)
			break
		}
	}
	return serviceNodePort
}
