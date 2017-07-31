package main

import (
	"github.com/plimble/ace"
	"fmt"
	"log"
	"devicehub/core"
	"os/user"
	"flag"
)

const EMPTY_STRING = ""
const UNKNOWN = "unknown"

func main() {
	var path, host string
	var port int
	usr, _ := user.Current()
	flag.StringVar(&path, "d", usr.HomeDir, "database file folder")
	flag.StringVar(&host, "h", "0.0.0.0", "host")
	flag.IntVar(&port, "p", 3000, "port")
	flag.Parse()

	log.Println("============== Device Hub Server ==============")
	log.Printf("database:\t[%s/android.txt,%s/ios.txt]\n", path, path)
	log.Printf("address:\t%s:%d\n", host, port)
	log.Println("===========================================")
	db1 := fmt.Sprintf("%s/android.txt", path)
	db2 := fmt.Sprintf("%s/ios.txt", path)
	indexer, err := core.New(db1, db2)
	if err != nil {
		log.Fatal(err)
	}
	server := ace.New()
	server.GET("/:name", func(c *ace.C) {
		name := c.Params.ByName("name")
		res := indexer.Search(&name)
		if res != EMPTY_STRING {
			c.JSON(200, res)
		} else {
			c.JSON(404, UNKNOWN)
		}
	})
	log.Println("server is started.")
	server.Run(fmt.Sprintf("%s:%d", host, port))
}
