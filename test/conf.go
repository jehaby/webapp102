package test

import "github.com/jehaby/webapp102/config"

func getConf() config.C {
	return config.C{
		DB: config.DB{
			User:     "postgres",
			Database: "webapp",
			Port:     "65432",
			Host:     "localhost",
		},
	}
}
