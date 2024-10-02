package connector

import (
	"bytes"
	"errors"
	"net/http"
	"os"
)

type Connector interface {
	Register() error
}

type DebeziumConnector struct {
	URL      string
	Name     string
	JSONPath string
}

func New(toURL, connectorName, connectorPath string) Connector {
	return &DebeziumConnector{
		URL:      toURL,
		Name:     connectorName,
		JSONPath: connectorPath,
	}
}

func (c *DebeziumConnector) Register() error {
	res, err := http.Get(c.URL + "/" + c.Name)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return nil
	}

	plan, err := os.ReadFile(c.JSONPath)
	if err != nil {
		return err
	}

	res, err = http.Post(c.URL, "application/json", bytes.NewBuffer(plan))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return errors.New("failed to register connector - " + http.StatusText(res.StatusCode))
	}

	return nil
}
