package connector

import (
	"bytes"
	"net/http"
	"os"
)

type Connector interface {
	Register() error
}

type DebeziumConnector struct {
	URL string
	Name string
	JSONPath string
}

func New(toURL, connectorName, connectorPath string) Connector {
	return &DebeziumConnector{
		URL: toURL,
		Name: connectorName,
		JSONPath: connectorPath,
	}
}

func (c *DebeziumConnector) Register() error {
	response, err := http.Get(c.URL + "/" + c.Name)

	if err != nil {
		return err
	} else if response.StatusCode == http.StatusOK {
		return nil
	}

	plan, err := os.ReadFile(c.JSONPath)
	if err != nil {
		return err
	}

	response, err = http.Post(c.URL, "application/json", bytes.NewBuffer(plan))

	if err != nil || response.StatusCode != http.StatusOK {
		return err
	}

	return nil
}