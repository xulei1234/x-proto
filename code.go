package x_proto

const (
	MCodeLogLine          uint32 = iota // 日志行信息
	MCodeCommon                         // 一般默认消息，不关注类型
	MCodeNodeInfo                       // 上报简单主机信息, 便于cruiser-channel管理
	MCodeSysInfo                        // 系统信息
	MCodeConfirm                        // 任务下发收到确认
	MCodeGossInfo                       // goss 运行结果
	MCodeIssueFile                      // 推送文件
	MCodeIssueUser                      // 推送跳板机用户结果
	MCodePingCheck                      // ping 检测输出
	MCodeLLDPTool                       // lldptool 输出信息
	MCodeAdminTask                      // cruiser-admin 批量执行命令结果
	MCodeServiceDeploy                  // 服务(中间件)部署
	MCodeCDNDetect                      // cdn探测
	MCodeCDNDetectV2                    // cdn探测v2
	MCodeProbeCheck                     //批量探测脚本
	MCodeCDNEdgeCollector               // cdn边缘采集
	MCodeRobot                          // 机器人
)

// 过渡使用，最终全部改为 code 表示
var str2Code = map[string]uint32{
	"agentinfo":        MCodeNodeInfo,
	"sysinfo":          MCodeSysInfo,
	"cmdConfirm":       MCodeConfirm,
	"gossinfo":         MCodeGossInfo,
	"issuefile":        MCodeIssueFile,
	"issueuser":        MCodeIssueUser,
	"pingcheck":        MCodePingCheck,
	"lldptool":         MCodeLLDPTool,
	"probecheck":       MCodeProbeCheck,
	"admintask":        MCodeAdminTask,
	"servicedeploy":    MCodeServiceDeploy,
	"cdndetect":        MCodeCDNDetect,
	"cdndetectv2":      MCodeCDNDetectV2,
	"cdnedgecollector": MCodeCDNEdgeCollector,
	"robot":            MCodeRobot,
}

// Str2Code by query table
func Str2Code(s string) uint32 {
	if code, ok := str2Code[s]; ok {
		return code
	}
	return MCodeCommon
}

// msg type
const (
	MsgTypeAgentInfo  = "agentinfo"  // 上报简单主机信息, 便于cruiser-channel管理
	MsgTypeSysInfo    = "sysinfo"    // 上报主机信息
	MsgTypeCmdConfirm = "cmdConfirm" // 命令收到确认消息
	MsgTypeGossInfo   = "gossinfo"   // 上报goss巡检信息
	MsgTypeIssueFile  = "issuefile"  // 上报文件下发结果
	MsgTypeIssueUser  = "issueuser"  // 上报跳板机下发结果
	MsgTypePingCheck  = "pingcheck"  // 上报ping检测结果
	MsgTypeUnknown    = "unknown"    // 上报未知类型信息
)
