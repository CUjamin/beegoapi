package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	AnythingMap map[string]*Anything
)

func init() {
	AnythingMap = make(map[string]*Anything)
	AnythingMap["cheben"] = &Anything{"chebenid", "chebenname", "chebenpw"}
	AnythingMap["cheben2"] = &Anything{"chebenid2", "chebenname2", "chebenpw2"}
}

type Anything struct {
	Id       string
	Username string
	Password string
}

func AddAnything(a Anything) string {
	a.Id = "anything_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	AnythingMap[a.Id] = &a
	return a.Id
}


func GetAnythingMap() map[string]*Anything {
	return AnythingMap
}


func GetAnything(AnythingId string,) (anything *Anything, err error) {
	if v, ok := AnythingMap[AnythingId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}