package awsome

type (
	Config struct {
		Location string
		Address  string
		GRPCPort string
		HTTPPort string
	}
)

func LoadConfiguration() *Config {
	return &Config{
		Address:  "localhost",
		GRPCPort: "8090",
		HTTPPort: "8080",
	}
}
