package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "mirinae-services",
	Short: "Short ---Description",
	Long: `
🗺 개발 환경에서 배포서버로 배포했을 시, 로깅 및 서비스 상태 조회의 불편사항을 느끼시는 것을 고려하여 만든 CLI 입니다. 🗺
   1. 기본 모듈 파악 및 AWS Jenkins EC2 와의 커넥션 상태 확인
   2. MobaXterm 의 커넥션 제한으로 불편함으로 인한 터널링 지원 및 터널링 상태 확인
   3. Kubernetes 클러스터 내부 서비스 및 자원 상태 확인 [배포 후 확인 용이성]
   4. 쿠버네티스 자원에 대한 로그 지원 [배포 후 디버깅 용이성]
   
   해당 CLI 는 go 언어로 만들었습니다. 부족한 점이 많으나, 작업하시는 데에 조금이라도 도움이 되시기를 바랍니다.
                                       - 정수민 매니저 올림
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfgFile)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// 커맨드 파싱
func init() {
	cobra.OnInitialize(initConfig)

	// 플래그와 설정 세팅이 확인되는 곳

	// 코브라는 어플리케이션 전역에서 사용되는 영속 플레그를 지원하는데, 여기서 확인 된다.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "default value", "config file (default is $HOME/.mirinae-services.yaml)")

	// 코브라는 직접 호출됬을때만 사용되는 로컬 플레그도 지원한다.
	// 로컬플래그.불리언타입( 플래그이름, 짧은이름, 기본값, 설명 )
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 영속 플래그.문자열타입( 입력값, 플래그이름, 기본값, 설명 )
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".mirinae-services" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mirinae-services")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
