package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main(){
	fmt.Println("Running Command for testing")
	cmd := exec.Command("ls", "-a")
	log.Print("cmd -->", cmd)



}
