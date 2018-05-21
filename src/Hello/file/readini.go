package file

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config"
	"log"
)

var (
	configFile = flag.String("configfile", "./conf/demo.ini", "General configuration file")
)

//topic list
var TOPIC = make(map[string]string)
func main(){

	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("topicArr") {
		section, err := cfg.SectionOptions("topicArr")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("topicArr", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	fmt.Println(TOPIC)
	fmt.Println(TOPIC["debug"])
}
