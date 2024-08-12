package utilities

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	config "github.com/spf13/viper"
)

type DatabaseSettings = struct {
	Host string
	Port string
	Username string
	Password string
	Name string
	RequireSSL bool
}

type ApplicationSettings = struct {
	Database DatabaseSettings
}

type Environment string

const (
	EnvironmentLocal      Environment = "local"
	EnvironmentProduction Environment = "production"
)

/** Identify what environment the application is running in **/
func whatEnvironment() (Environment, error) {
	// Load the enivronment variables:
	godotenv.Load(".env");
	environment := os.Getenv("APP_ENVIRONMENT");

	if environment == "" {
		return "", errors.New("Environment variable not set");
	}

	return Environment(environment), nil;
}

/** Configure the settings based on the environment **/
func configByEnvironment(env Environment) (*ApplicationSettings, error) {
	// Initialize config settings:
	config.SetConfigName(string(env));
	config.AddConfigPath("config");
	config.SetConfigType("yaml");
	config.ReadInConfig();

	// Marshal settings into application settings struct:
	var settings *ApplicationSettings;
	if err := config.Unmarshal(&settings); err != nil {
		return nil, err
	}

	// Load from the environment variables if in production:
	if env == EnvironmentProduction {
		settings.Database.Host = os.Getenv("DB_HOST");
		settings.Database.Port = os.Getenv("DB_PORT");
		settings.Database.Username = os.Getenv("DB_USERNAME");
		settings.Database.Password = os.Getenv("DB_PASSWORD");
	}

	return settings, nil;
}

/** Load the settings from the configuration file **/
func Settings(configPath string) ApplicationSettings {
	// Get the environment:
	env, err := whatEnvironment();
	if err != nil {
		fmt.Printf("Error: %s", err);
		os.Exit(1);
	}
	
	settings, err := configByEnvironment(env);
	if err != nil {
		fmt.Printf("Error: %s", err);
		os.Exit(1);
	}

	return *settings;
}