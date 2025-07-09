package proc

import (
	"testing"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/assert"
)

func TestGetProcessName(t *testing.T) {
	p, err := process.NewProcess(1754)
	if err != nil {
		t.Fatal(err)
	}
	cmdline, err := p.CmdlineSlice()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cmdline)
	assert.Equal(t, cmdline[0], "ai-agent")
}

func TestSearchProcessByName(t *testing.T) {
	p, err := SearchProcessByName("ai-agent")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p.Pid)
	assert.Equal(t, p.Pid, int32(1754))
}
