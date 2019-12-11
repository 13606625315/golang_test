package main
import (
	"encoding/json"
	"fmt"
)
type person struct {
    Name    string      `json:"name"`
    Sex     string      `json:"sex"`
    Age     string      `json:"age"`
}

type test struct {
    Class   int         `json:"class"`
    Person  []person    `json:"person"`
}



func main(){

	class6 := 
	`{  "class":6,
		"person":[{
			"name":"wangha",
			"sex":"male",
			"age":"18"
		},
		{
			"name":"zhang",
			"sex":"female",
			"age":"16"
		}]
	}`

	keys := &test{}
	if err := json.Unmarshal([]byte(class6), keys);err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%+v\n", keys)
		fmt.Printf("%+v\n", keys.Person[1])
		fmt.Printf("%s\n", keys.Person[0].Sex)
	}

}