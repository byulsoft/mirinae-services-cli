package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	_ "os/exec"
	"strings"
)

var requireCmd = &cobra.Command{
	Use:   "require",
	Short: "실행 전 mirinae 개발서버 연결 상태를 체크하는 명령입니다.",
	Long: `🐧 개발 서버 인바운드 개방 상태를 충족 했는지 확인합니다. 🐧 

🚥 결과는 다음과 같습니다 🚥 
모듈 설치 여부
ssh -> [result]

⚓️ 개발 서버 연결 ⚓️
15.165.251.109 / cmd/byul-jenkins.pem / "hostname"
🏖 Connection Success -> Print result [byul-jenkins] 🏖


사용 예
  >  mirinae-services require
`,

	Run: func(cmd *cobra.Command, args []string) {
		checkRuntimeOS()
		check()
	},
}

func check() {
	fmt.Println("\n🚥 모듈 설치 여부 확인 🚥")
	command := exec.Command("which", "ssh")
	result, err := command.Output()
	if err != nil {
		fmt.Println("ssh -> None")
	} else {
		fmt.Println("ssh -> Ready [" + strings.Trim(string(result), "\n") + "]")
	}
	sendCommandToJenkins("\"hostname\"", false)
}

func init() {
	rootCmd.AddCommand(requireCmd)
}
