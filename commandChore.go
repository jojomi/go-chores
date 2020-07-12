package chores

import (
	"github.com/jojomi/go-script/v2"
)

// CommandChore is a chore based on a cli command to be executed.
type CommandChore struct {
	BasicChore

	Command     []string
	WorkingDir  string
	Frequency   string
	Output      int
	SuccessFunc func(*script.ProcessResult) bool
	ctx         *script.Context
}

var successOnReturnZero = func(result *script.ProcessResult) bool {
	return result.Successful()
}

// NewCommandChore returns a new command based chore.
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
	lc := script.NewLocalCommand()
	lc.AddAll(c.Command...)
	// get execution function
	processResult, err := sc.Execute(script.CommandConfig{
		OutputStdout: c.Output&OUTPUT_OUT == 1,
		OutputStderr: c.Output&OUTPUT_ERR == 1,
		ConnectStdin: true,
	}, lc)
	c.ctx = sc
	if err != nil {
		return StateError
	}
	if c.SuccessFunc(processResult) {
		return StateSuccess
	}
	return StateError
}

// Context returns the CommandChore's script.Context
func (c *CommandChore) Context() *script.Context {
	return c.ctx
}
