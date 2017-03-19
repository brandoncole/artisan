package model

import (
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	location string   `json:",omit" yaml:",omit"`
	Version  string   `json:"version,omitempty" yaml:"version,omitempty"`
	Inputs   *Inputs  `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs  *Outputs `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

type Inputs struct {
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}

type Outputs struct {
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}

func (m *Manifest) Resolve(relative_path string) string {
	return path.Join(m.location, relative_path)
}

func (m *Manifest) Validate() error {
	// TODO: Validate Manifest
	return nil
}

func NewManifest(path string) (*Manifest, error) {

	data, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}

	result := &Manifest{}
	err = yaml.Unmarshal(data, result)
	if nil != err {
        return nil, err
    }

	err = result.Validate()
	if nil != err {
		return nil, err
	}

    return result, nil

}
