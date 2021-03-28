package main

import (
	"google.golang.org/grpc"
	"grpc-client/pb"
	"grpc-client/service"
	"log"
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

	//service.GetByNo(client)
	//service.GetAll(client)
	service.UploadAvatar(client)
}
