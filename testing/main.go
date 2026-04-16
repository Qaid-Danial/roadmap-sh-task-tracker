package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type test struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

var (
	test_list []test
	test_list2 []test
)

func writeJson() {
	data1 := test {
		A: "alpha",
		B: "beta",
		C: "charlie",
	}
	data2 := test {
		A: "Alasdfpha",
		B: "Beasdfta",
		C: "Chaadsfrlie",
	}
	data3 := test {
		A: "aLpha",
		B: "bEta",
		C: "cHarlie",
	}
	test_list2 = append(test_list2, data1)
	test_list2 = append(test_list2, data2)
	test_list2 = append(test_list2, data3)

	marshalled, _ := json.MarshalIndent(test_list2, "", " ")

	err := os.WriteFile("test.json", marshalled, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readJson() {
	reader, _ := os.ReadFile("test.json")

	err := json.Unmarshal(reader, &test_list)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(test_list)

	for i := range test_list {
		fmt.Println(test_list[i].A)
	}
}


func writeWithoutDelJson() {
	reader, _ := os.ReadFile("test.json")

	err := json.Unmarshal(reader, &test_list)
	if err != nil {
		fmt.Print(err)
	}

	data1 := test {
		A: "alpha",
		B: "beta",
		C: "charlie",
	}
	data2 := test {
		A: "Alasdfpha",
		B: "Beasdfta",
		C: "Chaadsfrlie",
	}
	data3 := test {
		A: "aLpha",
		B: "bEta",
		C: "cHarlie",
	}

	test_list = append(test_list, data1)
	test_list = append(test_list, data2)
	test_list = append(test_list, data3)


	marshalled, _ := json.MarshalIndent(test_list, "", " ")
	err1 := os.WriteFile("test.json", marshalled, 0644)

	if err1 != nil {
		fmt.Print(err1)
	}

}



func main() {

	writeWithoutDelJson()
	readJson()
}