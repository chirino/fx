package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antonmedv/fx/pkg/model"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

func usage(keyMap model.KeyMap) string {
	title := lipgloss.NewStyle().Bold(true)
	pad := lipgloss.NewStyle().PaddingLeft(4)
	return fmt.Sprintf(`
  %v
    Terminal JSON viewer

  %v
    fx data.json
    fx data.json .field
    curl ... | fx

  %v
    -h, --help            print help
    -v, --version         print version
    --themes              print themes
    -r, --raw             treat input as a raw string
    -s, --slurp           read all inputs into an array

  %v
%v

  %v
    [https://fx.wtf]
`,
		title.Render("fx "+version),
		title.Render("Usage"),
		title.Render("Flags"),
		title.Render("Key Bindings"),
		strings.Join(keyMapInfo(keyMap, pad), "\n"),
		title.Render("More info"),
	)
}

func keyMapInfo(keyMap model.KeyMap, style lipgloss.Style) []string {
	v := reflect.ValueOf(keyMap)
	fields := reflect.VisibleFields(v.Type())

	keys := make([]string, 0)
	for i := range fields {
		k := v.Field(i).Interface().(key.Binding)
		str := k.Help().Key
		if len(str) == 0 {
			str = strings.Join(k.Keys(), ", ")
		}
		keys = append(keys, fmt.Sprintf("%v    ", str))
	}

	desc := make([]string, 0)
	for i := range fields {
		k := v.Field(i).Interface().(key.Binding)
		desc = append(desc, fmt.Sprintf("%v", k.Help().Desc))
	}

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		strings.Join(keys, "\n"),
		strings.Join(desc, "\n"),
	)

	return strings.Split(style.Render(content), "\n")
}
