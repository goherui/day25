package setup

import (
	"day25/service/basic/config"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

var consulClient *api.Client
var serviceID string

// InitConsul 初始化Consul并注册服务
func InitConsul() error {
	// 创建Consul客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", config.GlobalConfig.Consul.Host, config.GlobalConfig.Consul.Port)

	var err error
	consulClient, err = api.NewClient(consulConfig)
	if err != nil {
		return fmt.Errorf("创建Consul客户端失败: %w", err)
	}

	// 生成服务ID
	serviceID = fmt.Sprintf("%s-%d", config.GlobalConfig.Consul.ServiceName, time.Now().Unix())

	// 注册服务 - 不使用结构体字面量，而是逐个设置字段
	registration := new(api.AgentServiceRegistration)
	registration.ID = serviceID
	registration.Name = config.GlobalConfig.Consul.ServiceName
	registration.Address = "localhost"
	registration.Port = config.GlobalConfig.Consul.ServicePort

	// 创建健康检查
	check := new(api.AgentServiceCheck)
	check.CheckID = fmt.Sprintf("%s-health", serviceID) // 明确设置CheckID
	check.Name = "TTL Health Check"
	check.TTL = fmt.Sprintf("%ds", config.GlobalConfig.Consul.TTL)
	check.DeregisterCriticalServiceAfter = "1m"

	// 添加健康检查到注册信息
	registration.Checks = []*api.AgentServiceCheck{check}

	// 注册服务
	if err := consulClient.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("注册服务失败: %w", err)
	}

	// 启动健康检查
	go func(checkID string) {
		ticker := time.NewTicker(time.Duration(config.GlobalConfig.Consul.TTL/2) * time.Second)
		defer ticker.Stop()

		log.Printf("健康检查已启动，检查ID: %s", checkID)

		for range ticker.C {
			if err := consulClient.Agent().UpdateTTL(checkID, "服务正常", api.HealthPassing); err != nil {
				log.Printf("更新健康检查失败: %v", err)
			} else {
				log.Printf("健康检查更新成功")
			}
		}
	}(fmt.Sprintf("%s-health", serviceID)) // 传递正确的checkID

	log.Println("Consul初始化成功，服务已注册")
	return nil
}

// DeregisterService 注销服务
func DeregisterService() error {
	if consulClient == nil || serviceID == "" {
		return nil
	}

	if err := consulClient.Agent().ServiceDeregister(serviceID); err != nil {
		return fmt.Errorf("注销服务失败: %w", err)
	}

	log.Println("服务已从Consul注销")
	return nil
}
