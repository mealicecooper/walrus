package utils

import (
	"fmt"
	"regexp"
)

// ANSI color escape codes
type COLOR string

const (
    RESET  		COLOR 	=  	"\033[0m"
    RED    		COLOR 	= 	"\033[31m"
	BOLD_RED 	COLOR 	= 	"\033[38;05;196m"
    GREEN  		COLOR 	= 	"\033[32m"
    YELLOW 		COLOR 	= 	"\033[33m"
	ORANGE 		COLOR 	= 	"\033[38;05;221m"
    BLUE   		COLOR 	= 	"\033[34m"
    PURPLE 		COLOR 	= 	"\033[35m"
    CYAN   		COLOR 	= 	"\033[36m"
    WHITE  		COLOR 	= 	"\033[37m"
	GREY   		COLOR 	= 	"\033[90m"
	BOLD   		COLOR 	= 	"\033[1m"
)

func Colorize(color COLOR, text string) string {

	regex := regexp.MustCompile(`\033\[[0-9]{1,3}[0-9;]*m`)

	//panic if the color is not valid
	if !regex.MatchString(string(color)) {
		panic("Invalid color")
	}

	return fmt.Sprintf("%s%s%s", color, text, RESET)
}

func ColorPrint(color COLOR, text string) {
	fmt.Println(Colorize(color, text))
}