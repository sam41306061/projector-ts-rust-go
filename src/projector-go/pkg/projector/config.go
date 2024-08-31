package projector

import (
	"fmt"
	"log"
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

func getPwd(opts *Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}
func getConfigPath(opts *Opts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return path.Join(config, "projector", "projector.json"), nil
}
func getOperation(opts *Opts) Operation {
	if len(opts.Args) == 0 {
		return Print
	}
	if opts.Args[0] == "add" {
		return Add
	}
	if opts.Args[0] == "remove" {
		return Remove
	}
	return Print
}

func getArgs(opts *Opts) ([]string, error) {
    log.Printf("Args: %v", opts.Args)
    if len(opts.Args) == 0 {
        return nil, fmt.Errorf("no Args provided")
    }
    operation := getOperation(opts)
    log.Printf("Operation: %v", operation)
    if operation == Add {
        if len(opts.Args) != 3 {
            return nil, fmt.Errorf("add requires 2 args , but recieved %v,",len(opts.Args) - 1)
        }
        return opts.Args[1:], nil
    }
    if operation == Remove {
        if len(opts.Args) != 2 {
            return nil, fmt.Errorf("remove requires 1 args , but recieved %v",len(opts.Args) - 1)
        }
        return opts.Args[1:], nil
    }
    if len(opts.Args) > 1{
        return nil, fmt.Errorf("print requires 0 or 1 Args, but recieved %v",len(opts.Args))
    }
    return opts.Args, nil
}

// might change
func NewConfig(opts *Opts) (*Config, error) {
	pwd, err := getPwd(opts)
	log.Printf("Got pwd: %s, err: %v", pwd, err)
	if err != nil {
		return nil, err
	}
	config, err := getConfigPath(opts)
	log.Printf("Got config path: %s, err: %v", config, err)
	if err != nil {
		return nil, err
	}
	args, err := getArgs(opts)
	log.Printf("Got args: %v, err: %v", args, err)

	operation := getOperation(opts)
    log.Printf("Operation: %v", operation)

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
