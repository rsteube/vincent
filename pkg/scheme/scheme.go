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
	s += fmt.Sprintf("%v\n", sc.render("                                                    "))
	s += fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v\n", sc.render("  "), sc.block(sc.Black), sc.block(sc.Red), sc.block(sc.Green), sc.block(sc.Yellow), sc.block(sc.Blue), sc.block(sc.Magenta), sc.block(sc.Cyan), sc.block(sc.White), sc.render("  "))
	s += fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v\n", sc.render("  "), sc.block(sc.BrightBlack), sc.block(sc.BrightRed), sc.block(sc.BrightGreen), sc.block(sc.BrightYellow), sc.block(sc.BrightBlue), sc.block(sc.BrightMagenta), sc.block(sc.BrightCyan), sc.block(sc.BrightWhite), sc.render("  "))
	s += fmt.Sprintf("%v\n", sc.render("                                                    "))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Black), sc.text(sc.BrightBlack))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Red), sc.text(sc.BrightRed))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Green), sc.text(sc.BrightGreen))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Yellow), sc.text(sc.BrightYellow))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Blue), sc.text(sc.BrightBlue))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Magenta), sc.text(sc.BrightMagenta))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.Cyan), sc.text(sc.BrightCyan))
	s += fmt.Sprintf("%v%v\n", sc.text(sc.White), sc.text(sc.BrightWhite))
	s += fmt.Sprintf("%v\n", sc.render("                                                    "))
	s += fmt.Sprintf("%v\n", sc.commandline())
	s += fmt.Sprintf("%v\n", sc.render("                                                    "))
	return
}

func (sc Scheme) render(s string) string {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(sc.Foreground)).
		Background(lipgloss.Color(sc.Background))

	return style.Render(s)
}

func (sc Scheme) block(color string) string {
	var style = lipgloss.NewStyle().
		Background(lipgloss.Color(color))

	return style.Render("      ")
}

func (sc Scheme) text(color string) string {
	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Background(lipgloss.Color(sc.Background))

	return style.Render(fmt.Sprintf("   AaBbMmYyZz - %v   ", color))
}

func (sc Scheme) commandline() string {
	return fmt.Sprintf("%v%v%v%v%v%v",
		sc.render("  $ "),
		lipgloss.NewStyle().Background(lipgloss.Color(sc.Background)).Foreground(lipgloss.Color(sc.Red)).Render("sudo "),
		lipgloss.NewStyle().Background(lipgloss.Color(sc.Background)).Foreground(lipgloss.Color(sc.Green)).Render("apt "),
		sc.render("install linux "),
		lipgloss.NewStyle().Blink(true).Background(lipgloss.Color(sc.Background)).Foreground(lipgloss.Color(sc.Foreground)).Render("|"),
		sc.render("                        "),
	)
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
