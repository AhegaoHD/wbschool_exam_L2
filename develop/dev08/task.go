package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Usage: kill PID")
				continue
			}

			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid PID: %s\n", args[1])
				continue
			}

			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Process not found: %d\n", pid)
				continue
			}

			// По умолчанию отправляется SIGTERM
			if err := process.Signal(syscall.SIGTERM); err != nil {
				fmt.Fprintf(os.Stderr, "Error killing process: %s\n", err)
				continue
			}

			fmt.Printf("Process %d killed\n", pid)

		case "ps":
			// Выполняем системную команду ps
			cmd := exec.Command("ps", "-e")
			cmd.Stdout = os.Stdout // выводим результат в STDOUT нашего шелла
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error running ps: %s\n", err)
			}

		case "exit":
			os.Exit(0)
		default:
			// Handling fork/exec
			execCmd(args)
		}
	}
}

func execCmd(args []string) {
	// Forking a new process
	binary, lookErr := exec.LookPath(args[0])
	if lookErr != nil {
		fmt.Fprintln(os.Stderr, lookErr)
		return
	}

	// Executing the command
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		fmt.Fprintln(os.Stderr, execErr)
	}
}

// Additional functions to implement 'kill' and 'ps' can be added here.
