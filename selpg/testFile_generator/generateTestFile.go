package main

import(
	"fmt"
	"os"
)

func main(){
	filename := "test.txt"
	fout, err := os.Create(filename)
	defer fout.Close()

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 500; i++{
		fout.WriteString(fmt.Sprintf("line %d\n", i))
		if i %10 == 0{
			fout.WriteString("\f")
		}
	}
}