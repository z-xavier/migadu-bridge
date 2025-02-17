package config

import (
	"strings"

	"github.com/spf13/viper"

	"migadu-bridge/internal/pkg/log"
)

type Config struct {
	LogConf    *LogConf    `json:"log" yaml:"log"`
	MigaduConf *MigaduConf `json:"migadu" yaml:"migadu"`
	ServerConf *ServerConf `json:"server" yaml:"server"`
}

type LogConf struct {
	DisableCaller     bool     `json:"disable_caller" yaml:"disable_caller"`
	DisableStacktrace bool     `json:"disable_stacktrace" yaml:"disable_stacktrace"`
	Level             string   `json:"level" yaml:"level"`
	Format            string   `json:"format" yaml:"format"`
	OutputPaths       []string `json:"output_paths" yaml:"output_paths"`
}

type MigaduConf struct {
	Email  string `json:"token" yaml:"token"`
	APIKey string `json:"api_key" yaml:"api_key"`
	Domain string `json:"domain" yaml:"domain"` // TODO 可配置
	//Timeout int    `json:"timeout" yaml:"timeout"`
}

type ServerConf struct {
	RunMode         string `json:"runmode" yaml:"runmode"`
	WebAddr         string `json:"web_addr" yaml:"web_addr"`
	InteriorWebAddr string `json:"interior_web_addr" yaml:"interior_web_addr"`
}

var config *Config

func GetConfig() *Config {
	return config
}

// InitConfig 设置需要读取的配置文件名、环境变量，并读取配置文件内容到 viper 中.
func InitConfig(cfgFile string) error {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath("/etc/migadu-provider/")
		viper.AddConfigPath("/config")
		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("./conf/")

		// 设置配置文件格式为 YAML (YAML 格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName("conf.yaml")
	}
	// 从命令行选项指定的配置文件中读取
	viper.SetConfigFile(cfgFile)

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 MIGADU_BRIDGE，如果是 migadu_bridge，将自动转变为大写。
	viper.SetEnvPrefix("MIGADU_BRIDGE")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Errorw("Failed to unmarshal viper configuration", "err", err)
		return err
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
	return nil
}

// LogOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func LogOptions() *log.Options {
	return &log.Options{
		DisableCaller:     config.LogConf.DisableCaller,
		DisableStacktrace: config.LogConf.DisableStacktrace,
		Level:             config.LogConf.Level,
		Format:            config.LogConf.Format,
		OutputPaths:       config.LogConf.OutputPaths,
	}
}
