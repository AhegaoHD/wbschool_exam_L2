package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Cannot connect: %v\n", err)
		return
	}
	fmt.Printf("Connected to %s\n", address)
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Read error: %v\n", err)
		}
		os.Exit(1) // exit if the connection is closed
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("STDIN read error: %v\n", err)
	}

	// Exit on EOF from STDIN
}
