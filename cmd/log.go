package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"time"
)

var podName string
var tail int8
var since string

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "미리내 개발 서버의 pod 로그를 확인하는 명령입니다.",
	Long: `🐧 미리내 개발 서버의 pod 로그를 확인하는 명령입니다. 배포 한 서비스의 로그 사항을 파악하기 용이합니다. 🐧

⚓️   먼저 show 명령에서 찾을 pod 의 이름을 확인하시고 진행하시길 바랍니다. ⚓️
      > mirinae-services show -r po -a

사용 예
  [pod 의 모든 로그 찾기] 
  >  mirinae-services log -n [name]

  [pod 의 로그 10 줄 찾기] 
  >  mirinae-services log -n [name] -t 10

  [pod 의 10 분 전 로그 찾기]
  >  mirinae-services log -n [name] -s 10m

  [pod 의 1 시간 전 로그 20 줄 찾기]
  >  mirinae-services log -n [name] -s 1h -t 20
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.PersistentFlags().GetString("pod-name")
		tail, _ := cmd.PersistentFlags().GetInt8("tail")
		since, _ := cmd.PersistentFlags().GetString("since")
		follow, _ := cmd.Flags().GetBool("follow")

		name = strings.TrimSpace(name)
		since = strings.TrimSpace(since)

		fmt.Printf("pod-name : '%s' , tail : '%d' , since : '%s' \n", name, tail, since)

		if strings.Trim(name, "\n") == "" {
			fmt.Println("🚧 Illegal Args, -n Flag must be specify 🚧")
			os.Exit(-1)
		} else if tail < -1 {
			fmt.Println("🚧 Illegal Args, -t Flag cannot be less than -1 🚧")
			os.Exit(-1)
		} else if _, err := time.ParseDuration(since); err != nil {
			fmt.Println("🚧 Illegal Args, -s Flag must follow the duration syntax 🚧")
			os.Exit(-1)
		}

		logCmdStr := modifyLogCmdStr(follow, name, tail, since)
		sendCommandToJenkins(logCmdStr, follow)
		os.Exit(1)
	},
}

func modifyLogCmdStr(follow bool, name string, tail int8, since string) string {
	var result string
	var b bytes.Buffer

	if strings.HasPrefix(name, "redis") {

		b.WriteString("sudo kubectl exec -t ")
		b.WriteString(name)
		b.WriteString(" -- tail")
		if follow {
			b.WriteString(" --follow")
		}
		b.WriteString(" /data/stdout")
		result = b.String()
		fmt.Printf("\nModified Command : %s \n", result)
		return result
	}
	b.WriteString("sudo kubectl logs ")
	b.WriteString(name)
	b.WriteString(" --tail ")
	b.WriteString(strconv.Itoa(int(tail)))
	b.WriteString(" --since ")
	b.WriteString(since)

	if follow {
		b.WriteString(" --follow")
	}
	result = b.String()
	fmt.Printf("\nModified Command : %s \n", result)
	return result
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed")
	logCmd.PersistentFlags().StringVarP(&podName, "pod-name", "n", "", "The name of the Pod resource for which to view logs.")
	logCmd.PersistentFlags().Int8VarP(&tail, "tail", "t", -1, "The number of output lines in the log. The default is all.")
	logCmd.PersistentFlags().StringVarP(&since, "since", "s", "0", "Only return logs newer than a relative duration like 5s, 2m, or 3h, default value is all")
}
