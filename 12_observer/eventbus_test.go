package observer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func sub1(msg1, msg2 string) {
	fmt.Printf("sub1, %s, %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s, %s\n", msg1, msg2)
}

func TestAsyncEventBus_Publish(t *testing.T) {
	bus := NewAsyncEventBus()
	assert.Nil(t, bus.Subscribe("topic:1", sub1))
	assert.Nil(t, bus.Subscribe("topic:1", sub2))
}

func ExampleAsyncEventBus_Publish() {
	bus := NewAsyncEventBus()
	_ = bus.Subscribe("topic:1", sub1)
	_ = bus.Subscribe("topic:1", sub2)
	bus.Publish("topic:1", "1", "2")
	bus.Publish("topic:1", "A", "B")
	time.Sleep(1 * time.Second)
	// unordered Output:
	// sub1, 1, 2
	// sub2, 1, 2
	// sub1, A, B
	// sub2, A, B
}