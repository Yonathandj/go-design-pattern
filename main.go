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

	newStudent := &design_pattern.CreateStudent{Name: "John Larren", Age: 12}
	newStudentValidation := design_pattern.ValidateSpecificOperation{ValidateData: newStudent}
	fmt.Println(newStudentValidation.ValidateData.Validate())

	updateStudent := &design_pattern.UpdateStudent{Id: 2, Name: "John Larren", Age: 20}
	updateStudentValidation := design_pattern.ValidateSpecificOperation{ValidateData: updateStudent}
	fmt.Println(updateStudentValidation.ValidateData.Validate())
}