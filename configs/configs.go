package configs

import "os"

var (
	SEACHERPORT = getPort()
)

func getPort() string {
	port := os.Getenv("SEACHERPORT")
	if port == "" {
		Logger.Warn("There is no SEACHERPORT!")
		port = "8000"
	}

	return port
}
