package design_pattern

import "fmt"

/*
	Strategy adalah sebuah design pattern yang memungkinkan kita untuk memilih satu di antara beberapa strategy / algoritma yang tersedia selama proses runtime. Cara kerjanya dengan memisahkan strategy / algoritma dengan kelas yang mengimplementasikannya. Umumnya terdapat 3 bagian dalam design pattern ini yaitu strategy interface yang bertugas mendeklarasikan method-method apa yang perlu ada dalam strategy konkrit nantinya. Strategy konkrit adalah kelas atau struct yang akan mengimplementasikan strategy interface. Context adalah kelas atau struct yang memanggil algoritma / kelas konkrit yang memungkinkan mengimplementasikan satu diantara beberapa strategy / algoritma yang ada.
*/

/*
	Dalam kasus ini saya akan mengimplementasikan 2 contoh strategy design pattern yaitu database connection dan input validation
*/

// Strategy interface
type IDBConnection interface {
	Connect()
}

// Strategy konkrit 1
type SqlConnection struct {
	DataSource string
}
func (sc *SqlConnection) Connect() {
	fmt.Println("Connected to sql with url", sc.DataSource)
}

// Strategy konkrit 2
type MongoDBConnection struct {
	DataSource string
}
func (mc *MongoDBConnection) Connect() {
	fmt.Println("Connected to mongo with url", mc.DataSource)
}

// Context
type ConnectSpecificDB struct {
	DB IDBConnection
}

//======================================//

// Strategy interface
type IValidation interface {
	Validate() string
}

// Strategy konkrit 1
type Student struct {
	Id int `json:"id"`
	 Name string `json:"name"`
	 Age int `json:"age"`
}
func (s *Student) Validate() string {
	if s.Name == "" || s.Age == 0 {
		return "Invalid user input"
	}
	return "Successfully validate user input student"
}

// Strategy konkrit 2
type Employee struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}
func (e *Employee) Validate() string  {
	if e.Name == "" || e.Age == 0 {
		return "Invalid user input"
	}
	return "Successfully validate user input employee"
}

// Context
type ValidateSpecificOperation struct {
	Validation IValidation
}
func (vso *ValidateSpecificOperation) SetAndValidateData(validation IValidation) string {
	vso.Validation = validation
	//
	return vso.Validation.Validate()
}