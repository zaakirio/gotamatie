package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

type PomodoroLog struct {
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	PomodoroDuration  int       `json:"pomodoro_duration"`
	BreakDuration     int       `json:"break_duration"`
	LongBreakDuration int       `json:"long_break_duration"`
	CompletedCount    int       `json:"completed_count"`
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "gotamatie",
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
	s.Prefix = "  "
	s.Suffix = "  "
	s.FinalMSG = "Completed!\n----------------------\n"

	start := time.Now()

	go func() {
		for {
			elapsed := time.Since(start)
			remaining := duration - elapsed
			if remaining < 0 {
				break
			}
			countdown := fmt.Sprintf("Countdown: %02d:%02d",
				int(remaining.Minutes()), int(remaining.Seconds())%60)
			s.Suffix = "  " + countdown + "  "
			time.Sleep(time.Second)
		}
	}()

	s.Start()
	time.Sleep(duration)
	s.Stop()
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

func logPomodoroSession(log PomodoroLog) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	logDir := filepath.Join(usr.HomeDir, ".gotamatie")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	logFile := filepath.Join(logDir, "pomodoro_log.json")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logEntry, err := json.Marshal(log)
	if err != nil {
		return err
	}

	if _, err := f.WriteString(string(logEntry) + "\n"); err != nil {
		return err
	}

	return nil
}

func initPomodoro(pomodoroDuration, breakDuration, longBreakDuration, pomodoroSessions int) {
	startTime := time.Now()
	completedCount := 0

	for pomodoroSessions > 0 {
		pomodoroCount := 0
		for pomodoroCount < 4 {
			setPomodoro := time.Duration(pomodoroDuration) * time.Minute
			setBreak := time.Duration(breakDuration) * time.Minute
			startTimer(setPomodoro, "Pomodoro started!")
			startTimer(setBreak, "Break started!")
			pomodoroCount++
			completedCount++
		}
		pomodoroSessions--
		fmt.Println("Pomodoro Session completed!")
		setLongBreak := time.Duration(longBreakDuration) * time.Minute
		startTimer(setLongBreak, "Long Break started!")
	}

	endTime := time.Now()
	fmt.Println("Logging Pomodoro Session!")
	log := PomodoroLog{
		StartTime:         startTime,
		EndTime:           endTime,
		PomodoroDuration:  pomodoroDuration,
		BreakDuration:     breakDuration,
		LongBreakDuration: longBreakDuration,
		CompletedCount:    completedCount,
	}

	if err := logPomodoroSession(log); err != nil {
		fmt.Println("Error logging Pomodoro session:", err)
	}
}
