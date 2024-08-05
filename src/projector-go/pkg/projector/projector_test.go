package projector_test

import "projects/projector-go/src/projector-go/pkg/projector"

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{},
			"/": map[string]string {
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