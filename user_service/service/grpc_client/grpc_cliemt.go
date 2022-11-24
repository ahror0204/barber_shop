package grpcclient

// import (
// 	"fmt"

// 	"github.com/barber_shop/user_service/config"
// 	pb "github.com/barber_shop/user_service/genproto"
// 	"google.golang.org/grpc"
// )

type GrpcClientI interface{
	// UserService() pb.UserServiceClient

}

// type GrpcClient struct {
// 	cfg config.Config
// 	connections map[string]interface{}
// }

// func (g *GrpcClient) OrderService() pb.UserServiceClient {
// 	return g.connections["user_service"].(pb.UserServiceClient)
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
// 		"user_service": pb.NewUserServiceClient(connpos),
// 	}, nil
// }