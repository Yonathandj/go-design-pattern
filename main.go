package main

import (
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
}