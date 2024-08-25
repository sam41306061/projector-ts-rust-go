package projector_test

import (
	"projects/projector-go/src/projector-go/pkg/projector"
	"testing"
)

func getData() *projector.Data {
	return &projector.Data {
		Projector: map[string]map[string]string {
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector (pwd string, data *projector.Data) *projector.Projector{
	return projector.CreateProjector(
		&projector.Config{
			Args: []string{},
			Operation: projector.Print,
			Pwd: pwd,
			Config: "Hell, Frontend Masters",
		},
		data,
	);
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)	
	if !ok {
		t.Errorf("exected to get value \"%v\"", value)
	}
	if value != "bar3"{
		t.Errorf("expected to find %v  but got %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t,proj,"foo","bar3");
}
