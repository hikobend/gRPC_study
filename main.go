package main

import (
	"coder/pb"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}

	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read", err)
	}

	// 空のEmployee構造体を用意
	readEmployee := &pb.Employee{}

	// デシリアライズされた結果がからの構造体の中に反映される。
	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Can't read", err)
	}

	fmt.Println(readEmployee)
}
