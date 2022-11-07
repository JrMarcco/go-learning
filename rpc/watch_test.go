package rpc

import (
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

func watch(client *rpc.Client) {
	go func() {
		var key string
		err := client.Call("KVStoreService.Watch", 30, &key)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("watch: %s\n", key)
	}()

}

func set(client *rpc.Client, kv [2]string) {
	err := client.Call("KVStoreService.Set", kv, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
}

func TestWatchServer(t *testing.T) {
	if err := rpc.RegisterName("KVStoreService", NewKVStoreService()); err != nil {
		log.Fatal("Register error: ", err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}

func TestWatchClient(t *testing.T) {
	time.Sleep(time.Second)

	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	watch(client)
	set(client, [2]string{"key", "val1"})
	set(client, [2]string{"key", "val2"})

	time.Sleep(time.Second)
}
