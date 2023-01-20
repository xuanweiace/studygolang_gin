package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"xwace/go_web_study/repos"
	"xwace/go_web_study/service"
)

func Init() error {
	if err := repos.Init("./data"); err != nil {
		return err
	}
	return nil
}
func main() {
	if err := Init(); err != nil {
		panic(err)
	}
	e := gin.Default()
	e.GET("topic/get/:id", func(c *gin.Context) {
		topic_id_str := c.Param("id")
		topic_id, _ := strconv.ParseInt(topic_id_str, 10, 0)
		queryPageInfoFlow := service.NewQueryPageInfoFlow(int(topic_id)) // 所以最好还是都用int64吧
		pageInfo, err := queryPageInfoFlow.Do()
		var data interface{}
		if err != nil {
			data = pack_error(err.Error())
		} else {
			data = pack_success(pageInfo)
		}
		log.Println(pageInfo)
		c.JSON(http.StatusOK, data)
	})
	e.Run("127.0.0.1:9000")
}

// 注意大写，然后再用json这个tag给还原为小写
type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func pack_success(data interface{}) *CommonResp {
	return &CommonResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func pack_error(msg string) *CommonResp {
	return &CommonResp{
		Code: -1,
		Msg:  msg,
		Data: nil,
	}
}
