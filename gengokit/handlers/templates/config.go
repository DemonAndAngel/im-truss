

package templates

const Config = `
package config

import (
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"log"
	"go-common/tools/db"
	es "go-common/tools/elasticsearch"
	"go-common/tools/tracing"
	"go-common/tools/redis"
	"errors"
)

// 加载配置文件
func init() {
	log.Println("加载配置文件...")
    viper.SetConfigName("app")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("加载配置文件失败;信息:" + err.Error())
	}
	// 使用配置
	if err := ConnResources(); err != nil {
		log.Fatal(err.Error())
	}
	// 配置监听
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件变更", e.Name)
		if err := ConnResources(); err != nil {
			log.Println(err.Error())
		}
	})
	log.Println("成功加载配置文件,并监听变更中...")
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
