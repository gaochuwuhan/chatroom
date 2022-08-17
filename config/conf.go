package config

import (
	"github.com/gaochuwuhan/goutils"
	"github.com/gaochuwuhan/goutils/pkg/cafe"
)

type VpConf struct {
	Env string
	Mongo MongoConf

}

type MongoConf struct{
	URI string
	UserName string
	Password string
}

var VP *VpConf

func Initialize() {
	env:=goutils.VP.GetString("env")
	mongo:=MongoConf{
		URI: goutils.VP.GetString(cafe.JoinStr(env,".mongo_uri")),
		UserName: goutils.VP.GetString(cafe.JoinStr(env,".mongo_user")),
		Password: goutils.VP.GetString(cafe.JoinStr(env,".mongou_pass")),
	}
	VP = &VpConf{
		Env:env,
		Mongo:mongo,
	}
}

