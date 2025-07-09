package proc

import (
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func NewProcess(pid int32) (*process.Process, error) {
	return process.NewProcess(pid)
}

func getAllProcess() ([]*process.Process, error) {
	return process.Processes()
}

func SearchProcessByName(name string) (*process.Process, error) {
	ps, err := getAllProcess()
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		pName, err := p.Name()
		if err != nil {
			return nil, err
		}
		if pName == name {
			return p, nil
		}
	}
	return nil, nil
}

func SearchProcessByPrefix(name string) ([]*process.Process, error) {
	ps, err := getAllProcess()
	if err != nil {
		return nil, err
	}
	var ret []*process.Process
	for _, p := range ps {
		pName, err := p.Name()
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(pName, name) {
			ret = append(ret, p)
		}
	}
	return ret, nil
}
