package main

import (
	__ "day25/proto"
	"day25/service/basic/setup"   // 直接导入，用于调用DeregisterService
	_ "day25/service/basic/setup" // 使用下划线导入，触发init()函数
	"day25/service/handler/service"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "服务端口")
)

func main() {
	log.Println("开始启动服务...")
	flag.Parse()
	log.Printf("使用端口: %d", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	log.Printf("监听成功: %v", lis.Addr())

	// 确保服务停止时注销
	defer func() {
		log.Println("开始注销服务...")
		if err := setup.DeregisterService(); err != nil {
			log.Printf("注销服务失败: %v", err)
		} else {
			log.Println("服务注销成功")
		}
	}()

	s := grpc.NewServer()
	__.RegisterStreamGreeterServer(s, &service.Server{})
	log.Printf("服务注册成功，开始监听请求...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

//package main
//
//import (
//	__ "day25/proto"
//	_ "day25/service/basic/setup"
//	"day25/service/handler/service"
//	"flag"
//	"fmt"
//	"google.golang.org/grpc"
//	"log"
//	"net"
//)
//
//var (
//	port = flag.Int("port", 50051, "The server port")
//)
//
//func main() {
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	__.RegisterStreamGreeterServer(s, &service.Server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
