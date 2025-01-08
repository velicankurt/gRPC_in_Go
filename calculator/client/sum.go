/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 15:32
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient, a, b int32) {
	// Sum function
	response, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  a,
		SecondNumber: b,
	})
	if err != nil {
		return
	}

	// Print the result
	println(response.Result)
}
