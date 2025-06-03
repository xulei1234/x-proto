package proto

import (
	"encoding/json"
)

// OSInfo Operation System information
type OSInfo struct {
	HostInfo      HostInfoStat      `json:"hostinfo"`
	InterfaceInfo []InterfaceStat   `json:"interfaceInfo"`
	MemInfo       VirtualMemoryStat `json:"memInfo"`
	CPUInfo       []CPUInfoStat     `json:"cpuInfo"`
}

// CPUInfoStat cpu information
type CPUInfoStat struct {
	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physicalId"`
	CoreID     string   `json:"coreId"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"modelName"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cacheSize"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
}

// InterfaceAddr is designed for represent interface addresses
type InterfaceAddr struct {
	Addr string `json:"addr"`
}

// InterfaceStat ...
type InterfaceStat struct {
	Speed        uint64          `json:"speed"`
	Index        int             `json:"index"`
	MTU          int             `json:"mtu"`          // maximum transmission unit
	Name         string          `json:"name"`         // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr string          `json:"hardwareaddr"` // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        []string        `json:"flags"`        // e.g., FlagUp, FlagLoopback, FlagMulticast
	Addrs        []InterfaceAddr `json:"addrs"`
}

// A HostInfoStat describes the host status.
// This is not in the psutil but it useful.
type HostInfoStat struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"bootTime"`
	Procs                uint64 `json:"procs"`           // number of processes
	OS                   string `json:"os"`              // ex: freebsd, linux
	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
	KernelArch           string `json:"kernelArch"`
	KernelVersion        string `json:"kernelVersion"` // version of the OS kernel (if available)
	VirtualizationSystem string `json:"virtualizationSystem"`
	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
	HostID               string `json:"hostid"`             // ex: uuid

}

// VirtualMemoryStat Memory usage statistics. Total, Available and Used contain numbers of bytes
// for human consumption.
//
// The other fields in this struct contain kernel specific values.
type VirtualMemoryStat struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total"`

	// RAM available for programs to allocate
	//
	// This value is computed from the kernel specific values.
	Available uint64 `json:"available"`

	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	Used uint64 `json:"used"`

	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	UsedPercent float64 `json:"usedPercent"`

	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	Free uint64 `json:"free"`

	// OS X / BSD specific numbers:
	// http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
	Active   uint64 `json:"active"`
	Inactive uint64 `json:"inactive"`
	Wired    uint64 `json:"wired"`

	// FreeBSD specific numbers:
	// https://reviews.freebsd.org/D8467
	Laundry uint64 `json:"laundry"`

	// Linux specific numbers
	// https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
	// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
	// https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
	Buffers        uint64 `json:"buffers"`
	Cached         uint64 `json:"cached"`
	Writeback      uint64 `json:"writeback"`
	Dirty          uint64 `json:"dirty"`
	WritebackTmp   uint64 `json:"writebacktmp"`
	Shared         uint64 `json:"shared"`
	Slab           uint64 `json:"slab"`
	SReclaimable   uint64 `json:"sreclaimable"`
	SUnreclaim     uint64 `json:"sunreclaim"`
	PageTables     uint64 `json:"pagetables"`
	SwapCached     uint64 `json:"swapcached"`
	CommitLimit    uint64 `json:"commitlimit"`
	CommittedAS    uint64 `json:"committedas"`
	HighTotal      uint64 `json:"hightotal"`
	HighFree       uint64 `json:"highfree"`
	LowTotal       uint64 `json:"lowtotal"`
	LowFree        uint64 `json:"lowfree"`
	SwapTotal      uint64 `json:"swaptotal"`
	SwapFree       uint64 `json:"swapfree"`
	Mapped         uint64 `json:"mapped"`
	VMallocTotal   uint64 `json:"vmalloctotal"`
	VMallocUsed    uint64 `json:"vmallocused"`
	VMallocChunk   uint64 `json:"vmallocchunk"`
	HugePagesTotal uint64 `json:"hugepagestotal"`
	HugePagesFree  uint64 `json:"hugepagesfree"`
	HugePageSize   uint64 `json:"hugepagesize"`
}

// IssueFile 下发文件结果
type IssueFile struct {
	Status string `json:"status"` // success/fail
	// TODO 补充
}

// Gossinfo 巡检结果
type Gossinfo struct {
	Status string `json:"status"`
	// TODO 补充
}

// cmd extra class
const (
	LogCmdClass = "log"
	MsgCmdClass = "msg"
)

// CmdExtra 定义
type CmdExtra struct {
	User    string `json:"user"`
	Timeout string `json:"timeout"`            // eg. 2m, using time.ParseDuration
	Class   string `json:"class"`              // deprecated
	MsgType string `json:"msg_type,omitempty"` // deprecated
	Code    uint32 `json:"code"`               // 标识上报的数据类型，MCodeLogLine 是使用异步日志按行上报
}

// NewCmdExtra with default user and timeout
func NewCmdExtra(class string) *CmdExtra {
	return &CmdExtra{
		Class:   class,
		Timeout: "1m",
		User:    "douyuops",
	}
}

// WithUser set user
func (ce *CmdExtra) WithUser(u string) *CmdExtra {
	ce.User = u
	return ce
}

// WithTimeout set user
func (ce *CmdExtra) WithTimeout(timeout string) *CmdExtra {
	ce.Timeout = timeout
	return ce
}

// WithCode set user
func (ce *CmdExtra) WithCode(code uint32) *CmdExtra {
	ce.Code = code
	return ce
}

// WithMsgType set user
func (ce *CmdExtra) WithMsgType(mt string) *CmdExtra {
	ce.MsgType = mt
	return ce
}

// Bytes ...
func (ce CmdExtra) Bytes() []byte {
	b, _ := json.Marshal(ce)
	return b
}

// LogCmdExtra must be upload as log, deprecated
func LogCmdExtra() []byte {
	extra, _ := json.Marshal(CmdExtra{
		Class: LogCmdClass,
	})
	return extra
}

// MsgCmdExtra must be upload as msg, deprecated
func MsgCmdExtra(t string) []byte {
	extra, _ := json.Marshal(CmdExtra{
		Class:   MsgCmdClass,
		MsgType: t,
	})
	return extra
}
