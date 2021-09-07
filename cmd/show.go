
package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
)

var name string
var resourceArr []string

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "미리내 개발 서버의 서비스와 자원을 확인하는 명령입니다.",
	Long: `🐧 미리내 개발 서버의 자원의 상태를 확인합니다. 배포 후 서비스 상태를 파악하기 용이합니다. 🐧

kubernetes 각 자원에 대한 상태를 확인할 수 있으며, 원한다면 자세한 내용(-d)도 확인할 수 있습니다.
CLI 내부적으로 Jenkins EC2 에 터널링하여 kubernetes 명령을 보내도록 하는 구조입니다.

 ⚓️   [resource] 의 종류는 다음과 같습니다 => po, svc, deploy, rs  [기본 값은 all 입니다]  ⚓️

사용 예
  [모든 서비스 찾기] : -a
  >  mirinae-services show -a 

  [클러스터 내부 특정 자원 찾기] : -a + -r (플래그를 여러 개 주던 하나에 , 로 나누던 알아서 잘 해석하게 해놨습니다)
  >  mirinae-services show -r deploy,po -a  
  >  mirinae-services show -r deploy -r po -a

  [개별 자원 인스턴스 찾기] : -r + -n
  >  mirinae-services show -r [resources] -n [name] 

  [개별 자원 인스턴스를 자세하게 확인] : -d + -r + -n
  >  mirinae-services show -d -r [resources] -n [name]  
`,
	Run: func(cmd *cobra.Command, args []string) {
		all,_ := cmd.Flags().GetBool("all")
		describe, _ := cmd.Flags().GetBool("describe")
		check,paramArr := checkArrParameter(3, resourceArr)
		if all  {
			sort.Strings(paramArr)
				if check == false {
					fmt.Println("🚧 Illegal Args, -r Flag usage : po | deploy | svc | rs 🚧")
					os.Exit(-1)
				} else if strings.TrimSpace(name) != "" {
					fmt.Println("🚧 Illegal Args : -n flag cannot be used with -a 🚧")
					os.Exit(-1)
				} else if describe {
					fmt.Println("🚧 Illegal Args : -d flag cannot be used with -a 🚧")
					os.Exit(-1)
				} else if contains(paramArr,"all") && len(paramArr) != 1 {
					fmt.Println("🚧 Illegal Args, -r Flag usage : po | deploy | svc | rs 🚧")
					os.Exit(-1)
				}
				//kubectl cmd resource name -> kubectl get resource
				showCmdStr := modifyShowCmdStr(all,describe, paramArr,name)
			sendCommandToJenkins(showCmdStr,false)
			os.Exit(1)
		}
		if strings.TrimSpace(name) != "" && check {
			// 하나의 자원에 하나의 이름만 검색
			fmt.Println(len(paramArr))
			if len(paramArr) != 1 {
				fmt.Println("🚧 Illegal Args, Only one argument can be used when using the -r and -n flags together 🚧")
				os.Exit(-1)
			} else if paramArr[0] == "all" {
				fmt.Println("🚧 Illegal Args, When using the -r and -n flags together, the argument of the -r flag must not be all  🚧")
				os.Exit(-1)
			}
			showCmdStr := modifyShowCmdStr(all,describe,resourceArr,name)
			sendCommandToJenkins(showCmdStr,false)
			os.Exit(1)
		}

	},
}

func modifyShowCmdStr(allFlag bool, describeFlag bool, resourceArr []string, name string) string  {
	var b bytes.Buffer
	b.WriteString("sudo kubectl ")
	if describeFlag {
		b.WriteString("describe ")
	} else {
		b.WriteString("get ")
	}
	if allFlag {
		resourceCmd := strings.Join(resourceArr,",")
		b.WriteString(resourceCmd)
	} else {
		resourceCmd := strings.Join(resourceArr,",")
		b.WriteString(resourceCmd)
		b.WriteString(" ")
		b.WriteString(name)
		b.WriteString(" ")
	}
	showCmdStr := b.String()

	fmt.Printf("\nModified Command : %s \n",showCmdStr)
	return showCmdStr
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().BoolP("all","a",false, "Find all resources supported by cli")
	showCmd.Flags().BoolP("describe","d",false,"View details about individual resources")
	showCmd.PersistentFlags().StringArrayVarP(&resourceArr, "resources","r",[]string{"all"},"The type of resource you want to check. Multiple declarations are possible, po | deploy | svc | rs is supported")
	showCmd.PersistentFlags().StringVarP(&name, "name","n","","The name of the resource to look up. Multi-flag is not supported due to log characteristics")
}
