package cgroup

type Cgroup interface {
	Init() error
	Get() (string, error)
}
