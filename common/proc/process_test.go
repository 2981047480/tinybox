package proc

import (
	"testing"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/assert"
)

func TestGetProcessName(t *testing.T) {
	p, err := process.NewProcess(530179)
	if err != nil {
		t.Fatal(err)
	}
	cmdline, err := p.CmdlineSlice()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cmdline)
	assert.Equal(t, cmdline[0], "nginx: master process nginx -g daemon off;")
}

func TestSearchProcessByName(t *testing.T) {
	p, err := SearchProcessByName("nginx")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p.Pid)
	assert.Equal(t, p.Pid, int32(530179))
}

func TestSearchProcessByPrefix(t *testing.T) {
	p, err := SearchProcessByPrefix("ngin")
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range p {
		t.Log(item.Pid)
	}
	// assert.Equal(t, p.Pid, int32(530179))
}

func TestGetContainerIDOfProcess(t *testing.T) {
	ppid, err := GetContainerIDOfProcess(int32(530252))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ppid)
}
