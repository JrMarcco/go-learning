package _chan

import (
	"reflect"
	"testing"
	"time"
)

func TestClearChan(t *testing.T) {

	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	t.Log(len(ch))

	go func() {
		// 清空 channel
		for range ch {
		}
	}()

	time.Sleep(time.Second)

	t.Log(len(ch))
}

func TestSelect(t *testing.T) {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	cases := createCases(ch1, ch2)

	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)

		if recv.IsValid() {
			// recv case
			t.Log("recv: ", cases[chosen].Dir, recv, ok)
		} else {
			// send case
			t.Log("send: ", cases[chosen].Dir, ok)
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	for i, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(i),
		})
	}

	return cases
}
