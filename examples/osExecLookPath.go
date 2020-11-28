package main


import (
	"log"
	"fmt"
	"os/exec"
)

func main(){
	path , err := exec.LookPath("ls")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the path =%q\n", path)

	invalid , err := exec.LookPath("invalid")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the path =%q\n", invalid)
}
