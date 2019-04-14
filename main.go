package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/reynld/carbtographer/pkg/database"
	"github.com/reynld/carbtographer/pkg/server"
	"github.com/reynld/carbtographer/pkg/utils"
)

func main() {
	godotenv.Load()
	if err := utils.CheckEnviroment(); err != nil {
		log.Fatal(err)
	}

	serve := flag.Bool("serve", false, "runs server")
	migrate := flag.Bool("migrate", false, "migrates database")
	seed := flag.Bool("seed", false, "seeds database")
	flag.Parse()

	if len(os.Args) > 1 {
		if flag.NFlag() != 1 {
			fmt.Println("pass just one argument")
			flag.Usage()
			os.Exit(1)
		}

		s := server.Server{}
		s.Initialize()

		if *serve {
			s.Run()
		}
		if *migrate {
			database.RunMigrations(s.DB)
		}
		if *seed {
			database.RunSeeds(s.DB)
		}

	} else {
		fmt.Println("pass at least one argument")
		flag.Usage()
	}
}
