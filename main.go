package main

import (
	"fmt"

	"github.com/Yonathandj/go-design-pattern/design_pattern"
)

func main() {
	// System Design Pattern

	sqlDB := &design_pattern.SqlConnection{DataSource: "http://localhost/5432"}
	sqlDBConnection := design_pattern.ConnectSpecificDB{DB: sqlDB}
	sqlDBConnection.DB.Connect()

	mongoDB := &design_pattern.MongoDBConnection{DataSource: "http://localhost/8080"}
	MongoDBConnection := design_pattern.ConnectSpecificDB{DB: mongoDB}
	MongoDBConnection.DB.Connect()

	//=========================================//

	newStudent := &design_pattern.Student{Id: 1, Name: "Kayla Larren", Age: 20}
	newStudentValidation := design_pattern.ValidateSpecificOperation{}
	fmt.Println(newStudentValidation.SetAndValidateData(newStudent))

	newEmployee := &design_pattern.Employee{Id: 1, Name: "John Larren", Age: 20}
	newEmployeeValidation := design_pattern.ValidateSpecificOperation{}
	fmt.Println(newEmployeeValidation.SetAndValidateData(newEmployee))

	//=========================================//
	//=========================================//

	// Observer Design Pattern

	nathan := design_pattern.Subscriber{Name: "Nathan"}
	larren := design_pattern.Subscriber{Name: "Larren"}

	mrBeastChannel := design_pattern.YoutubeChannel{Name: "Mr.Beast Channel"}

	nathan.Subscribe(&mrBeastChannel)
	larren.Subscribe(&mrBeastChannel)

	mrBeastChannel.NotifyLive()

	fmt.Print("Larren has subscribed to", larren.YoutubeChannels)
	for _, v := range larren.YoutubeChannels {
		fmt.Print(v)
	}

	fmt.Println()
	
	fmt.Print("Nathan has subscribed to", nathan.YoutubeChannels)
	for _, v := range larren.YoutubeChannels {
		fmt.Print(v)
	}
	
	fmt.Println()
	
	fmt.Print("Mr.Beast Channel has subscribers", mrBeastChannel.Subscribers)
	for _, v := range mrBeastChannel.Subscribers {
		fmt.Print(v)
	}
}