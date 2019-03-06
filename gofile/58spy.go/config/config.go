package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

//Config is to store config file name
type Config struct {
	Name string
}

// Init is to init server with config file
func Init(configPath string) (err error) {
	c := Config{Name: configPath}

	//初始化配置文件
	if err = c.initConfig(); err != nil {
		log.Error("Config field", err)
		return

	}

	//初始化日志配置
	c.initLog()

	//监控配置文件变化，热加载程序
	c.watchConfig()

	return
}
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	if err := log.InitWithConfig(&passLagerCfg); err != nil {
		log.Warn("faild config log 1111111111111111111111" + err.Error())
	}
	log.Info("success config log 222222222222222222" + viper.GetString("log.logger_file"))
}

func (c *Config) initConfig() error {
	ns := strings.Split(c.Name, ".")

	if c.Name != "" {
		viper.SetConfigFile(c.Name)
		viper.SetConfigType(ns[1])
	} else {
		viper.SetConfigFile("conf/config.yaml")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("58spy")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("Config file changed %s", in.Name)
	})
}
