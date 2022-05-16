package client

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

// Node 表示参与选举的节点。
type Node struct {
	elect     *concurrency.Election
	prefix    string
	session   *concurrency.Session
	leaderKey string
}

// New 返回一个新创建的Node对象。
func New(client *clientv3.Client, service string, leaderKey string, ttl int) (*Node, error) {
	if client == nil {
		return nil, fmt.Errorf("unexpected client, nil")
	}

	prefix := fmt.Sprintf("/m/elect/%s/", service)
	session, err := concurrency.NewSession(client, concurrency.WithTTL(ttl))
	if err != nil {
		err = fmt.Errorf("new session error, %w", err)
		return nil, err
	}

	elect := concurrency.NewElection(session, prefix)
	return &Node{
		elect:     elect,
		prefix:    prefix,
		session:   session,
		leaderKey: leaderKey,
	}, nil
}

// WaitToBeLeader 阻塞等待成为leader。
func (n *Node) WaitToBeLeader(timeout time.Duration) error {
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	return n.elect.Campaign(ctx, n.leaderKey)
}

// ObserverDone 将session断开、过期事件关联上回调事件。例如：leader断开服务器连接了，丢失leadership，则os.Exit(1)。
func (n *Node) ObserverDone(cb func()) {
	ch := n.session.Done()
	for {
		_, ok := <-ch
		if !ok {
			break
		}
	}

	if cb != nil {
		cb()
	}
}

// Close 关闭内部session。
func (n *Node) Close() error {
	return n.session.Close()
}
