package conf

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

//读入配置文件
func ReadSettingsFromFile(settingFilePath string) (config *Config) {
	//尝试打开文件
	var Config Config
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		panic(any("No such file named " + settingFilePath))
	}
	defer jsonFile.Close()
	
	//将数据读出
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		logrus.Error(err)
	}
	config = &Config
	return config
}