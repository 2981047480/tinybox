package cpu

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/stretchr/testify/assert"
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

func TestGetUsage(t *testing.T) {
	c := NewCpu()
	c.Init()
	t.Log(c.GetUsage())
	assert.Equal(t, c.last_usage, c.GetUsage())
}

func TestGetLoad(t *testing.T) {
	c := NewCpu()
	c.Init()
	a := c.GetLoad()
	t.Log(a)
}
