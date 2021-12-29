package ex6_4
import (
	"fmt"
	"time"
)

func PrintHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func TimeNow() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC822))
}
