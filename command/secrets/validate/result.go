package validate

import (
	"bytes"
	"context"
	"fmt"
	"github.com/0xPolygon/polygon-edge/command/helper"
	"github.com/0xPolygon/polygon-edge/server/proto"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

type SecretsValidateResult struct {
	NodeID       string `json:"nodeId"`
	Address      string `json:"address"`
	ValidateInfo string `json:"validateInfo"`
}

func (r *SecretsValidateResult) GetOutput() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n[VALIDATE INFO]\n")
	buffer.WriteString(helper.FormatKV([]string{
		fmt.Sprintf("Node ID|%s", r.NodeID),
		fmt.Sprintf("Node Address|%s", r.Address),
		fmt.Sprintf("Validate Info| %s", r.ValidateInfo),
	}))
	buffer.WriteString("\n")

	return buffer.String()
}

func getSystemStatus(grpcAddress string) (*proto.ServerStatus, error) {
	client, err := helper.GetSystemClientConnection(
		grpcAddress,
	)
	if err != nil {
		return nil, err
	}

	return client.GetStatus(context.Background(), &empty.Empty{})
}
