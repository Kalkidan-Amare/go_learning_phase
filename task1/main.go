package main

import "fmt"


func calculateAverage(dict map[string]int) float64 {
	total := 0
	for _, grade := range dict {
		total += grade
	}
	average := float64(total) / float64(len(dict))
	
	return average
}

func checkDataType(input interface{}, expectedType string) bool {
	switch expectedType {
	case "string":
		_, ok := input.(string)
		return ok
	case "int":
		_, ok := input.(int)
		return ok
	case "float64":
		_, ok := input.(float64)
		return ok
	default:
		return false
	}
}

func main() {
	var name string
	fmt.Println("Enter name of student: ")
	fmt.Scanln(&name)
	var no_sub int
	fmt.Println("Enter number of subjects: ")
	fmt.Scanln(&no_sub)

	dict := map[string]int{}
	var sub_name string
	var grade int
	// total := 0
	for i := 0; i < no_sub; i++ {

		fmt.Printf("Enter name of subject %d: ", i)
		fmt.Scanln(&sub_name)
	
		fmt.Printf("Enter grade of subject %d: ", i)
		fmt.Scanln(&grade)
		for !checkDataType(grade, "int") {
			fmt.Println("Invalid grade.")
			fmt.Println("Enter grade of subject")
			fmt.Scanln(&grade)
		}
	
		for grade < 0 || grade > 100 {
			fmt.Println("Invalid grade.")
			fmt.Println("Enter grade of subject 0-100")
			fmt.Scanln(&grade)
		}
		dict[sub_name] = grade
		// total += grade
	}
	
	for key, val := range dict {
		fmt.Printf("Subject: %s, Grade: %d\n", key, val)
	}
	fmt.Println("Average: ", calculateAverage(dict))
	
}