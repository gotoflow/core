package config

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Driver string `yaml:"driver"`
	} `yaml:"database"`

	Base struct {
		DagDir string `yaml:"dag_dir"`
	} `yaml:"base"`
}