package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mattn/go-tty"
)

func main() {
	var amount int

	// Clear screen and move cursor to top-left
	fmt.Print("\033[2J") // Clear screen
	fmt.Print("\033[H")  // Move cursor to row 1, col 1

	// Print frame
	fmt.Println("Game Simulation Program")
	fmt.Println("Programmer Gadungan")
	fmt.Println()
	fmt.Print("Set initial Amount: ")
	fmt.Scanln(&amount)
	fmt.Println()
	fmt.Println("Press A to Increase amount by 100")
	fmt.Println("Press S to Decrease amount by 100")
	fmt.Println("Press Q to Quit")
	fmt.Println()
	fmt.Printf("Amount: %d\n", amount)
	fmt.Print("What's your action: ")

	amountRow := 10
	promptRow := amountRow + 1

	// Open terminal for single-key input (cross-platform)
	tty, err := tty.Open()
	if err != nil {
		fmt.Println("Failed to open tty:", err)
		os.Exit(1)
	}
	defer tty.Close()

	// Handle Ctrl+C safely
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		// Clear screen and restore terminal
		fmt.Print("\033[H\033[2J")
		fmt.Println("Exiting...")
		os.Exit(0)
	}()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			continue
		}

		switch r {
		case 'A', 'a':
			amount += 100
		case 'S', 's':
			amount -= 100
		case 'Q', 'q':
			fmt.Printf("\033[%d;1H\nExiting...\n", promptRow+1)
			return
		default:
			continue
		}

		// Update Amount line in place
		fmt.Printf("\033[%d;1H\033[2KAmount: %d", amountRow, amount)
		// Update prompt line in place
		fmt.Printf("\033[%d;1H\033[2KWhat's your action: ", promptRow)
	}
}
