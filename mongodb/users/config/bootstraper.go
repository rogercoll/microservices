package config

func StartUp() {

	// Initialize AppConfig variable
	initConfig()
	// Start a MongoDB session
	createDbSession()
}