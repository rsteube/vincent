package theme

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	Name    string `yaml:"name"`
	Black   string `yaml:"color_01"`
	Red     string `yaml:"color_02"`
	Green   string `yaml:"color_03"`
	Yellow  string `yaml:"color_04"`
	Blue    string `yaml:"color_05"`
	Magenta string `yaml:"color_06"`
	Cyan    string `yaml:"color_07"`
	White   string `yaml:"color_08"`

	BrightBlack   string `yaml:"color_09"`
	BrightRed     string `yaml:"color_10"`
	BrightGreen   string `yaml:"color_11"`
	BrightYellow  string `yaml:"color_12"`
	BrightBlue    string `yaml:"color_13"`
	BrightMagenta string `yaml:"color_14"`
	BrightCyan    string `yaml:"color_15"`
	BrightWhite   string `yaml:"color_16"`

	Background string `yaml:"background"`
	Foreground string `yaml:"foreground"`

	Cursor string `yaml:"cursor"`
}

func (t Theme) Render() (s string) {
	s += fmt.Sprintf("%v\n", t.render("                                                    "))
	s += fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v\n", t.render("  "), t.block(t.Black), t.block(t.Red), t.block(t.Green), t.block(t.Yellow), t.block(t.Blue), t.block(t.Magenta), t.block(t.Cyan), t.block(t.White), t.render("  "))
	s += fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v\n", t.render("  "), t.block(t.BrightBlack), t.block(t.BrightRed), t.block(t.BrightGreen), t.block(t.BrightYellow), t.block(t.BrightBlue), t.block(t.BrightMagenta), t.block(t.BrightCyan), t.block(t.BrightWhite), t.render("  "))
	s += fmt.Sprintf("%v\n", t.render("                                                    "))
	s += fmt.Sprintf("%v%v\n", t.text(t.Black), t.text(t.BrightBlack))
	s += fmt.Sprintf("%v%v\n", t.text(t.Red), t.text(t.BrightRed))
	s += fmt.Sprintf("%v%v\n", t.text(t.Green), t.text(t.BrightGreen))
	s += fmt.Sprintf("%v%v\n", t.text(t.Yellow), t.text(t.BrightYellow))
	s += fmt.Sprintf("%v%v\n", t.text(t.Blue), t.text(t.BrightBlue))
	s += fmt.Sprintf("%v%v\n", t.text(t.Magenta), t.text(t.BrightMagenta))
	s += fmt.Sprintf("%v%v\n", t.text(t.Cyan), t.text(t.BrightCyan))
	s += fmt.Sprintf("%v%v\n", t.text(t.White), t.text(t.BrightWhite))
	s += fmt.Sprintf("%v\n", t.render("                                                    "))
	s += fmt.Sprintf("%v\n", t.commandline())
	s += fmt.Sprintf("%v\n", t.render("                                                    "))

	//"$ sudo apt install linux |"
	return
}

func (t Theme) render(s string) string {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.Foreground)).
		Background(lipgloss.Color(t.Background))

	return style.Render(s)
}

func (t Theme) block(color string) string {
	var style = lipgloss.NewStyle().
		Background(lipgloss.Color(color))

	return style.Render("      ")
}

func (t Theme) text(color string) string {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Background(lipgloss.Color(t.Background))

	return style.Render(fmt.Sprintf("   AaBbMmYyZz - %v   ", color))
}

func (t Theme) commandline() string {
	return fmt.Sprintf("%v%v%v%v%v%v",
		t.render("  $ "),
		lipgloss.NewStyle().Background(lipgloss.Color(t.Background)).Foreground(lipgloss.Color(t.Red)).Render("sudo "),
		lipgloss.NewStyle().Background(lipgloss.Color(t.Background)).Foreground(lipgloss.Color(t.Green)).Render("apt "),
		t.render("install linux "),
		lipgloss.NewStyle().Blink(true).Background(lipgloss.Color(t.Background)).Foreground(lipgloss.Color(t.Foreground)).Render("|"),
		t.render("                        "),
	)
}
