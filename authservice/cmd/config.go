package cmd

type ServerConfig struct {
	Port string `arg:"PORT"`
	Name string `arg:"SERVICE_NAME"`
}

func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		Name: "authservice",
		Port: "8081",
	}

}
