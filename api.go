package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/JerryCheese/dlems/model"
	"github.com/JerryCheese/dlems/store"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
)

// APIConf _
type APIConf struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

// defaultAPIConf _
var defaultAPIConf APIConf

// initAPI create & start HTTP server
func initAPI(conf APIConf) {

	r := gin.Default()

	registerAPIRoutes(r)

	r.Run(fmt.Sprintf("%s:%d", conf.Addr, conf.Port))
}

// registerAPIRoutes _
func registerAPIRoutes(r *gin.Engine) {

	r.POST("/api/exp", handlerAddExp)
	r.GET("/api/exp", handlerListExp)

	r.POST("/api/exp/data", handlerAddExpData)
	r.GET("/api/exp/data", handlerListExpData)
}

func echoError(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"msg":  msg,
		"code": -1,
	})
}

func echoOK(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":  "ok",
		"code": 0,
	})
}

func echoOKData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"msg":  "ok",
		"code": 0,
		"data": data,
	})
}

/*****************************
	Experiment Manage
/*****************************/
func handlerAddExp(c *gin.Context) {
	var d model.DRun
	c.BindJSON(&d)

	if len(d.Name) == 0 {
		echoError(c, "name is empty")
		return
	} else if len(d.ExecStr) == 0 {
		echoError(c, "execStr is empty")
		return
	}

	d.StartTime = time.Now()
	m := d.AsMap()
	delete(m, "_id")
	m = store.AddMapData("run", m)

	echoOKData(c, m)
}

func handlerListExp(c *gin.Context) {
	name := strings.TrimSpace(c.Query("name"))
	_sort := strings.TrimSpace(c.Query("_sort"))

	//
	filter := map[string]interface{}{}
	if len(name) > 0 {
		filter["Name"] = name
	}

	// set sort params
	var sortOpts []bson.E
	if len(_sort) > 0 {
		_sortSlice := strings.Split(_sort, ",")
		for _, field := range _sortSlice {
			s := strings.Split(strings.TrimSpace(field), " ")
			key, asc := s[0], s[1]
			sortOpts = append(sortOpts, bson.E{key, cast.ToInt(asc)})
		}
	}

	data := store.FindWithSort("run", filter, sortOpts...)

	echoOKData(c, data)
}

/*****************************
	Experiment data Manage
/*****************************/
func handlerAddExpData(c *gin.Context) {
	var d model.DValue
	c.BindJSON(&d)

	if len(d.RunID) == 0 {
		echoError(c, "run_id is empty")
		return
	} else if len(d.Name) == 0 {
		echoError(c, "name is empty")
		return
	}

	d.Time = time.Now()

	m := d.AsMap()
	delete(m, "_id")
	m = store.AddMapData("data", m)

	echoOKData(c, m)
}

func handlerListExpData(c *gin.Context) {
	runID := strings.TrimSpace(c.Query("runID"))
	name := strings.TrimSpace(c.Query("name"))
	_sort := strings.TrimSpace(c.Query("_sort"))

	//
	filter := map[string]interface{}{}
	if len(runID) > 0 {
		filter["RunID"] = runID
	}
	if len(name) > 0 {
		filter["Name"] = name
	}

	// set sort params
	var sortOpts []bson.E
	if len(_sort) > 0 {
		_sortSlice := strings.Split(_sort, ",")
		for _, field := range _sortSlice {
			s := strings.Split(strings.TrimSpace(field), " ")
			key, asc := s[0], s[1]
			sortOpts = append(sortOpts, bson.E{key, cast.ToInt(asc)})
		}
	}

	data := store.FindWithSort("data", filter, sortOpts...)

	echoOKData(c, data)
}
