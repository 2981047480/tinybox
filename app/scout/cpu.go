package scout

import (
	"time"

	"github.com/2981047480/tinybox/common/log"
	"github.com/2981047480/tinybox/ioc"
	"github.com/shirou/gopsutil/v3/cpu"
)

type Cpu struct {
	usage   []float64 // CPU使用率，内部变量。
	CpuNum  int       `json:"cpu_num"`  // CPU个数
	CpuCore int       `json:"cpu_core"` // CPU核数
}

const MEASURE_TIME_DURATION = time.Second

func NewCpu() *Cpu {
	// c, err := cpu.Info()
	c := new(Cpu)
	return c
}

func (c *Cpu) Name() string {
	return CPU
}

func (c *Cpu) Init() error {
	cpumap, err := cpu.Info()
	if err != nil {
		return err
	}
	c.CpuNum = len(cpumap)
	c.CpuCore = int(cpumap[0].Cores)
	ioc.DefaultIoc.Register(CPU, c)
	return nil
}

func (c *Cpu) GetUsage() []float64 {
	var err error
	c.usage, err = cpu.Percent(MEASURE_TIME_DURATION, true)
	if err != nil {
		log.Log(err.Error())
	}
	return c.usage
}
