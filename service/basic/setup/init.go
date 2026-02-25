package setup

import "log"

func init() {
	InitViper()
	InitMySQL()
	if err := InitConsul(); err != nil {
		log.Printf("Consul 初始化失败: %v，服务将继续运行但不注册到 Consul\n", err)
	}
}
