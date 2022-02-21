package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word string
	visit int
	UpdateAt *time.Time
}

// Clone 利用序列化进行深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

func (ks Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range ks {
		// 浅拷贝
		newKeywords[k] = v
	}

	for _, item := range updateWords {
		// 深拷贝
		newKeywords[item.word] = item.Clone()
	}

	return newKeywords
}