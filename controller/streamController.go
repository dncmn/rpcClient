package controller

import (
	"context"
	"io"
	"log"
	pb "rpcClient/pb"
)

// 服务端流rpc
func PrintList(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pg.name=%s,pg.value=%v\n", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

// 客户端流rpc
func PrintRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	return nil
}

// 双向流rpc
func PrintRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {

		// 客户端发送
		err = stream.Send(r)
		if err != nil {
			return err
		}

		// 客户端接收流
		resp, err := stream.Recv()
		if err == io.EOF {
			return err
		}

		if err != nil {
			return err
		}

		log.Printf("resp: pj.Name=%v,pt.Value=%v\n", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}
