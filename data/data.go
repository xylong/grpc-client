package data

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-client/pb"
	"time"
)

var Employees = []pb.Employee{
	{
		Id:        3,
		No:        3,
		FirstName: "佟",
		LastName:  "丽娅",
		MonthSalary: &pb.MonthSalary{
			Basic: 50000,
			Bouns: 4589.58,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
	{
		Id:        4,
		No:        4,
		FirstName: "迪丽",
		LastName:  "热巴",
		MonthSalary: &pb.MonthSalary{
			Basic: 56874,
			Bouns: 123.45,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
}
