package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/trevatk/anastasia/internal/adapter/port/rpc/msg"
)

func (b *Bundle) modifyAccessControlList(ctx context.Context, payload []byte) (*msg.ModifyACLResponse, error) {

	// unmarshal request
	var modify msg.ModifyACLPayload
	err := json.Unmarshal(payload, &modify)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload %v", err)
	}

	// validate request
	v := modify.Validate()
	if v != nil {
		return nil, fmt.Errorf("invalid policy provided %v", err)
	}

	return b.ac.ModifyAccessControlList(ctx, &modify)
}
