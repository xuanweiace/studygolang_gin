package repos

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Topic struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
}

var topicIndexMap map[int]*Topic

// 把topic内容从磁盘加载进内存
func initTopicIndexMap(path string) error {
	topicIndexMap = make(map[int]*Topic)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var topic *Topic
	//topic = &Topic{} // 注意不能放到这里，共用一个指针了
	for scanner.Scan() {
		topic = &Topic{}
		text := scanner.Text()
		if err := json.Unmarshal([]byte(text), topic); err != nil {
			log.Println(err)
			return fmt.Errorf("err:", err)
		}
		topicIndexMap[topic.ID] = topic
	}
	return nil
}

type TopicDao struct{}

func (t *TopicDao) QueryById(id int) *Topic {
	return topicIndexMap[id] // 后续可以把这个map放到topicDao里
}

var topicDao *TopicDao
var topicOnce sync.Once // 单例模式，减少内存的浪费

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})
	return topicDao
}
