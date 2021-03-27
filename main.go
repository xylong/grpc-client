package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-client/pb"
	"io"
	"log"
	"os"
	"time"
)

const (
	port = ":9000"
)

/*
func main() {
	creds,err:=credentials.NewClientTLSFromFile("server.pem", "")
	if err!=nil {
		log.Fatalln(err.Error())
	}
	options:=[]grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn,err:=grpc.Dial("127.0.0.1"+port,options...)
	if err!=nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	client:=pb.NewEmployeeServiceClient(conn)
	GetByNo(client)
}
*/

func main() {
	conn, err := grpc.Dial("127.0.0.1"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)
	//GetByNo(client)
	//GetAll(client)
	addPhoto(client)
}

func GetByNo(client pb.EmployeeServiceClient) {
	res, err := client.GetByNo(context.Background(), &pb.GetByNoRequest{No: 1990})
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(res)
	fmt.Println(res.Employee)
	fmt.Println(res.Employee.Id)
}

func GetAll(client pb.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(res.Employee)
	}
}

func addPhoto(client pb.EmployeeServiceClient) {
	imgFile, err := os.Open("beauty.jpg")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer imgFile.Close()

	md := metadata.New(map[string]string{
		"no": "2004",
	})
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)

	// 传数据
	stream, err := client.AddPhoto(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		chunk := make([]byte, 1024*128)
		chunkSize, err := imgFile.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
		// 读取到最后一块
		if chunkSize < len(chunk) {
			chunk = chunk[:chunkSize]
		}
		stream.Send(&pb.AddPhotoRequest{Data: chunk})
		time.Sleep(time.Millisecond * 500)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res.Ok)
}
