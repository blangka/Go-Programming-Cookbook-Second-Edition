package main

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

type key int

const logFields key = 0

func initialize() {
	log.SetHandler(text.New(os.Stdout))
	ctx := context.Background()
	ctx, e := fromContext(ctx, log.Log)

	ctx = withField(ctx, "id", "123")
	e.Info("starting")
	withField(ctx, "name", "Go Cookbook")
	e.Info("after gatherName")
	withFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
	e.Info("after gatherLocation")
}

func fromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

func withField(ctx context.Context, key string, value interface{}) context.Context {
	return withFields(ctx, log.Fields{key: value})
}

func withFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	return context.WithValue(ctx, logFields, f)
}

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

func main() {
	initialize()
}
