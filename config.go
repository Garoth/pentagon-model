package pentagonmodel

import (
    "log"
    "io/ioutil"
    "encoding/json"
    "os/user"
)

const (
    CONFIG_FILE = ".pentagon/config" // starting in user's home dir
)

type Config struct {
    Workdir string

    GmailAddress, GmailPassword string

    Geeknote string
}

func GetConfig() *Config {

    usr, err := user.Current()
    if err != nil {
        log.Fatalln("Couldn't figure out current user:", err)
    }

    bytes, err := ioutil.ReadFile(usr.HomeDir + "/" + CONFIG_FILE)
    if err != nil {
        log.Println("Error reading config file:", err)
        log.Fatalln("Ensure ~/.pentagon/config exists.")
    }

    config := &Config{}
    if err := json.Unmarshal(bytes, config); err != nil {
        log.Fatalln("Error parsing config:", err)
    }

    return config
}
