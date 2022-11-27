package business

import (
	"context"
	"time"
)

type ObjectMetadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InternalContext struct {
	TrackerID string
}

func NewInternalContext(trackerID string) InternalContext {
	return InternalContext{
		TrackerID: trackerID,
	}
}

func (ic InternalContext) ToContext() context.Context {
	ctx := context.WithValue(context.Background(), "tracker_id", ic.TrackerID)
	return ctx
}
