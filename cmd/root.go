package cmd

import (
	config2 "fiber-web/config"
	"fiber-web/db"
	"fiber-web/pkg/logger"
	"fiber-web/router"
	"fiber-web/util"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
)

var (
	config  string
	port    string
	mode    string
	rootCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API fiber-web",
		Example: "fiber-web server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8888", "Tcp port server listening on")
}

func run() error {
	app := router.InitRouter()
	if port != "" {
		config2.SetConfig(config, "settings.application.port", port)
	}
	go func() {
		if err := app.Listen(":" + config2.ApplicationConfig.Port); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Printf("%s Server Run http://%s:%s/ \r\n", util.GetCurrntTimeStr(),
		config2.ApplicationConfig.Host, config2.ApplicationConfig.Port)
	//关闭mongo
	defer db.DestroyMongo()
	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", util.GetCurrntTimeStr())
	if err := app.Shutdown(); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
	return nil
}

func usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	// 1. 读取配置
	config2.ConfigSetup(config)
	// 2. 数据库连接
	SetupDB()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetupDB() {
	//初始化数据库
	db.InitMongo()
}
