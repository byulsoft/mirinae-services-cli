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
πΊ κ°λ° νκ²½μμ λ°°ν¬μλ²λ‘ λ°°ν¬νμ μ, λ‘κΉ λ° μλΉμ€ μν μ‘°νμ λΆνΈμ¬ν­μ λλΌμλ κ²μ κ³ λ €νμ¬ λ§λ  CLI μλλ€. πΊ
   1. κΈ°λ³Έ λͺ¨λ νμ λ° AWS Jenkins EC2 μμ μ»€λ₯μ μν νμΈ
   2. MobaXterm μ μ»€λ₯μ μ νμΌλ‘ λΆνΈν¨μΌλ‘ μΈν ν°λλ§ μ§μ λ° ν°λλ§ μν νμΈ
   3. Kubernetes ν΄λ¬μ€ν° λ΄λΆ μλΉμ€ λ° μμ μν νμΈ [λ°°ν¬ ν νμΈ μ©μ΄μ±]
   4. μΏ λ²λ€ν°μ€ μμμ λν λ‘κ·Έ μ§μ [λ°°ν¬ ν λλ²κΉ μ©μ΄μ±]
   
   ν΄λΉ CLI λ go μΈμ΄λ‘ λ§λ€μμ΅λλ€. λΆμ‘±ν μ μ΄ λ§μΌλ, μμνμλ λ°μ μ‘°κΈμ΄λΌλ λμμ΄ λμκΈ°λ₯Ό λ°λλλ€.
                                       - μ μλ―Ό λ§€λμ  μ¬λ¦Ό
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfgFile)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// μ»€λ§¨λ νμ±
func init() {
	cobra.OnInitialize(initConfig)

	// νλκ·Έμ μ€μ  μΈνμ΄ νμΈλλ κ³³

	// μ½λΈλΌλ μ΄νλ¦¬μΌμ΄μ μ μ­μμ μ¬μ©λλ μμ νλ κ·Έλ₯Ό μ§μνλλ°, μ¬κΈ°μ νμΈ λλ€.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "default value", "config file (default is $HOME/.mirinae-services.yaml)")

	// μ½λΈλΌλ μ§μ  νΈμΆλ¬μλλ§ μ¬μ©λλ λ‘μ»¬ νλ κ·Έλ μ§μνλ€.
	// λ‘μ»¬νλκ·Έ.λΆλ¦¬μΈνμ( νλκ·Έμ΄λ¦, μ§§μμ΄λ¦, κΈ°λ³Έκ°, μ€λͺ )
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// μμ νλκ·Έ.λ¬Έμμ΄νμ( μλ ₯κ°, νλκ·Έμ΄λ¦, κΈ°λ³Έκ°, μ€λͺ )
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
