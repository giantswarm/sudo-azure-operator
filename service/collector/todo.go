package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelInstallation = "installation"
	labelClusterID    = "cluster_id"
)

var (
	ScheduleDesc *prometheus.Desc = prometheus.NewDesc(
		prometheus.BuildFQName("sudo-azure-operator", "todo", "info"),
		"Todo description of the sudo-azure-operator todo metric",
		[]string{
			labelInstallation,
			labelClusterID,
		},
		nil,
	)
)

type TodoConfig struct {
}

type Todo struct {
}

func NewTodo(config TodoConfig) (*Todo, error) {
	r := &Todo{}

	return r, nil
}

func (r *Todo) Collect(ch chan<- prometheus.Metric) error {
	return nil
}

func (r *Todo) Describe(ch chan<- *prometheus.Desc) error {
	ch <- ScheduleDesc

	return nil
}