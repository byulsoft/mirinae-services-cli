package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	_ "go/types"
	"os"
	_ "os"
	"os/exec"
	_ "os/exec"
	"sort"
	"strconv"
	"strings"
)

var serviceNames []string
var localPort string

var tunnelCmd = &cobra.Command{
	Use:   "tunnel",
	Short: "미리내 개발 서버의 서비스와 터널링 및 터널링 상태를 확인하는 명령입니다.",
	Long: `🐧 미리내 개발 서버의 PostgreSQL, MongoDB, Redis 등의 서비스에 대한 터널링 및 터널링 연결 상태를 확인합니다. 🐧

단일 플레그는 우선 순위는 kill > status 순 이며, 단일 플래그는 이외의 플래그를 무시합니다.

사용 예 
   >  mirinae-services tunnel -s [resource] [ 서비스 터널링 및 포트포워딩 ]
   >  mirinae-services tunnel -a [ 모든 서비스 터널링 및 포트포워딩 ]
   >  mirinae-services tunnel --status [ 터널링 및 포트 포워딩 상태의 서비스 확인 ]
   >  mirinae-services tunnel --kill [ 터널링 및 포트포워딩 상태 서비스 전부 삭제 ]
`,
	Run: func(cmd *cobra.Command, args []string) {
		kill, _ := cmd.Flags().GetBool("kill")
		if kill {
			fmt.Println("🚧 The --kill tag is single flag. it has high priority. Other flags are ignored. 🚧")
			fmt.Println("🚥 Warnning : Kill all port forwarding processes. Do you want to proceed? [y/n] 🚥")
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			if strings.Trim(answer,"\n") == "y" {
				killAllPs()
				return
			} else if answer == "n" {
				os.Exit(-1)
			}
			fmt.Println("\nThe answer only recognizes y / n")
			os.Exit(-1)
		}

		status, _ := cmd.Flags().GetBool("status")
		if status {
			fmt.Println("🚧 The --status tag is single flag. it has high priority. Other flags are ignored. 🚧")
			showStatus()
			return
		}

		bool, _ := cmd.Flags().GetBool("all")

		if bool {
			tunnelingOneService([]string{"wok1","wok2","postgre","mongo","redis"},localPort)
		} else {
			checkArgs()
			tunnelingOneService(serviceNames, localPort)
		}
		fmt.Println("tunnel called : " + strings.Join(serviceNames,",") + " / "+localPort + " / "+keyPath)
	},
}

func checkArgs()  {
	if check := len(serviceNames); check > 0 && checkServiceType(serviceNames) {
		return
	}
	fmt.Println("\n🚥 Illegal Argument : Service Flag Must Be Specify 🚥\n")
	os.Exit(-1)
}

func checkServiceType(serviceNames []string) bool  {
	arr := []string{"wok1","wok2","postgre","mongo","redis"}
	sort.Strings(arr)

	for i := 0; i < len(serviceNames); i++ {
		if !contains(arr, serviceNames[i]) {
			fmt.Println("\n🚥 ServiceName usage : \"wok1\" | \"wok2\" | \"postgre\" | \"mongo\" | \"redis\" 🚥\n")
			return false
		}
	}

	return true
}

func showStatus()  {
	psCmd := exec.Command("ps","aux")
	grepCmd := exec.Command("grep","ssh -L")
	var err error
	grepCmd.Stdin, err = psCmd.StdoutPipe()
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("Error -> ["+err.Error()+"]")
		return
	}

	fmt.Println("\nRun ps aux | grep ssh Command\n")
	grepCmd.Stdout = os.Stdout
	grepCmd.Start()
	psCmd.Run()
	grepCmd.Wait()
	fmt.Println()

}
func killAllPs()  {
	fmt.Println("\n🚥 Progress Kill SSH PS 🚥\n")

	//kill -9 $(pgrep -f byul-jenkins.pem)
	findCmd := exec.Command("pgrep", "-f", "byul-jenkins.pem")
	result, _ := findCmd.Output()
	if len(result) > 0 {
		resultArr := strings.Split(string(result),"\n")
		for i := 0; i< len(resultArr); i++ {
			i, _ := strconv.Atoi(resultArr[i])
			process := os.Process{Pid: i}
			process.Kill()
		}
		fmt.Println("🏖 Kill Success -> Print result [ "+strings.Trim(strings.Join(resultArr," "),"\n")+"] 🏖\n")
	} else {
		fmt.Println("🏖 No process to kill 🏖\n")
	}
}

func init() {
	rootCmd.AddCommand(tunnelCmd)
	tunnelCmd.PersistentFlags().StringArrayVarP(&serviceNames,"service-name", "s",[]string{""}, "Tunneling Target ServiceNames enable multi value -s ~, -s ~ ..., wok1 | wok2 | postgre | mongo | redis")
	tunnelCmd.Flags().BoolP("all","a",false,"Port forwarding to all tunneling related services. -> wok1 | wok2 | postgre | mongo | redis ")
	tunnelCmd.PersistentFlags().Bool("status",false,"View all port-forwarded services [Single Flag]")
	tunnelCmd.PersistentFlags().Bool("kill",false,"Kill all port forwarding processes [Single Flag]")

}
