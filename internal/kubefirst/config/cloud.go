package config

type Cloud struct {
	AWS    AWS    `yaml:"aws"`
	Azure  Azure  `yaml:"azure"`
	Google Google `yaml:"google"`
}

type AWS struct {
	Region    string `yaml:"region"`
	AccountID string `yaml:"accountID"`
}

type Azure struct {
	Region           string `yaml:"region"`
	IdentityClientID string `yaml:"identityClientID"`
}

type Google struct {
	Region      string `yaml:"region"`
	ProjectName string `yaml:"projectName"`
}
