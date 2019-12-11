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
	`[
		{
			"class":6,
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
	},
	{
		"class":6,
		"person":[{
			"name":"wangha",
			"sex":"male",
			"age":"18"
		},
		{
			"name":"zhang",
			"sex":"female",
			"age":"8"
		}]	
	}
	]`

	fmt.Printf("%+v\n", class6)
	tests := make([]test, 0)
	if err := json.Unmarshal([]byte(class6), &tests);err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%+v\n", tests[0])
		fmt.Printf("%+v\n", tests[0].Person[1])
		fmt.Printf("%s\n", tests[0].Person[0].Sex)
		fmt.Printf("%+v\n", tests[1])
	}

}