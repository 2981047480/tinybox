package cpu

import (
	"time"

	"github.com/2981047480/tinybox/common/log"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
)

type Cpu struct {
	last_usage []float64 // CPU使用率，内部变量。
	CpuNum     int       `json:"cpu_num"`  // CPU个数
	CpuCore    int       `json:"cpu_core"` // CPU核数
	Arch       string    `json:"cpu_arch"` // CPU架构
	*load.AvgStat
}

const MEASURE_TIME_DURATION = time.Second

func NewCpu() *Cpu {
	return new(Cpu)
}

func (c *Cpu) Init() error {
	cpumap, err := cpu.Info()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	c.CpuNum = len(cpumap)
	c.CpuCore = int(cpumap[0].Cores)

	c.AvgStat, err = load.Avg()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	c.Arch, err = host.KernelArch()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (c *Cpu) GetUsage() []float64 {
	var err error
	c.last_usage, err = cpu.Percent(MEASURE_TIME_DURATION, true)
	if err != nil {
		log.Log(err.Error())
	}
	return c.last_usage
}

func (c *Cpu) GetLoad() *load.AvgStat {
	return c.AvgStat
}

func (c *Cpu) GetLoad1() float64 {
	return c.AvgStat.Load1
}

func (c *Cpu) GetLoad5() float64 {
	return c.AvgStat.Load5
}

func (c *Cpu) GetLoad15() float64 {
	return c.AvgStat.Load15
}

func (c *Cpu) GetUptime() uint64 {
	uptime, err := host.Uptime()
	if err != nil {
		log.Error(err.Error())
	}
	return uptime
}
