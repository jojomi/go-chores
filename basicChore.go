package chores

type BasicChore struct {
	name        string
	description string

	exec func() ChoreState 
}

func (c *BasicChore) Name() string {
	return c.name
}

func (c *BasicChore) SetName(name string) {
	c.name = name
}

func (c *BasicChore) Description() string {
	return c.description
}

func (c *BasicChore) SetDescription(description string) {
	c.description = description
}

func (c *BasicChore) Necessary() bool {
	// TODO implement frequency requirements
	return true
}

func (c *BasicChore) Possible() bool {
	// TODO provide hook for own function
	return true
}

func (c *BasicChore) Execute() ChoreState {
	return c.exec()
}

func (c *BasicChore) SetExec(exec func() ChoreState) {
	c.exec = exec
}
