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
        t.Errorf("expected to find %v but got nil", value)
    } else if value != v {
        t.Errorf("expected to find %v but got %v", value, v)
    }
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t,proj,"foo","bar3");
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar",data)
	test(t,proj, "foo", "bar3")
	proj.SetValue("foo", "bar4")
	test(t,proj,"foo", "bar4")
	proj.SetValue("fem", "is_super_great")
	test(t,proj,"fem", "is_super_great")

	proj = getProjector("/", data)
	test(t,proj, "fem", "is_great")
}
func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar",data)
	test(t,proj, "foo", "bar3")

	proj.RemoveValue("foo", "bar2")
	test(t,proj,"foo", "bar2")
}
