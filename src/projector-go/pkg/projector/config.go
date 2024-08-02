package projector

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota // enum
	Add
	Remove
)

type Config struct {
	Args      []string
	Operation Operation
	Config    string
	Pwd       string
}

func getPwd(opts *ProjectorOpts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}
func getConfigPath(opts *ProjectorOpts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(config, "projector", "projector.json"), nil
}
func getOperation(opts *ProjectorOpts) Operation {
	if len(opts.Arguments) == 0 {
		return Print
	}
	if opts.Arguments[0] == "add" {
		return Add
	}
	if opts.Arguments[0] == "remove" {
		return Remove
	}
	return Print
}

func getArgs(opts *ProjectorOpts) ([]string, error) {
	if len(opts.Arguments) == 0 {
		return []string{}, nil
	}
	opertation := getOperation(opts)
	if opertation == Add {
		if len(opts.Arguments) != 3 {
			return nil, fmt.Errorf("add requires 2 args , but recieved %v,",len(opts.Arguments) - 1)
		}
		return opts.Arguments[1:], nil
	}
	if opertation == Remove {
		if len(opts.Arguments) != 2 {
			return nil, fmt.Errorf("remove requires 1 args , but recieved %v",len(opts.Arguments) - 1)
		}
		return opts.Arguments[1:], nil
	}
	if len (opts.Arguments) > 1{
		return nil, fmt.Errorf("print requires 0 or 1 arguments, but recieved %v",len(opts.Arguments))
	}
	return opts.Arguments, nil
}

// might change
func NewConfig(opts *ProjectorOpts) (*Config, error) {
	pwd, err := getPwd(opts)
	if err != nil {
		return nil, err
	}
	config, err := getConfigPath(opts)
	if err != nil {
		return nil, err
	}
	args, err := getArgs(opts)
	if err != nil {
		return nil, err
	}
	return &Config{
		Pwd: pwd,
		Config: config,
		Operation: getOperation(opts),
		Args: args,
	}, nil
}
