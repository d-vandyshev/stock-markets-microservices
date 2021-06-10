package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var conf Config

func init() {
	conf.SetFromEnvOrDie()
}

func main() {
	fmt.Println("conf.DataPath:", conf.DataPath)
	fmt.Println("conf.RabbitmqUrl:", conf.RabbitmqUrl)

	stocks, err := ioutil.ReadDir(conf.DataPath)
	if err != nil {
		fmt.Printf("ERROR. Can't read data directory %s", conf.DataPath)
		os.Exit(1)
	}

	for _, stock := range stocks {
		if !stock.IsDir() {
			continue
		}
		fmt.Println(stock.Name())
		filenames, err := ioutil.ReadDir(filepath.Join(conf.DataPath, stock.Name(), "csv"))
		if err != nil {
			fmt.Printf("ERROR. Can't read data directory %s", conf.DataPath)
			os.Exit(1)
		}
		// TODO Run RabbitMQ in docker
		// TODO Put the list in RabbitMQ
		fmt.Println(len(filenames))
	}
}
