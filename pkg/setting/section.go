package setting

import (
	"fmt"
	"log"
	"webconsole/pkg/zinx/ziface"
)

type ServerSettingS struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Locale  string `mapstructure:"locale"`
	Port    int    `mapstructure:"port"`

	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type LoggerSettingS struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_bacups"`
}

type CacheSettingS struct {
	CacheType string `mapstructure:"cachetype"`
	Port      string `mapstructure:"port"`
	TTL       int    `mapstructure:"ttl"`
	CacheDir  string `mapstructure:"cachedir"`
}

type DatabaseSettingS struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idel_conns"`
}

type RBACSettingS struct {
	RoleFile  string `mapstructure:"roleFile"`
	InherFile string `mapstructure:"inherFile"`
}

type OssSettingS struct {
	MqAddr string `mapstructure:"mqaddr"`
	EsAddr string `mapstructure:"esaddr"`
}

type ZinxSettingS struct {
	/*
		Zinx
	*/
	Version        string `mapstructure:"version"`          // 当前Zinx的版本号
	MaxConn        int    `mapstructure:"max_conn"`         // 当前服务器允许的最大连接数
	MaxPackageSize uint32 `mapstructure:"max_package_size"` // 当前Zinx框架数据包的最大值
	WorkerPoolSize uint32 `mapstructure:"worker_pool_size"` //当前Zinx资源池限制
	TaskQueueSize  uint32 `mapstructure:"task_queue_size"`  //当前Zinx等待队列限制

	/*
		Server
	*/
	TcpServer ziface.IServer // 当前zinx全局的Server对象
	Host      string         `mapstructure:"host"` // 当前服务器监听的IP
	Port      int            `mapstructure:"port"` // 当前服务器监听的Port
	Name      string         `mapstructure:"name"` // 当前服务器的名称
}

func (s *Setting) ReadSection(key string, v interface{}) error {
	fmt.Println(v)
	err := s.VP.UnmarshalKey(key, v)
	if err != nil {
		log.Println(key)
		return err
	}
	return nil
}
