package request

import "fmt"

type Configuration struct {
	ApiKey       string
	Organization string
}

func NewConfiguration(apiKey string) *Configuration {
	return &Configuration{
		ApiKey: apiKey,
	}
}

func NewConfigurationWithOrg(apiKey string, organization string) *Configuration {
	return &Configuration{
		ApiKey:       apiKey,
		Organization: organization,
	}
}

func (c *Configuration) Headers() map[string]string {
	var header = map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.ApiKey),
	}
	if len(c.Organization) != 0 {
		header["OpenAI-Organization"] = c.Organization
	}
	return header
}
