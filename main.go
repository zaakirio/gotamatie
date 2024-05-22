package main

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-pomodoro-cli",
		Short: "A CLI pomodoro timer",
		Run: func(cmd *cobra.Command, args []string) {
			pomodoroDuration, err := cmd.Flags().GetInt("d")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			breakDuration, err := cmd.Flags().GetInt("b")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			longBreakDuration, err := cmd.Flags().GetInt("l")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			pomodoroSessions, err := cmd.Flags().GetInt("s")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if !isValidInput(pomodoroDuration, breakDuration, longBreakDuration, pomodoroSessions) {
				os.Exit(1)
			}

			initPomodoro(pomodoroDuration, breakDuration, longBreakDuration, pomodoroSessions)

		},
	}
	rootCmd.Flags().IntP("d", "d", 0, "Pomodoro Duration")
	rootCmd.Flags().IntP("b", "b", 0, "Break Duration")
	rootCmd.Flags().IntP("l", "l", 0, "Long Break Duration")
	rootCmd.Flags().IntP("s", "s", 0, "Number of Pomodoro Sessions")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startTimer(duration time.Duration, message string) {
	fmt.Println(message)

	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Start()

	timer := time.NewTimer(duration)
	<-timer.C

	s.Stop()

	fmt.Println("Completed!\n----------------------")
}

func isValidInput(pomodoroDuration, breakDuration, longBreakDuration, pomodoroSessions int) bool {
	isValidPomodoro := pomodoroDuration > 1 && pomodoroDuration <= 60
	isValidBreak := breakDuration > 1 && breakDuration <= 60
	isValidLongBreak := longBreakDuration > 1 && longBreakDuration <= 60
	isValidPomodoroSessions := pomodoroSessions > 1 && pomodoroSessions <= 23

	if !isValidPomodoro {
		fmt.Println("Invalid pomodoro duration")
		return false
	}
	if !isValidBreak {
		fmt.Println("Invalid break duration")
		return false
	}
	if !isValidLongBreak {
		fmt.Println("Invalid long break duration")
		return false
	}
	if !isValidPomodoroSessions {
		fmt.Println("Invalid pomodoro sessions")
		return false
	}

	return true
}

func initPomodoro(pomodoroDuration int, breakDuration int, longBreakDuration int, pomodoroSessions int) {
	for pomodoroSessions > 0 {

		pomodoroCount := 0

		for pomodoroCount < 4 {
			setPomodoro := time.Duration(pomodoroDuration) * time.Minute
			setBreak := time.Duration(breakDuration) * time.Minute
			startTimer(setPomodoro, "Pomodoro started!")
			startTimer(setBreak, "Break started!")
			pomodoroCount++
		}

		pomodoroSessions--

		fmt.Println("Pomodoro Session completed!")

		setLongBreak := time.Duration(longBreakDuration) * time.Minute
		startTimer(setLongBreak, "Long Break started!")
	}
}
