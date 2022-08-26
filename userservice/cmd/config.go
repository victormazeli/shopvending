package cmd

type DatabaseConfig struct {
	Url string `arg:"env:CONNECTION_URL"`
}

type ServerConfig struct {
	Port string `arg:"env:PORT"`
	Name string `arg:"env:SERVICE_NAME"`
	DatabaseConfig
}

func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		Name: "userservice",
		Port: "8081",
		DatabaseConfig: DatabaseConfig{
			Url: "host=localhost user=vend password=rootpass dbname=userservice port=5432 sslmode=disable",
		},
	}

}
