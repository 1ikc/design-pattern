package singleton

import "sync"

var (
	lazySg *Singleton
	once sync.Once
)

func GetLazyInstance() *Singleton {
	if lazySg == nil {
		once.Do(func() {
			lazySg = &Singleton{}
		})
	}

	return lazySg
}