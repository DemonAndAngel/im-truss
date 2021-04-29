

package templates

const Config = `
package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-common/tools/db"
	es "go-common/tools/elasticsearch"
	"go-common/tools/logger"
	"go-common/tools/redis"
	"go-common/tools/tracing"
	"os"
	"time"
)

// 加载配置文件
func init() {
	logger.Info("ConfigInit", "SUCCESS", "加载配置文件", nil)
    viper.SetConfigName("app")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("ConfigInit", "LOAD_FAIL", "加载配置文件失败;异常信息:" + err.Error(), nil)
	}
	// 使用配置
	if err := ConnResources(); err != nil {
		logger.Error("ConfigInit", "USE_FAIL", "应用配置文件失败;异常信息:" + err.Error(), nil)
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}
	// 配置监听
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("ConfigInit", "SUCCESS", "配置文件变更", nil)
		if err := ConnResources(); err != nil {
			logger.Error("ConfigInit", "RELOAD_USE_FAIL", "重新加载配置文件失败;异常信息:" + err.Error(), nil)
		}
	})
	logger.Info("ConfigInit", "SUCCESS", "成功加载配置文件,并监听变更中...", nil)
}

func ConnResources() (error) {
	// 重载配置
	str := "启动失败!"
	err := db.InitWithEnv()
	if err != nil {
		return errors.New(str + "数据库连接异常:"+err.Error())
	}
	err = redis.InitWithEnv()
	if err != nil {
		return errors.New(str + "缓存连接异常:"+err.Error())
	}
	err = es.InitWithEnv()
	if err != nil {
		return errors.New(str + "es连接异常:"+err.Error())
	}
	err = tracing.InitWithEnv()
	if err != nil {
		return errors.New(str + "链路跟踪异常:" + err.Error())
	}
	return nil
}`

const ConfigAppJsonExample = `
{
    "appEnv": "local",
    "appDebug": true,
    "appName": "服务名",
    "appEnName": "serviceName",
    "driver": {
        "mysql": "default",
        "redis": "default",
        "jaeger": "default",
        "es": "default"
    },
    "mysql": {
        "default": {
            "database": "database",
            "prefix": "prefix_",
            "connMaxLifetime": 3600,
            "host": "172.17.0.1",
            "port": 33057,
            "username": "root",
            "password": "root"
        }
    },
    "redis": {
        "default": {
            "host": "172.17.0.1",
            "port": 6379,
            "password": "962464543",
            "database": 0
        }
    },
    "jaeger": {
        "default": {
            "host": "172.17.0.1",
            "port": 6831
        }
    },
    "es": {
        "default": {
            "addr": "http://172.17.0.1:9200",
            "username": "elastic",
            "password": "changeme"
        }
    }
}
`
const ConfigAppJson = `
{
    "appEnv": "local",
    "appDebug": true,
    "appName": "服务名",
    "appEnName": "serviceName",
    "driver": {
        "mysql": "default",
        "redis": "default",
        "jaeger": "default",
        "es": "default"
    },
    "mysql": {
        "default": {
            "database": "database",
            "prefix": "prefix_",
            "connMaxLifetime": 3600,
            "host": "172.17.0.1",
            "port": 33057,
            "username": "root",
            "password": "root"
        }
    },
    "redis": {
        "default": {
            "host": "172.17.0.1",
            "port": 6379,
            "password": "962464543",
            "database": 0
        }
    },
    "jaeger": {
        "default": {
            "host": "172.17.0.1",
            "port": 6831
        }
    },
    "es": {
        "default": {
            "addr": "http://172.17.0.1:9200",
            "username": "elastic",
            "password": "changeme"
        }
    }
}
`
