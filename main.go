package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-client/pb"
	"io"
	"log"
)

const (
	port=":9000"
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
	conn,err:=grpc.Dial("127.0.0.1"+port,grpc.WithInsecure())
	if err!=nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	client :=pb.NewEmployeeServiceClient(conn)
	//GetByNo(client)
	GetAll(client)
}

func GetByNo(client pb.EmployeeServiceClient)  {
	res,err:=client.GetByNo(context.Background(), &pb.GetByNoRequest{No: 1990})
	if err!=nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(res)
	fmt.Println(res.Employee)
	fmt.Println(res.Employee.Id)
}

func GetAll(client pb.EmployeeServiceClient)  {
	stream,err:=client.GetAll(context.Background(),&pb.GetAllRequest{})
	if err!=nil {
		log.Fatalln(err.Error())
	}

	for  {
		res,err:=stream.Recv()
		if err==io.EOF {
			break
		}
		if err!=nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(res.Employee)
	}
}

