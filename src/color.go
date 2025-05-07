package main

import "fmt"

func Magenta(a ...interface{}) string {
	return "\033[35m" + fmt.Sprint(a...) + "\033[0m"
}

func Bold(a ...interface{}) string {
	return "\033[1m" + fmt.Sprint(a...) + "\033[0m"
}

func Yellow(a ...interface{}) string {
	return "\033[33m" + fmt.Sprint(a...) + "\033[0m"
}

func Green(a ...interface{}) string {
	return "\033[32m" + fmt.Sprint(a...) + "\033[0m"
}

func Blue(a ...interface{}) string {
	return "\033[34m" + fmt.Sprint(a...) + "\033[0m"
}

func Cyan(a ...interface{}) string {
	return "\033[36m" + fmt.Sprint(a...) + "\033[0m"
}
