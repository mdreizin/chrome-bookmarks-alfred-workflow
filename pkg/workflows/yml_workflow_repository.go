package workflows

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YmlWorkflowRepository struct {
	Filename string
}

func (r *YmlWorkflowRepository) GetWorkflowMetadata() (WorkflowMetadata, error) {
	b, err := ioutil.ReadFile(r.Filename)

	if err != nil {
		return nil, err
	}

	m := WorkflowMetadata{}
	err = yaml.Unmarshal(b, &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}
