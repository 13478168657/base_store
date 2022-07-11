/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 16:59:11
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-03 16:07:24
 * @FilePath: /allfunc/grpc_test/server/server.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"net"
	"project/allfunc/grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("no get metada")
	}
	if val, ok := md["name"]; ok {
		for key, vl := range val {
			fmt.Println(key, vl)
		}
		// fmt.Println(val)
	}
	// for key, val := range md {
	// 	fmt.Println(key, val)
	// }
	return &proto.HelloReply{Message: "hello :" + req.Name}, nil
}

func main() {

	inter := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("服务端拦截器,接收请求")

		res, err := handler(ctx, req)
		fmt.Println("执行结束")
		return res, err
	}
	opt := grpc.UnaryInterceptor(inter)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		panic("start fail" + err.Error())
	}
}
