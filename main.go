package main

import (
	"fmt"

	"github.com/Yonathandj/go-design-pattern/design_pattern"
)

func main() {
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
}