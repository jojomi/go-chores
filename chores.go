package chores

import (
	"fmt"

	"github.com/fatih/color"
)

// terminal colors
var fmtErr = color.New(color.FgRed).Add(color.Bold)
var fmtSucc = color.New(color.FgGreen).Add(color.Bold)
var fmtInfo = color.New(color.FgYellow)

type ChoreState int

const (
	STATE_SUCCESS ChoreState = iota
	STATE_WARNING
	STATE_ERROR
)

const (
	OUTPUT_OUT int = 1 << iota
	OUTPUT_ERR
)

type Chore interface {
	Name() string
	SetName(name string)
	Description() string
	SetDescription(description string)

	Execute() ChoreState

	Possible() bool
	Necessary() bool
}

// Execute executes a list of chores supplied outputting status on stdout.
func Execute(choreList []Chore) error {
	fmtInfo.Println("Doing chores...")
	for _, chore := range choreList {
		fmtInfo.Printf("Looking at '%s'...\n", chore.Name())
		if !chore.Necessary() {
			fmtInfo.Printf("Not necessary.")
			continue
		}
		if !chore.Possible() {
			fmtErr.Printf("Necessary, but not possible!")
			continue
		}
		state := chore.Execute()
		if state == STATE_SUCCESS {
			fmtSucc.Printf("✓ (success)\n")
		} else if state == STATE_ERROR {
			fmtErr.Printf("⚠ (error)\n")
			fmt.Println("[Output]")
			//fmt.Println(ctx.LastOutput().String())
			//fmt.Println(ctx.LastError().String())
		}
	}
	fmtInfo.Println("Chores done.")
	return nil
}
