package cmd

import (
	"fmt"
	"github.com/is-a-gamer/file-service/http"
	"github.com/is-a-gamer/file-service/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&confPath, "config", "c", "config.yaml", "config file (default is ./config.yaml)")
}

var (
	confPath string
	rootCmd  = &cobra.Command{
		Use:   `fileserver`,
		Short: ``,
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			// 初始化日志配置
			log.InitLogger("file-server")
			// 初始化配置文件
			initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			http.Start()
		},
	}
)

func initConfig() {
	if confPath != "" {
		viper.SetConfigFile(confPath)
	} else {
		viper.SetConfigFile("config.yaml")
	}
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Error("read config error", zap.String("confFile", confPath), zap.Error(err))
	} else  {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
