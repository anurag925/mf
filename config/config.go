package config

type secret struct {
	DBUrl string `env:"DB_URL" envDefault:"mysql://root:@tcp(localhost:3306)/mf_development?parseTime=true"`
}

type setting struct {
	TestConfig string `yaml:"test_config"`
}
