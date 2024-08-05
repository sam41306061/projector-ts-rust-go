package projector

import (
	"encoding/json"
	"os"
	"path"
)

// struct
type Data struct {
	Projector map[string]map[string]string `json:"projector"`	
}

type Projector struct {
	config *Config
	data *Data
}

func (p *Projector) GetValue(key string) (string, bool) {
	cur := p.config.Pwd
	prev := ""

	out := ""
	found := false
	// while loop example
	for ; cur != prev; {
		if dir, ok := p.data.Projector[cur]; ok {
			if value, ok := dir[key]; ok {
				out = value 
				found = true
				break
			}
		} 

		prev = cur
		cur = path.Dir(cur)
	}

	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	out := map[string]string{}
	paths := []string{}

	cur := p.config.Pwd
	prev := ""
	
	for ; cur != prev; {
		paths = append(paths, cur)
		prev = cur
		cur = path.Dir(cur)
	}
	// map mergeing example 
	for i := len(paths) - 1; i >= 0; i -- {
		if dir, ok := p.data.Projector[paths[i]]; ok {
			for k, v := range dir {
				out[k] = v
			}
		}
	}

	return out
}

func (p *Projector) SetValue( key, value string) {
	pwd := p.config.Pwd	
	if _, ok := p.data.Projector[pwd]; !ok {
		p.data.Projector[pwd] =  map[string]string{}
	}
	p.data.Projector[pwd][key] = value
}
func (p *Projector) RemoveValue( key, value string) {
	pwd:= p.config.Pwd	
	if dir, ok := p.data.Projector[pwd]; ok {
		delete(dir,key)
	}
}

func defaultProjector(config *Config) *Projector {
	return &Projector{
		config: config,
		data: &Data{
			Projector: map[string]map[string]string{},
		},
	}
}


func NewProjector(config *Config) *Projector {
	if _, err :=  os.Stat(config.Config); err == nil {
		contents, err := os.ReadFile(config.Config)
		if err != nil {
			return defaultProjector(config)
		}
		var data Data
		err = json.Unmarshal(contents, &data);

		if err != nil {
			return defaultProjector(config)
		}
		return &Projector{
			data: &data,
			config: config,
		}
	}	
	return defaultProjector(config)
}