package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ActiveProfile     string
	MigrationFilePath string

	SecretKey       string
	Issuer          string
	TokenExpiryTime int `mapstructure:"TOKEN_EXPIRY_TIME"`

	ProjectRootPath string `mapstructure:"PROJECT_ROOT_PATH"`

	GinMode string `mapstructure:"GIN_MODE"`

	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	AppReadTimeOut    int    `mapstructure:"APP_READ_TIMEOUT"`
	AppWriteTimeOut   int    `mapstructure:"APP_WRITE_TIMEOUT"`
	AppIdleTimeOut    int    `mapstructure:"APP_IDLE_TIMEOUT"`

	DevelopmentMode bool `mapstructure:"DEVELOPMENT_MODE"`
}

func InitConfig(context.Context) (appConfig *AppConfig, err error) {
	activeProfile := os.Getenv("ACTIVE_PROFILE")

	if activeProfile != "dev" && activeProfile != "test" {
		isEnvAvailable, err := allSecretEnvVarsSet()
		if !isEnvAvailable || err != nil {
			return appConfig, fmt.Errorf("[InitConfig, allSecretEnvVarsSet], env vars not set, error:%w", err)
		}
	}

	viper.AddConfigPath(fmt.Sprintf("%s/%s", RootPath(), "env"))

	if activeProfile != "" {
		viper.SetConfigName(activeProfile)
	} else {
		viper.SetConfigName("dev")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return appConfig, fmt.Errorf("[InitConfig, viper.ReadInConfig], cannot load appConfig, error:%w", err)
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return appConfig, fmt.Errorf("[InitConfig, viper.Unmarshal], Viper unmarshal error, error:%w", err)
	}

	appConfig.ActiveProfile = activeProfile
	appConfig.ProjectRootPath = RootPath()
	appConfig.MigrationFilePath = "database/migration"

	return appConfig, nil
}

func RootPath() (path string) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return path
	}
	path = filepath.Dir(currentFile)
	path = filepath.Dir(path)
	return path
}

func allSecretEnvVarsSet() (isEnvAvailable bool, err error) {
	requiredEnvVars := []string{}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return false, fmt.Errorf("[allSecretEnvVarsSet], required env var not set, var:%s", envVar)
		}
	}

	return true, nil
}
