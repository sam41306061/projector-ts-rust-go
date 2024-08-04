package projector_test

import (
	"testing"

	"projects/projector-go/src/projector-go/pkg/projector"
)

func TestConfigPrint(t *testing.T) {
    opts := projector.ProjectorOpts {
        Config: "",
        Pwd: "",
        Arguments: []string{},
    }

    config, err := projector.NewConfig(&opts)

    if err != nil {
        t.Errorf("error returned from projector config %v", err)
    }

    if config.Operation != projector.Print {
        t.Errorf("operation expected was print but got %v", config.Operation)
    }
}

func TestConfigAdd(t *testing.T) {
    opts := projector.ProjectorOpts {
        Config: "",
        Pwd: "",
        Arguments: []string{"add", "foo", "bar"},
    }

    config, err := projector.NewConfig(&opts)

    if err != nil {
        t.Errorf("error returned from projector config %v", err)
    }

    if config.Operation != projector.Add {
        t.Errorf("operation expected was add but got %v", config.Operation)
    }

    if config.Args[0] != "foo" || config.Args[1] != "bar" {
        t.Errorf("expected arguments to equal {'foo', 'bar'} but got %+v", config.Args)
    }
}

func TestConfigRemove(t *testing.T) {
	opts := projector.ProjectorOpts {
		Config: "",
		Pwd: "",
		Arguments: []string("remove", "foo", "bar"),
	}
	config, err := projector.NewConfig(&opts)
	if err != nil {
		t.Errorf("error returned from projector confg %v", config.Operation)
	}
	if config.Operation != projector.Remove {
		t.Errorf("operation expected was remove but got %v", config.Operation)
	}
	if config.Args[0] != "foo" {
		t.Errorf("expected arugments to equal {'foo'} but got %v", config.Args)
	}
}