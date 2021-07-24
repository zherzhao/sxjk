package global

import (
	"webconsole/pkg/setting"
)

var (
	Conf            *setting.Setting
	ServerSetting   *setting.ServerSettingS
	LoggerSetting   *setting.LoggerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RBACSetting     *setting.RBACSettingS
)

func Init() (err error) {
	Conf, err = setting.NewSetting()
	if err != nil {
		return
	}
	_ = Conf
	return

}
