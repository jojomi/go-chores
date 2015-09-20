package chores

import (
	"github.com/jojomi/go-script"
)

type CommandChore struct {
	BasicChore

	Command     []string
	WorkingDir  string
	Frequency   string
	Output      int
	SuccessFunc func(*script.Context) bool
	ctx         *script.Context
}

var successOnReturnZero = func(c *script.Context) bool {
	return c.LastExitCode() == 0
}

// New returns a new command based chore.
func NewCommandChore() (chore *CommandChore) {
	// initialize Chore
	chore = &CommandChore{
		SuccessFunc: successOnReturnZero,
	}

	return
}

// Execute executes a single chore.
func (c *CommandChore) Execute() (choreState ChoreState) {
	// TODO check frequency requirement
	sc := script.NewContext()
	sc.SetWorkingDir(c.WorkingDir)
	// get execution function
	stdoutSilent := c.Output&OUTPUT_OUT == 0
	stderrSilent := c.Output&OUTPUT_ERR == 0
	err := sc.Execute(stdoutSilent, stderrSilent, c.Command[0], c.Command[1:]...)
	c.ctx = sc
	if err != nil {
		return STATE_ERROR
	}
	if c.SuccessFunc(sc) {
		return STATE_SUCCESS
	} else {
		return STATE_ERROR
	}
}

func (c *CommandChore) Context() *script.Context {
	return c.ctx
}
