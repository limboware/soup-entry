package configsv2

import (
	"maps"
	"strconv"
	"strings"

	errnov1 "github.com/rejchev/errno"
)

const OptionPrefix = '-'

type Config map[string]string

var v = Config{}

func Get() Config {
	return v
}

func GetConfig() Config {
	return Get()
}

func GetConfigIntValue[T int | int64 | int32](k string, buff *T) bool {
	return GetConfigValue(Get(), k, buff, func(s string, t *T) bool {
		c, err := strconv.ParseInt(s, 10, 64)

		if err != nil {
			return false
		}

		*buff = (T)(c)
		return true
	})
}

func GetConfigStringValue(k string, buff *string) bool {
	return GetConfigValue(Get(), k, buff, func(s string, t *string) bool {
		*t = s
		return true
	})
}

func GetConfigValue[T any](v Config, k string, buff *T, castFn func(string, *T) bool) bool {
	if x, ok := v[k]; ok {
		return castFn(x, buff)
	}

	return false
}

func (x Config) Parse(args []string) errnov1.Code {
	valAwait := false

	pathBinIdx := strings.LastIndexFunc(args[0], func(r rune) bool {
		return r == '\\' || r == '/'
	})

	if pathBinIdx == -1 {
		pathBinIdx = len(args[0])
	}

	x["bin"] = args[0][0:pathBinIdx]

	for i := range len(args) {
		arg := strings.TrimSpace(args[i])

		if arg[0] == OptionPrefix {
			x[arg[1:]] = ""
			valAwait = true
			continue
		}

		if valAwait {
			x[args[i-1][1:]] = arg
			valAwait = false
		}
	}

	// fmt.Printf("%s\n", x["bin"])

	return errnov1.OK
}

// os.Environ
func (x Config) ParseEnvi(v []string) errnov1.Code {
	prefix := "LIMBOWARE_"

	for _, y := range v {
		if strings.Index(y, prefix) == 0 {
			if b := strings.Split(y[len(prefix):], "="); len(b) == 2 {
				x[strings.ToLower(b[0])] = b[1]
			}
		}
	}

	return errnov1.OK
}

func (x Config) Value(v string) string {
	return x.ValueA(v, "")
}

func (x Config) ValueA(v, def string) string {
	if y, ok := x[v]; ok {
		return y
	}

	return ""
}

func (x Config) Exist(v string) bool {
	_, ok := x[v]
	return ok
}

func (x Config) Append(v Config) Config {
	for k, v := range v {
		if _, ok := x[k]; !ok {
			x[k] = v
		}
	}

	return x
}

func (x Config) Insert(v Config) Config {
	maps.Copy(x, v)

	return x
}
