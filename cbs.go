package cbs

// RunningMode 表示运行模式
type RunningMode int

// 定义运行模式
const (
	RunningModeAuto RunningMode = 0 + iota
	RunningModeAgent
	RunningModeServer
	RunningModeBoth
	RunningModeNone
)
