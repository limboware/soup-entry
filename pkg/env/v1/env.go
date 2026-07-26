package env1

import "strings"

type EnvType uint16

const envTypeSize = 16

const (
	UnknownEnv EnvType = iota
	DevEnv
	QAEnv EnvType = 1 << (iota - 1)
	StageEnv
	ProductionEnv
	LocalEnv
	ExperimentalEnv
)

var envTypes = map[EnvType]string{
	UnknownEnv:      "unknown",
	DevEnv:          "dev",
	QAEnv:           "qa",
	StageEnv:        "stage",
	ProductionEnv:   "prod",
	LocalEnv:        "local",
	ExperimentalEnv: "exp",
}

// Contains using like: et.Contains(DevEnv.And(LocalEnv))
func (et EnvType) Contains(t EnvType) bool {
	return (et & t) != 0
}

// Is eq. of  '==' operation
func (et EnvType) Is(t EnvType) bool {
	return et == t
}

func (et EnvType) With(t EnvType) EnvType {
	return et | t
}

func (et EnvType) Without(t EnvType) EnvType {
	return et &^ t
}

func (et EnvType) Int() int {
	return int(et)
}

func (et EnvType) String() string {
	if v, ok := envTypes[et]; ok {
		return v
	}

	ac := ""

	for i := 0; i < envTypeSize; i++ {
		if !et.Contains(EnvType(1 << i)) {
			continue
		}

		if v, ok := envTypes[EnvType(1<<i)]; ok {
			if ac == "" {
				ac = v
			} else {
				ac += "-" + v
			}
		}

		if et < 1<<(i+1) {
			break
		}
	}

	return ac
}

func isEnvAcronym(envA, buf string) bool {
	pos := strings.Index(envA, buf)

	return pos == 0 || (pos > 0 && envA[pos-1] == ';')
}

func ParseEnv(buf string) EnvType {
	env := UnknownEnv

	if len(buf) != 0 {
		// subs and some special acronyms
		m := map[EnvType]string{
			DevEnv:          "dv;devt;development",
			QAEnv:           "qa;quality",
			StageEnv:        "stg;stage;pp",
			ProductionEnv:   "production",
			LocalEnv:        "lcl;local",
			ExperimentalEnv: "x;el;ep;experimental",
		}

		var subs = []string{buf}
		if strings.Contains(buf, "-") {
			subs = strings.Split(buf, "-")
		}

		for _, sub := range subs {

			sub = strings.ToLower(strings.TrimSpace(sub))

			for k, v := range m {
				if isEnvAcronym(v, sub) {
					env = env.With(k)
				}
			}
		}
	}

	return env
}