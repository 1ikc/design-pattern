package singleton

type Singleton struct {}

var sg *Singleton

func init()  {
	sg = &Singleton{}
}

func GetInstance() *Singleton {
	return sg
}