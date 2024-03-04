package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"second_lab/pkg/style"
	"strings"
	"syscall"
)

func main() {
	// Создаем канал для обработки сигнала CTRL+C и CTRL+Z
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGTERM)

	var processIds []int

	go func() {
		for {
			s := <-sig
			switch s {
			case syscall.SIGINT:
				kill := exec.Command("kill", fmt.Sprintf("%d", processIds[len(processIds)-1]))
				kill.Run()
				processIds = processIds[:len(processIds)-1]
			case syscall.SIGTSTP:
				for _, proc := range processIds {
					kill := exec.Command("kill", fmt.Sprintf("%d", proc))
					kill.Run()
				}
				os.Exit(0)
			}
		}
	}()

	// Запускаем бесконечный цикл для чтения ввода пользователя
	for {
		style.Beauty()

		// Считываем строку ввода пользователя
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := scanner.Text()
		if len(input) == 0 {
			continue
		}

		command := strings.Split(input, " ")
		cmd := exec.Command(command[0], command[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println("Ошибка запуска команды:", err)
			return
		}

		processId := cmd.Process.Pid
		processIds = append(processIds, processId)

		cmd.Wait()
		fmt.Printf("Процесс %s с номером %d завершен\n", command[0], processId)
	}
}
