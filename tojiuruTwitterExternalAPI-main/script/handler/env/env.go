package env

import (
	"fmt"
	"log"
	"os"
)

func ReadEnv(statement string) string {
	envstr := os.Getenv(statement)
	if len(envstr) < 1 {
		log.Println(fmt.Errorf("no value or key is wrong"))
		return ""
	} else if envstr == "default" {
		return "http://localhost:3000/"
	}
	return envstr
}
