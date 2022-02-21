package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAWSClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AWSClientAdapter{
		Client: AWSClient{},
	}

	err := a.CreateServer(1.0, 2.0)
	assert.Nil(t, err)
}

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}

	err := a.CreateServer(1.0, 2.0)
	assert.Nil(t, err)
}