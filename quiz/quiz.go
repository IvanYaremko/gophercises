package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := flag.String("file", "problems.csv", "enter file name")

	flag.Parse()
	fmt.Println("filename: ", *fileName)

	f, err := os.Open(*fileName)

	if err != nil {
		fmt.Println("Error with reading file", err)
	}

	r := csv.NewReader(f)
	count := 0
	for {

		record, err := r.Read()

		if err == io.EOF {
			fmt.Printf("Finished! \n%v correct", count)
			break
		}

		if err != nil {
			break
		}

		question, answer := record[0], record[1]

		fmt.Printf("What is %v\n", question)
		var userAnswer string
		_, nErr := fmt.Scanln(&userAnswer)

		if nErr != nil {
			fmt.Println("reading user input error")
			break
		}

		fmt.Printf("ans: %v userAns: %v \n", answer, userAnswer)
		if answer == userAnswer {
			count++
		}
	}

}
