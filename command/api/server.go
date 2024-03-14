package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"orca-service/application/common"
	"orca-service/application/router"
	"orca-service/global"
	"orca-service/global/log"
	"orca-service/global/security"
	"orca-service/global/util"
	"os"
)

var (
	configYaml string
	Command    = &cobra.Command{Use: "server",
		Short:   "Start API server",
		Example: "orca server",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	Command.PersistentFlags().StringVarP(&configYaml, "config", "c", global.ConfigFilePath, "Start server with provided configuration file")
}

func run() error {
	file, err := os.ReadFile(configYaml)
	if err != nil {
		log.Error("error reading configuration file", err)
	}
	err = yaml.Unmarshal(file, &global.Config)
	if err != nil {
		log.Error("error parsing configuration file", err)
	}

	gin.SetMode(global.Config.Server.Mode)
	r := gin.New()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug("endpoint %v %v %v %v", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	log.Debug("server run mode: " + global.Config.Server.Mode)

	// 定义404处理路由
	r.NoRoute(router.NoRouteHandler)
	// 定义ping处理路由
	r.GET("/ping", router.PingHandler)

	if err := util.InitTranslator("zh"); err != nil {
		return err
	}

	// 初始化路由
	router.InitRouter(r)
	// 初始化Redis
	err = common.InitRedis()
	if err != nil {
		return err
	}

	// 初始化数据库
	err = common.InitDatabase()
	if err != nil {
		return err
	}

	err = security.InitSecurityEngine()
	if err != nil {
		return err
	}

	err = common.Migrate()
	if err != nil {
		log.Error("database migration failure")
		return err
	}

	address := fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port)
	return r.Run(address)
}
