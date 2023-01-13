package common

import (
	"context"
	"time"
)

const TotalRequestTimeOut = 2500 * time.Millisecond
const ControllerContextTimeOut = 2000 * time.Millisecond
const ServiceContextTimeOut = 1000 * time.Millisecond
const ModelContextTimeOut = 1000 * time.Millisecond
const DatabaseClientTimeOut = 2000 * time.Millisecond

func NewContext(t time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	return ctx, cancel
}
