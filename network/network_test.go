package network

import (
	"testing"
	"time"
)

func TestFromNetworkTime(t *testing.T) {
	t.Log(TESTNET.EpochTime.UnixMilli())

	timestamp := TESTNET.Time(time.Hour * 2)
	t.Log(timestamp)
}
