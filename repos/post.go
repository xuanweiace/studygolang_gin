package repos

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

//////////////////////////////////////////////////

type Post struct {
	ID         int    `json:"id"`
	ParentID   int    `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
}

//var postIndexMap map[int][]*Post// 后续可以把这个map放到topicDao里 缺点是让Dao层带有了数据的属性。应该是行为和数据分开，让DAO只负责数据，这是DDD。

// 把post内容从磁盘加载进内存
func (p *PostDao) initPostIndexMap(path string) error {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var post *Post
	for scanner.Scan() {
		post = &Post{}
		text := scanner.Text()
		if err := json.Unmarshal([]byte(text), post); err != nil {
			log.Println(err)
			return fmt.Errorf("err:", err)
		}
		_, ok := p.postIndexMap[post.ParentID]
		if !ok {
			p.postIndexMap[post.ParentID] = make([]*Post, 0)
		}
		p.postIndexMap[post.ParentID] = append(p.postIndexMap[post.ParentID], post)
	}
	return nil
}

func (p *PostDao) QueryByParentId(id int) []*Post {
	return p.postIndexMap[id]
}

type PostDao struct {
	postIndexMap map[int][]*Post
}

var postDao *PostDao
var postOnce sync.Once // 单例模式，减少内存的浪费

func NewPostDaoInstance() *PostDao {
	topicOnce.Do(func() {
		postDao = &PostDao{
			postIndexMap: make(map[int][]*Post),
		}
	})
	return postDao
}
