package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-pomodoro-cli",
		Short: "A CLI pomodoro timer",
		Run: func(cmd *cobra.Command, args []string) {
			pomodoroDuration, _ := cmd.Flags().GetString("d")
			breakDuration, _ := cmd.Flags().GetString("b")
			longBreakDuration, _ := cmd.Flags().GetString("l")
			pomodoroSessions, _ := cmd.Flags().GetString("s")

			initPomodoro(pomodoroDuration, breakDuration, longBreakDuration, pomodoroSessions)

		},
	}
	rootCmd.Flags().String("d", "", "Pomodoro Duration")
	rootCmd.Flags().String("b", "", "Break Duration")
	rootCmd.Flags().String("l", "", "Long Break Duration")
	rootCmd.Flags().String("s", "", "Number of Pomodoro Sessions")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startPomodoro(duration time.Duration) {
	fmt.Println("Pomodoro started!")
	time.Sleep(duration)
	fmt.Println("Pomodoro completed!")
}

func startBreak(duration time.Duration) {
	fmt.Println("Break started!")
	time.Sleep(duration)
	fmt.Println("Break completed!")
}

func startLongBreak(duration time.Duration) {
	fmt.Println("Long Break started!")
	time.Sleep(duration)
	fmt.Println("Long Break completed!")
}

func initPomodoro(pomodoroDuration int, breakDuration int, longBreakDuration int, pomodoroSessions int) {
	isValidPomodoro := pomodoroDuration > 1 || pomodoroDuration <= 60
	println(pomodoroDuration)
	isValidBreak := breakDuration > 1 || breakDuration <= 60
	isValidLongBreak := longBreakDuration > 1 || longBreakDuration <= 60
	isValidPomodoroSessions := pomodoroSessions > 1 || pomodoroSessions <= 23

	if !isValidPomodoro {
		fmt.Println("Invalid pomodoro duration")
		return
	}
	if !isValidBreak {
		fmt.Println("Invalid break duration")
		return
	}
	if !isValidLongBreak {
		fmt.Println("Invalid long break duration")
		return
	}
	if !isValidPomodoroSessions {
		fmt.Println("Invalid pomodoro sessions")
		return
	}

	for pomodoroSessions > 0 {

		pomodoroCount := 0

		for pomodoroCount < 4 {
			setPomodoro := time.Duration(pomodoroDuration) * time.Minute
			println(setPomodoro)
			setBreak := time.Duration(breakDuration) * time.Minute
			startPomodoro(setPomodoro)
			startBreak(setBreak)
			pomodoroCount++
		}

		pomodoroSessions--

		fmt.Println("Pomodoro Session completed!")

		setLongBreak := time.Duration(longBreakDuration) * time.Minute
		startLongBreak(setLongBreak)
	}

}
