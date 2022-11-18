package grpcclient

// import (
// 	"fmt"

// 	"github.com/barber_shop/user_service/config"
// 	pb "github.com/barber_shop/user_service/genproto"
// 	"google.golang.org/grpc"
// )

// type grpcClientI interface{
// 	OrderService() pb.OrderServiceClient
// }

// type GrpcClient struct {
// 	cfg config.Config
// 	connections map[string]interface{}
// }

// func New(cfg config.Config) (*GrpcClient, error) {
// 	connpos, err := grpc.Dial(
// 		fmt.Sprintf("%s%s", cfg.OrderServiceHost, cfg.OrderServicePort),
// 		grpc.WithInsecure(),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("post service dial host: %s port: %s err:%s",
// 	cfg.OrderServiceHost, cfg.OrderServicePort, err.Error())
// 	}
// 	return &GrpcClient{
// 		"order_service": pb.NewOrderServiceClient(connpos),
// 	}, nil
// }

// func (g *GrpcClient) OrderService() pb.OrderServiceClient {
// 	return g.connections["order_service"].(pb.OrderServiceClient)
// }
