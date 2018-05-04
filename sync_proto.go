package main

import (
	"os/exec"
	"log"
)

func main() {
	if out, err := exec.Command("bash", "./sync_proto.sh").CombinedOutput(); err != nil {
		log.Fatalf("Failed: %s\n%s\n", err, out)
	} else {
		log.Printf("Success:\n%s\n", out)
	}
}
