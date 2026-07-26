package loggerv1

import "strings"

type LoggerFlag uint16

const (
	None        LoggerFlag = 0
	PrintLogs   LoggerFlag = 0x1
	PrintDebug  LoggerFlag = 0x2
	PrintSystem LoggerFlag = 0x4
	PrintWarns  LoggerFlag = 0x8
	PrintErrs   LoggerFlag = 0x10
)

var flAliases = map[LoggerFlag]string{
	None:        "",
	PrintLogs:   "logs",
	PrintDebug:  "debugs",
	PrintSystem: "sys",
	PrintWarns:  "warns",
	PrintErrs:   "errs",
}

func Parse(v string, buff *LoggerFlag) bool {
	if vs := strings.Split(v, ","); buff != nil && len(vs) > 0 {
		*buff = None

		for _, y := range vs {
			for flag, alias := range flAliases {
				if alias == y {
					*buff = (*buff).With(flag)
				}
			}
		}

		return *buff != None
	}

	return false
}

func (x LoggerFlag) With(v LoggerFlag) LoggerFlag {
	return x | v
}

func (x LoggerFlag) Is(v LoggerFlag) bool {
	return (x & v) != 0
}

func (x LoggerFlag) Closer() uint16 {
	return (uint16)(x)
}

func (x LoggerFlag) Contains(v LoggerFlag) bool {
	return (x & (1 << v.Closer())) != 0
}

func (x LoggerFlag) String() string {
	out := ""
	i := 1
	for y := x; y != 0; y = y >> 1 {
		if (y & 0x1) != 0 {
			if out != "" {
				out += ","
			}

			out += flAliases[LoggerFlag(i)]
		}

		i = i << 1
	}

	return out
}
