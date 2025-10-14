package proc

import (
	"strings"

	"github.com/2981047480/tinybox/exception"
	"github.com/shirou/gopsutil/v4/process"
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
	return nil, exception.ProcessNotFoundError
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
	if ret == nil {
		return ret, exception.ProcessNotFoundError
	}
	return ret, nil
}

func GetContainerPIDOfProcess(pid int32) (ppid int32, err error) {
	p, err := NewProcess(pid)
	if err != nil {
		return 0, err
	}
	for {
		ppid, err = p.Ppid()
		if err != nil {
			return 0, err
		}
		if ppid == 1 {
			break
		}
		p, err = NewProcess(ppid)
		if err != nil {
			return 0, err
		}
	}
	return p.Pid, nil
}
