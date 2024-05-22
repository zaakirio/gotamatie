

# Gotamatie 🍅🐇

Gotamatie is a command-line interface (CLI) application written in Go that implements the Pomodoro Technique, a time management method that alternates between focused work periods (called "pomodoros") and short breaks. This application allows you to customize the duration of pomodoros, short breaks, long breaks, and the number of pomodoro sessions.

## Installation

To install Gotamatie, you'll need to have Go installed on your system. You can download Go from the official website: https://golang.org/dl/

Once you have Go installed, you can clone the repository and build the application:

```bash
git clone https://github.com/zaakirio/gotamatie.git
cd go-pomodoro-cli
go build
```

This will create an executable file `Gotamatie` (or `Gotamatie.exe` on Windows) in the same directory.

## Usage

To run Gotamatie, use the following command:

```bash
./gotamatie --d <pomodoro-duration> --b <break-duration> --l <long-break-duration> --s <pomodoro-sessions>
```

Replace the placeholders with the desired values:

- `<pomodoro-duration>`: The duration of each pomodoro in minutes (must be between 2 and 60).
- `<break-duration>`: The duration of each short break in minutes (must be between 2 and 60).
- `<long-break-duration>`: The duration of each long break in minutes (must be between 2 and 60).
- `<pomodoro-sessions>`: The number of pomodoro sessions (must be between 2 and 23).

For example, to start a pomodoro session with 25-minute pomodoros, 5-minute short breaks, 15-minute long breaks, and 4 pomodoro sessions, you would run:

```bash
./gotamatie --d 25 --b 5 --l 15 --s 4
```

The application will display a spinner animation and the remaining time for each pomodoro, short break, and long break.

## Future Features
- Progress tracking in local storage
- Productivity analytics
- Graphs

