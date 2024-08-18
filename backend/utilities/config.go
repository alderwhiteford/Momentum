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

type AuthSettings = struct {
	JwtSecret string
}

type ApplicationSettings = struct {
	Database DatabaseSettings
	AuthSettings AuthSettings
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
func configAppSettings(env Environment) (*ApplicationSettings, error) {
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

	// Load Prod Database Settings:
	if env == EnvironmentProduction {
		prodDBSettings, err := configProdDatabase();
		if err != nil {
			return nil, err;
		}
		settings.Database = *prodDBSettings;
	}

	// Load Auth Settings:
	authSettings, err := configAuth();
	if err != nil {
		return nil, err;
	}
	settings.AuthSettings = *authSettings;

	return settings, nil;
}

func configProdDatabase() (*DatabaseSettings, error) {
	host := os.Getenv("DB_HOST");
	if host == "" {
		return nil, errors.New("database host not set");
	}
	port := os.Getenv("DB_PORT");
	if port == "" {
		return nil, errors.New("database port not set");
	}
	name := os.Getenv("DB_NAME");
	if name == "" {
		return nil, errors.New("database name not set");
	}
	username := os.Getenv("DB_USERNAME");
	if username == "" {
		return nil, errors.New("database username not set");
	}
	password := os.Getenv("DB_PASSWORD");
	if password == "" {
		return nil, errors.New("database password not set");
	}

	databaseSettings := DatabaseSettings{
		Host: host,
		Port: port,
		Username: username,
		Password: password,
	}

	return &databaseSettings, nil;
}

func configAuth() (*AuthSettings, error) {
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET");
	if jwtSecret == "" {
		return nil, errors.New("JWT Secret not set");
	}
	
	authSettings := AuthSettings{
		JwtSecret: jwtSecret,
	}

	return &authSettings, nil;
}

/** Load the settings from the configuration file **/
func Settings(configPath string) ApplicationSettings {
	// Get the environment:
	env, err := whatEnvironment();
	if err != nil {
		fmt.Printf("Error: %s", err);
		os.Exit(1);
	}
	
	// Config the non-secure settings:
	settings, err := configAppSettings(env);
	if err != nil {
		fmt.Printf("Error: %s", err);
		os.Exit(1);
	}

	return *settings;
}