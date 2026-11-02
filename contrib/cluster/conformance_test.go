package cluster_test

import (
	"testing"

	"github.com/lynxland/lynx/contrib/cluster"
	"github.com/lynxland/lynx/contrib/cluster/clustertest"
)

// TestMemoryConformance：进程内 Coordinator 接入 clustertest 契约套件
// （consul / cluster-redis 在各自模块接入）。
func TestMemoryConformance(t *testing.T) {
	clustertest.TestCoordinator(t, func(t *testing.T) cluster.Coordinator {
		return cluster.NewMemory()
	})
}
