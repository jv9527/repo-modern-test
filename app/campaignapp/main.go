package main

import (
	"net/http"

	dCampaign "github.com/tokopedia/go-test/domain/campaign/delivery/http"
	redisCampaign "github.com/tokopedia/go-test/domain/campaign/repository/cache/redis"
	postgresCampaign "github.com/tokopedia/go-test/domain/campaign/repository/sql/postgres"
	uCampaign "github.com/tokopedia/go-test/domain/campaign/usecase"
	cassandraSegmentation "github.com/tokopedia/go-test/domain/segmentation/repository/db/cassandra"
)

func main() {
	cfg := initConfig("development")

	// init repositories
	postgresCampaign := postgresCampaign.New(cfg.Campaign.Postgres.ConnString)
	redisCampaign := redisCampaign.New(cfg.Campaign.Redis.Address)
	cassandraSegmentation := cassandraSegmentation.New(cfg.Segment.Cassandra.CassandraIP)

	// init usecases
	usecaseCampaign := uCampaign.New(postgresCampaign, redisCampaign, cassandraSegmentation)

	// init handlers
	handlerCampaign := dCampaign.New(*usecaseCampaign)

	// register handlers
	router := initHttpRouter(*handlerCampaign)

	http.ListenAndServe(":9000", router)
}

func initConfig(env string) *Config {
	return &Config{}
}

type Config struct {
	Segment struct {
		Cassandra struct {
			CassandraIP string
		}
	}
	Campaign struct {
		Postgres struct {
			ConnString string
		}
		Redis struct {
			Address string
		}
	}
}
