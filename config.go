package pentagonmodel

import (
    "log"
    "io/ioutil"
    "encoding/json"
    "os/user"
)

const (
    CONFIG_FILE = ".pentagon" // starting in user's home dir
)

type Config struct {
    GmailAddress, GmailPassword string
}

func GetConfig() *Config {

    usr, err := user.Current()
    if err != nil {
        log.Fatalln("Couldn't figure out current user:", err)
    }

    bytes, err := ioutil.ReadFile(usr.HomeDir + "/" + CONFIG_FILE)
    if err != nil {
        log.Fatalln("Error reading config file:", err)
    }

    config := &Config{}
    if err := json.Unmarshal(bytes, config); err != nil {
        log.Fatalln("Error parsing config:", err)
    }

    return config
}
