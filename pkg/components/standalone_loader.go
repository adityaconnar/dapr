package components

import (
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	config "github.com/actionscore/actions/pkg/config/modes"
	"github.com/ghodss/yaml"
)

// StandaloneComponents loads components in a standalone mode environment
type StandaloneComponents struct {
	config config.StandaloneConfig
}

// NewStandaloneComponents returns a new standalone loader
func NewStandaloneComponents(configuration config.StandaloneConfig) *StandaloneComponents {
	return &StandaloneComponents{
		config: configuration,
	}
}

// LoadComponents loads actions components from a given directory
func (s *StandaloneComponents) LoadComponents() ([]Component, error) {
	dir := s.config.ComponentsPath
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	list := []Component{}

	for _, file := range files {
		if !file.IsDir() {
			b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, file.Name()))
			if err != nil {
				log.Warnf("error reading file: %s", err)
				continue
			}

			var component Component
			err = yaml.Unmarshal(b, &component)
			if err != nil {
				log.Warnf("error parsing file: %s", err)
				continue
			}

			list = append(list, component)
		}
	}

	return list, nil
}