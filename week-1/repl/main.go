package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Goodbye and happy coding!!!")
			os.Exit(0)
		}
	}()

	fmt.Print("> ")
	for scanner.Scan() {
		exp := scanner.Text()
		res, err := Evaluate(&exp)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(exp, "=", res)
		}
		fmt.Print("> ")
	}
}
