package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// ReadConfigs read configds from config.env file
func ReadConfigs() {
	viper.SetConfigFile("config.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
}

// ServicePort returns http.Server's port as string
func ServicePort() string {
	return ":" + viper.GetString("SERVICE_PORT")
}

// JobTimerDuration returns duration for worker job
func JobTimerDuration() time.Duration {

	return time.Duration(viper.GetInt("JOB_TIMER_SEC"))
}

// DBConnectString returns database connect string.
// It returns sslmode=disable also
func DBConnectString() string {
	return "host=" + viper.GetString("DB_HOST") +
		" port=" + viper.GetString("DB_PORT") +
		" user=" + viper.GetString("DB_USER") +
		" dbname=" + viper.GetString("DB_NAME") +
		" password=" + viper.GetString("DB_PASS") +
		" sslmode=disable"
}
