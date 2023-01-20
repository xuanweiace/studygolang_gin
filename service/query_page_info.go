package service

import (
	"fmt"
	"sync"
	"xwace/go_web_study/repos"
)

type PageInfoFlow struct {
	topic_id int
	postDao  *repos.PostDao
	topicDao *repos.TopicDao
}

func NewQueryPageInfoFlow(topic_id int) *PageInfoFlow {
	return &PageInfoFlow{
		topic_id: topic_id,
		postDao:  repos.NewPostDaoInstance(),
		topicDao: repos.NewTopicDaoInstance(),
	}
}

func (p *PageInfoFlow) Do() (*PageInfo, error) {
	if err := p.checkParam(); err != nil {
		return nil, err
	}
	topic, posts, err := p.prepareInfo()
	if err != nil {
		return nil, err
	}
	pageInfo := p.packPageInfo(topic, posts)
	return pageInfo, nil
}

func (p *PageInfoFlow) checkParam() error {
	if p.topic_id <= 0 {
		return fmt.Errorf("topic_id must > 0")
	}
	return nil
}

// 准备数据 也可以这个函数返回(topic,post,error) 但是缺点在于这样需要考虑子协程如何将数据拷贝回来。
func (p *PageInfoFlow) prepareInfo() (topic *repos.Topic, posts []*repos.Post, err error) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic = p.topicDao.QueryById(p.topic_id) // 返回的是一个topic
	}()
	go func() {
		defer wg.Done()
		posts = p.postDao.QueryByParentId(p.topic_id)
	}()
	wg.Wait()
	return
}

type PageInfo struct {
	Topic *repos.Topic  //记得大写
	Posts []*repos.Post //记得大写
}

func (p *PageInfoFlow) packPageInfo(topic *repos.Topic, posts []*repos.Post) *PageInfo {
	return &PageInfo{
		Topic: topic,
		Posts: posts,
	}

}
