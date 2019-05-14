package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"log"
	"rpcClient/controller"
	pb "rpcClient/pb"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
	// OpenTLS 是否开启TLS认证
	OpenTLS = true
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		//"appuid": "1000", // error
		"appuid": "100", // success
		"appkey": "i am key",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	if OpenTLS {
		return true
	}
	return false
}

func main() {
	var (
		err  error
		opts []grpc.DialOption
	)

	if OpenTLS {
		// TLS连接
		creds, err := credentials.NewClientTLSFromFile("./ssl/server.pem", "CN")
		if err != nil {
			grpclog.Fatal(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		grpclog.Fatal(err)
	}
	defer conn.Close()

	// 客户端接收服务端的流消息
	streamClient := pb.NewStreamServiceClient(conn)
	err = controller.PrintList(streamClient, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "grpc Stream Client:list", Value: 2019}})
	if err != nil {
		log.Fatal(err)
	}

	// 客户端流rpc
	err = controller.PrintRecord(streamClient, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	if err != nil {
		log.Fatal(err)
	}

	err = controller.PrintRoute(streamClient, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}

	//c := pb.NewGreeterClient(conn)
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//
	//r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("greeting:", r.Message)
	//loginReply, err := c.Login(context.Background(), &pb.LoginRequest{Name: "root", Password: "123456"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	////fmt.Println(loginReply.Message)
	////fmt.Println(loginReply.Color)
	////fmt.Println(loginReply.RewardMap)
	//fmt.Println(loginReply.DateList)
	//
	//canSetResp, err := c.CanSet(context.Background(), &pb.CanSetRequest{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(canSetResp.CanSet)
	//
	//re, err := c.CanUpdate(context.Background(), &pb.CanUpdateRequest{Username: "root"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("final=", re)
	//fmt.Println(pb.UserRegisterType_GuestRegister, pb.UserRegisterType_NormalRegister)

	//u := pb.NewUserClient(conn)
	//r, err := u.Register(context.Background(), &pb.RegisterRequest{Username: "", Password: "1qaz2wsx", Country: 1, PhoneNum: "15737345574"})
	//if err != nil {
	//	grpclog.Fatal(err)
	//	return
	//}
	//fmt.Println(r.Uid)
	//switch  pb.Profile{}.Avatar.(type){
	//case *pb.Profile_ImageData:
	//		fmt.Println("aaaaa")
	//case *pb.Profile_ImageUrl:
	//		fmt.Println("bbbbb")
	//default:
	//		fmt.Println("cccc")
	//
	//}
	//t := pb.SearchRequest{}
	//t.List = make(map[int32]*pb.ListDate)
	//l1 := &pb.ListDate{}
	//l1.List = make(map[int32]int32)
	//l1.List[1] = 1
	//l1.List[2] = 1
	//l1.List[3] = 1
	//l1.List[4] = 1
	//
	//l2 := &pb.ListDate{}
	//l2.List = make(map[int32]int32)
	//l2.List[1] = 10
	//l2.List[2] = 10
	//l2.List[3] = 10
	//l2.List[4] = 10
	//t.List = map[int32]*pb.ListDate{
	//	1: l1,
	//	2: l2,
	//}
	//fmt.Println(t.List)
	//for key, val := range t.List {
	//	fmt.Printf("key=%v,val=%v\n", key, val)
	//	for boxID, res := range val.List {
	//		fmt.Printf("----------boxID=%v,result=%v\n", boxID, res)
	//	}
	//}
}
