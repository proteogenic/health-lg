package health_lg

import (
	"fmt"

	"github.com/gocraft/health"
	"github.com/pressly/lg"
)

// Sink emits errors to lg.
type Sink struct{}

func (s *Sink) EmitEvent(job string, event string, kvs map[string]string) {
	lg.WithField("job", job).
		WithField("event", event).
		WithFields(mapFields(kvs)).
		Info(job + " : " + event)
}

func (s *Sink) EmitEventErr(job string, event string, inputErr error, kvs map[string]string) {
	lg.WithField("job", job).
		WithField("event", event).
		WithField("panic", inputErr.Error()).
		WithFields(mapFields(kvs)).
		Error(job + " : " + event)
}

func (s *Sink) EmitTiming(job string, event string, nanos int64, kvs map[string]string) {
	lg.WithField("job", job).
		WithField("event", event).
		WithField("elapsed", float64(nanos)/1000000.0).
		WithFields(mapFields(kvs)).
		Info(job + " : " + event)
}

func (s *Sink) EmitGauge(job string, event string, value float64, kvs map[string]string) {
	lg.WithField("job", job).
		WithField("event", event).
		WithField("gauge", fmt.Sprintf("%g", value)).
		WithFields(mapFields(kvs)).
		Info(job + " : " + event)
}

func (s *Sink) EmitComplete(job string, status health.CompletionStatus, nanos int64, kvs map[string]string) {
	lg.WithField("job", job).
		WithField("status", status.String()).
		WithField("elapsed", float64(nanos)/1000000.0).
		WithFields(mapFields(kvs)).
		Info(job + " : " + status.String())
}

func mapFields(kvs map[string]string) map[string]interface{} {
	if kvs == nil {
		return nil
	}

	fields := make(map[string]interface{})
	for k, v := range kvs {
		fmt.Printf("%s:%v\n", k, v)
		fields[k] = v
	}
	return fields
}
