package controlproto

import (
	"testing"

	"github.com/vadimkozak/centrifuge/v1/internal/controlpb"

	"github.com/stretchr/testify/require"
)

func TestEncoder(t *testing.T) {
	encoder := NewProtobufEncoder()

	cmd := &controlpb.Command{
		UID:    "test",
		Method: controlpb.MethodTypeDisconnect,
		Params: controlpb.Raw("{}"),
	}
	d, err := encoder.EncodeCommand(cmd)
	require.NoError(t, err)
	require.NotNil(t, d)

	node := &controlpb.Node{
		UID:         "test",
		Name:        "test name",
		Version:     "v1.0.0",
		NumChannels: 2,
		NumClients:  3,
		NumUsers:    1,
		Uptime:      12,
		Metrics: &controlpb.Metrics{
			Interval: 60,
			Items: map[string]float64{
				"item": 1,
			},
		},
	}
	d, err = encoder.EncodeNode(node)
	require.NoError(t, err)
	require.NotNil(t, d)

	disconnect := &controlpb.Disconnect{
		User: "test",
	}
	d, err = encoder.EncodeDisconnect(disconnect)
	require.NoError(t, err)
	require.NotNil(t, d)

	unsub := &controlpb.Unsubscribe{
		User:    "test",
		Channel: "test channel",
	}
	d, err = encoder.EncodeUnsubscribe(unsub)
	require.NoError(t, err)
	require.NotNil(t, d)
}
