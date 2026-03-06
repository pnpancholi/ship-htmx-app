package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	cyan  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BCD4"))
	bold  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BCD4"))
	dim   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4A9EAD"))
	white = lipgloss.NewStyle().Foreground(lipgloss.Color("#E0E0E0"))
)

const banner = `
▞▀▖▌ ▌▜▘▛▀▖   ▌ ▌▀▛▘▙▗▌▌ ▌  ▞▀▖▛▀▖▛▀▖
▚▄ ▙▄▌▐ ▙▄▘▄▄▖▙▄▌ ▌ ▌▘▌▝▞▄▄▖▙▄▌▙▄▘▙▄▘
▖ ▌▌ ▌▐ ▌     ▌ ▌ ▌ ▌ ▌▞▝▖  ▌ ▌▌  ▌  
▝▀ ▘ ▘▀▘▘     ▘ ▘ ▘ ▘ ▘▘ ▘  ▘ ▘▘  ▘                                                                                                           
`

func printBanner() {
	fmt.Println(cyan.Render(banner))
	fmt.Println(bold.Render("ship-htmx-app") + white.Render(" v0.1.0"))
	fmt.Println()
	fmt.Println(dim.Render("Build robust HTMX websites and applications rapidly"))
	fmt.Println()
	fmt.Println(dim.Render(" Docs ") + white.Render("https://shiphtmx.com"))
	fmt.Println()
	fmt.Println(dim.Render(" Available Commands"))
	// ToDo: Add table of available commands//
}
