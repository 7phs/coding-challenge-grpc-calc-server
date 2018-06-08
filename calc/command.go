package calc

import (
	"strings"
	"errors"
	"strconv"
)

// Type of supported arithmetic commands
type CommandType int

const (
	UNKNOWN CommandType = iota
	ADD
	SUB
	DIV
	MUL
	FIB
)

// Parsing a string and getting a command type numerical representation.
// A string might contain characters in lower or/and upper cases.
// UNKNOWN type will return for an unsupported command name.
func ParseCommandType(cmd string) CommandType {
	switch strings.ToLower(cmd) {
	case "add":
		return ADD
	case "sub":
		return SUB
	case "div":
		return DIV
	case "mul":
		return MUL
	case "fib":
		return FIB
	}

	return UNKNOWN
}

// Type of a command with arguments. Using for a general description arithmetical operations.
type Command struct {
	command CommandType
	values  []int64
}

// Parse a list of string parameters.
// Getting a command type and a list of integer arguments of an arithmetical operation.
// Error will return for an unknown command, empty arguments list, non numerical params and
// after additional checking of arguments using a specific command type rules.
func NewCommand(args []string) (*Command, error) {
	if len(args)==0 {
		return nil, errors.New("empty arguments list")
	}

	command := ParseCommandType(args[0])
	if command==UNKNOWN {
		return nil, errors.New("an unknown command '" + args[0] + "'")
	}

	cmd := &Command{
		command: command,
	}

	cmd.values = make([]int64, 0, len(args[1:]))
	for _, v := range args[1:] {
		n, err := strconv.ParseInt(v, 10, 64)
		if err!=nil {
			return nil, errors.New("one of the argument isn't a number: " + err.Error())
		}

		cmd.values = append(cmd.values, n)
	}

	if len(cmd.values)==0 {
		return nil, errors.New("empty arguments list")
	}

	if err:=cmd.validateValues(); err!=nil {
		return nil, errors.New("invalid arguments list: " + err.Error())
	}

	return cmd, nil
}

// Additional checking an arguments list using specific command type rules.
func (o *Command) validateValues() error {
	switch o.command {
	case FIB:
		if len(o.values)>1 {
			return errors.New("count of arguments more than one")
		}

		if o.values[0]<0 {
			return errors.New("negative value")
		}
	}

	return nil
}

// Return a command type of a command.
func (o *Command) Command() CommandType {
	return o.command
}

// Return an arguments list of a command.
func (o *Command) Values() []int64 {
	return o.values
}
