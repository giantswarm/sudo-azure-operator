package collector

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/kubernetes"
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
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger
}

type Todo struct {
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger
}

func NewTodo(config TodoConfig) (*Todo, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	r := &Todo{
		K8sClient: config.K8sClient,
		Logger: config.Logger,
	}

	return r, nil
}

func (r *Todo) Collect(ch chan<- prometheus.Metric) error {
	return nil
}

func (r *Todo) Describe(ch chan<- *prometheus.Desc) error {
	ch <- ScheduleDesc

	return nil
}