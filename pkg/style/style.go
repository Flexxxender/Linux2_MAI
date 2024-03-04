package style

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/fatih/color"
)

func Beauty() {
	// Получение текущей директории и пользователя
	dir, err := os.Getwd()
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		return
	}
	currentUser, err := user.Current()
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		return
	}

	// Выводим имя текущего пользователя вместе с абсолютным путем
	username := color.New(color.Bold, color.FgRed).PrintfFunc()
	username(currentUser.Username)
	fmt.Print(":")

	// Выводим путь до текущей директории
	absolutePath := color.New(color.Bold, color.FgGreen).PrintfFunc()
	absolutePath(dir)
	fmt.Print("$ ")
}
