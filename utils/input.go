package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetInputFromKeyboard() interface{} {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func ClearCmd() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Loading(duration time.Duration) {
	spinner := []string{"[ \\ ]", "[ / ]", "[ - ]", "[ \\ ]"}
	var i int
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		i = (i + 1) % len(spinner)
		fmt.Printf("\r%s", spinner[i])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}
