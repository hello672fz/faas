package policies

import (
	"hash/fnv"
	"math/rand"
	"net/url"
	"time"
)

// TODO: Policy
func SelectNodeByPolicy(urls []*url.URL) int {
	if len(urls) == 0 {
		return -1
	}
	//todo:
	//find the node with idle container(local ctn table) or scale a new one to a node using random/haxi/cxl-snapshot scheduled
	//
	//

	randomIndex := randomSelect(urls)
	randomIndex = 0
	//hashIndex := hashSelect(urls, name)
	//heatIndex:=

	return randomIndex
}

func randomSelect(urls []*url.URL) int {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为种子
	index := rand.Intn(len(urls))    // 生成一个 0 到 len(arr)-1 之间的随机索引
	return index
}

func hashSelect(urls []*url.URL, key string) int {
	hash := fnv.New32a()                      // 创建一个哈希实例
	hash.Write([]byte(key))                   // 使用输入的 key 来生成哈希值
	index := hash.Sum32() % uint32(len(urls)) // 获取哈希值并将其映射到数组的索引
	return int(index)
}

func heatSelect(urls []*url.URL, key string) int {
	index := 0
	return index
}
