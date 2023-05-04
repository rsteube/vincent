package scheme

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Scheme struct {
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

func (sc Scheme) Colors() []string {
	return append(sc.ColorsNormal(), sc.ColorsBright()...)
}

func (sc Scheme) ColorsNormal() []string {
	return []string{
		sc.Black,
		sc.Red,
		sc.Green,
		sc.Yellow,
		sc.Blue,
		sc.Magenta,
		sc.Cyan,
		sc.White,
	}

}

func (sc Scheme) ColorsBright() []string {
	return []string{
		sc.BrightBlack,
		sc.BrightRed,
		sc.BrightGreen,
		sc.BrightYellow,
		sc.BrightBlue,
		sc.BrightMagenta,
		sc.BrightCyan,
		sc.BrightWhite,
	}
}

func (sc Scheme) IsDark() bool {
	c := termenv.ConvertToRGB(termenv.RGBColor(sc.Background))
	_, _, l := c.Hsl()
	return l < 0.5
}

func (sc Scheme) Format(format string) (string, error) {
	if f := formats[format]; f != nil {
		return f(sc)
	}
	return "", fmt.Errorf("unknown format: %v - expected one of [%v]", format, strings.Join(Formats(), ", "))
}

func (sc Scheme) Render() (s string) {
	s += fmt.Sprintln(sc.render(strings.Repeat(" ", 52)))
	s += fmt.Sprintln(sc.render("  ") + sc.block(sc.ColorsNormal()...) + sc.render("  "))
	s += fmt.Sprintln(sc.render("  ") + sc.block(sc.ColorsBright()...) + sc.render("  "))
	s += fmt.Sprintln(sc.render(strings.Repeat(" ", 52)))
	s += fmt.Sprintln(sc.text(sc.Black, sc.BrightBlack))
	s += fmt.Sprintln(sc.text(sc.Red, sc.BrightRed))
	s += fmt.Sprintln(sc.text(sc.Green, sc.BrightGreen))
	s += fmt.Sprintln(sc.text(sc.Yellow, sc.BrightYellow))
	s += fmt.Sprintln(sc.text(sc.Blue, sc.BrightBlue))
	s += fmt.Sprintln(sc.text(sc.Magenta, sc.BrightMagenta))
	s += fmt.Sprintln(sc.text(sc.Cyan, sc.BrightCyan))
	s += fmt.Sprintln(sc.text(sc.White, sc.BrightWhite))
	s += fmt.Sprintln(sc.render(strings.Repeat(" ", 52)))
	s += fmt.Sprintln(sc.commandline())
	s += fmt.Sprintln(sc.render(strings.Repeat(" ", 52)))
	fmt.Sprintln()
	return
}

func (sc Scheme) render(s string) string {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(sc.Foreground)).
		Background(lipgloss.Color(sc.Background))

	return style.Render(s)
}

func (sc Scheme) block(color ...string) (s string) {
	for _, c := range color {
		s += lipgloss.NewStyle().
			Background(lipgloss.Color(c)).
			Render("      ")
	}
	return
}

func (sc Scheme) text(color ...string) (s string) {
	for _, c := range color {
		s += lipgloss.NewStyle().
			Background(lipgloss.Color(sc.Background)).
			Foreground(lipgloss.Color(c)).
			Render(fmt.Sprintf("   AaBbMmYyZz - %v   ", c))
	}
	return
}

func (sc Scheme) commandline() (s string) {
	style := lipgloss.NewStyle().Background(lipgloss.Color(sc.Background))

	s += sc.render("  $ ")
	s += style.Foreground(lipgloss.Color(sc.Red)).Render("sudo ")
	s += style.Foreground(lipgloss.Color(sc.Green)).Render("apt ")
	s += sc.render("install linux ")
	s += style.Blink(true).Foreground(lipgloss.Color(sc.Foreground)).Render("|")
	s += sc.render("                        ")
	return
}

var formats = map[string]func(sc Scheme) (string, error){
	"render": func(sc Scheme) (string, error) {
		return fmt.Sprintf("%v\n\n%v", sc.Name, sc.Render()), nil
	},
}

func Formats() []string {
	result := make([]string, 0)
	for key := range formats {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}
