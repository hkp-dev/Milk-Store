package database

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestConnect(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful connection", func(mt *mtest.T) {
		Connect()
		if Client == nil {
			t.Error("Expected Client to be initialized")
		}
	})
}

func TestDisconnect(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("successful disconnection", func(mt *mtest.T) {
		Connect()
		Disconnect()
		if Client != nil {
			t.Error("Expected Client to be nil after disconnection")
		}
	})
}
