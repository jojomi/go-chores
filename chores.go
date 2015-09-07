package chores

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jojomi/go-script"
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

type Chore struct {
	Name        string
	Description string
	Command     []string
	WorkingDir  string
	Frequency   string
	Output      int
	SuccessFunc func(*script.Context) bool
}

var successOnReturnZero = func(c *script.Context) bool {
	return c.LastExitCode() == 0
}

// New returns a new standard chore.
func New() (chore *Chore) {
	// initialize Chore
	chore = &Chore{
		SuccessFunc: successOnReturnZero,
	}

	return
}

// Execute executes a single chore.
func (c *Chore) Execute() (choreState ChoreState, ctx *script.Context) {
	// TODO check frequency requirement
	sc := script.NewContext()
	sc.SetWorkingDir(c.WorkingDir)
	// get execution function
	stdoutSilent := c.Output&OUTPUT_OUT == 0
	stderrSilent := c.Output&OUTPUT_ERR == 0
	err := sc.Execute(stdoutSilent, stderrSilent, c.Command[0], c.Command[1:]...)
	if err != nil {
		return STATE_ERROR, sc
	}
	if c.SuccessFunc(sc) {
		return STATE_SUCCESS, sc
	} else {
		return STATE_ERROR, sc
	}
}

// Execute executes a list of chores supplied outputting status on stdout.
func Execute(choreList []*Chore) error {
	fmtInfo.Println("Doing chores...")
	for _, chore := range choreList {
		fmtInfo.Printf("Executing '%s'...\n", chore.Name)
		state, ctx := chore.Execute()
		if state == STATE_SUCCESS {
			fmtSucc.Printf("✓ (success)\n")
		} else if state == STATE_ERROR {
			fmtErr.Printf("⚠ (error)\n")
			fmt.Println(ctx.LastOutput().String())
			fmt.Println(ctx.LastError().String())
		}
	}
	fmtInfo.Println("Chores done.")
	return nil
}
