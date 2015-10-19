package graphite

import (
	"fmt"
	"net"
	"testing"
)

// Change these to be your own graphite server if you so please
var graphiteHost = "carbon.hostedgraphite.com"
var graphitePort = 2003
var graphiteAddress = fmt.Sprintf("%s:%d", graphiteHost, graphitePort)

func TestNewGraphite(t *testing.T) {
	gh, err := NewGraphite(graphiteHost, graphitePort)
	if err != nil {
		t.Error(err)
	}

	if _, ok := gh.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}
}

func TestNewGraphiteFromAddress(t *testing.T) {
	gh, err := NewGraphiteFromAddress(graphiteAddress)
	if err != nil {
		t.Error(err)
	}

	if _, ok := gh.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}
}

// Uncomment the following method to test sending an actual metric to graphite
//
//func TestSendMetric(t *testing.T) {
//	gh, err := NewGraphite(graphiteHost, graphitePort)
//	if err != nil {
//		t.Error(err)
//	}
//	err = gh.SimpleSend("stats.test.metric11", "1")
//	if err != nil {
//		t.Error(err)
//	}
//}
