package builder

import (
	"fmt"
	"strings"
)

func Logo() {
	var ascii_logo []string
	ascii_logo = append(ascii_logo, "\033[38;5;99m\033[1m╔═╗╦  ╔═╗╔═╗╦ ╦   ╦  ╦╦\r\n")
	ascii_logo = append(ascii_logo, "\033[38;5;98m╠╣ ║  ╠═╣╚═╗╠═╣───╚╗╔╝║\r\n")
	ascii_logo = append(ascii_logo, "\033[38;5;105m╚  ╩═╝╩ ╩╚═╝╩ ╩    ╚╝ ╩\033[0m\r\n")
	ascii_logo = append(ascii_logo, "\033[38;5;185m--\033[38;5;196m>\033[0m \033[38;5;226m@\033[38;5;128m\033[1mz3ntl3 \033[0m\033[1m(\033[38;5;165maka Efdal\033[1m)\033[0m\r\n")
	ascii_logo = append(ascii_logo, "   \033[1m\033[38;5;46m#\033[38;5;213m-- \033[1mSuper Fast Vidmoly BOT \033[1m\033[38;5;213m--\033[38;5;46m#\033[0m\r\n")
	ascii_logo = append(ascii_logo, "\t\033[1m\033[38;5;46m#\033[38;5;213m-- \033[1mBypass Vidmoly \033[1m\033[38;5;216m--\033[38;5;46m#\033[0m\r\n")


	logo := strings.Join(ascii_logo,"")

	fmt.Println(logo)
}