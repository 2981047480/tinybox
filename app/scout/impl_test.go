package scout

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func TestNewCpu(t *testing.T) {
	c, err := cpu.Info()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}

func TestCpuPersent(t *testing.T) {
	c, err := cpu.Percent(time.Second, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}
