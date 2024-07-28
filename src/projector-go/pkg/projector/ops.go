package projector

import (
	"github.com/hellflame/argparse"
)

type ProjectorOpts struct {
	Pwd       string
	Config    string
	Arguments []string
}

func GetOpts() (*ProjectorOpts, error) {
	parser := argparse.NewParser("projector", "gets all the values", &argparse.ParserConfig{DisableDefaultShowHelp: true})
	args := parser.Strings("f", "foo", &argparse.Option{
		Positional: true,
		Default:    "",
		Required:   false,
	})

	config := parser.String("c", "config", &argparse.Option{Required: false, Default: ""})
	pwd := parser.String("p", "pwd", &argparse.Option{Required: false, Default: ""})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &ProjectorOpts{
		Pwd:       *pwd,
		Config:    *config,
		Arguments: *args,
	}, nil
}
