package config

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/magiconair/properties"
)

//type AppConfigProperties map[string]string

func GetProps(prop string) string {
	//dir, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	p := properties.MustLoadFile("C:\\ambiente\\workspace-go\\git\\GoApiTesting\\test\\config\\config.properties", properties.UTF8)
	value, _ := p.Get(prop)
	return value
}

// func ReadPropertiesFile() (AppConfigProperties, error) {
// 	config := AppConfigProperties{}
// 	//dir, err := os.Getwd()
// 	//if err != nil {
// 	//	log.Fatal(err)
// 	//}
// 	//userDir, _ := os.UserHomeDir()
// 	filename := "C:\\ambiente\\workspace-go\\git\\GoApiTesting\\test\\config\\config.properties"

// 	if len(filename) == 0 {
// 		return config, nil
// 	}
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if equal := strings.Index(line, "="); equal >= 0 {
// 			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
// 				value := ""
// 				if len(line) > equal {
// 					value = strings.TrimSpace(line[equal+1:])
// 				}
// 				config[key] = value
// 			}
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	return config, nil
// }
