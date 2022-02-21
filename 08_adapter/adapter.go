package adapter

import "fmt"

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AWSClient struct {}
func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu： %f, mem: %f", cpu, mem)
	return nil
}

type AWSClientAdapter struct {
	Client AWSClient
}
func (a *AWSClientAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.RunInstance(cpu, mem)
}


type AliyunClient struct{}
func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	return a.Client.CreateServer(int(cpu), int(mem))
}