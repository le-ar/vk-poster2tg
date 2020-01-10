package cores

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type AppSettings struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	TgApi       string `json:"api_token"`
	TgChannel   string `json:"channel_name"`
	SortPostsBy string `json:"sort_posts_by"`
}

func IsSettingsFileExists() bool {
	_, err := os.Stat("settings.json")
	return err == nil
}

func ReadOrInit() *AppSettings {
	if IsSettingsFileExists() {
		conf, err := ReadAppSettingFromFile()
		if err == nil {
			return conf
		}
	}

	return InitAppSetting()
}

func ReadAppSettingFromFile() (*AppSettings, error) {
	jsonFile, err := os.Open("settings.json")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var appSettings AppSettings

	json.Unmarshal(byteValue, &appSettings)

	return &appSettings, nil
}

func InitAppSetting() *AppSettings {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email for vk-poster: ")
	email, _ := reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)
	email = strings.Replace(email, "\r", "", -1)

	fmt.Print("Enter password for vk-poster: ")
	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)
	password = strings.Replace(password, "\r", "", -1)

	fmt.Print("Enter api token for telegram bot: ")
	tgApi, _ := reader.ReadString('\n')
	tgApi = strings.Replace(tgApi, "\n", "", -1)
	tgApi = strings.Replace(tgApi, "\r", "", -1)

	fmt.Print("Enter telegram channel name: ")
	channelName, _ := reader.ReadString('\n')
	channelName = strings.Replace(channelName, "\n", "", -1)
	channelName = strings.Replace(channelName, "\r", "", -1)
	if channelName[0] != '@' {
		channelName = "@" + channelName
	}

	fmt.Print("Enter posts sort alg (\"IA\" or \"Percent\"): ")
	sortPostsBy, _ := reader.ReadString('\n')
	sortPostsBy = strings.Replace(sortPostsBy, "\n", "", -1)
	sortPostsBy = strings.Replace(sortPostsBy, "\r", "", -1)
	if channelName[0] != '@' {
		channelName = "@" + channelName
	}

	appSettings := AppSettings{
		Email:       email,
		Password:    password,
		TgApi:       tgApi,
		TgChannel:   channelName,
		SortPostsBy: sortPostsBy,
	}

	settingsJson, _ := json.Marshal(appSettings)

	err := ioutil.WriteFile("settings.json", settingsJson, 0644)

	if err != nil {
		fmt.Println(err)
	}

	return &appSettings
}
