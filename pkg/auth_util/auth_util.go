package auth_util

import (
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"os"
)

// secret.yaml の内容がマウントされる
// $ echo -n "social-graph-manager: 123" | base64
// => c29jaWFsLWdyYXBoLW1hbmFnZXI6IDEyMw==
const callersFilename = "/etc/delinkcious/mutual-auth.yaml"

var callersByName = map[string]string{}
var callersByToken = map[string][]string{}

func init() {
	if os.Getenv("DELINKCIOUS_MUTUAL_AUTH") == "false" {
		return
	}

	data, err := os.ReadFile(callersFilename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, callersByName)
	if err != nil {
		panic(err)
	}

	for caller, token := range callersByName {
		callersByToken[token] = append(callersByToken[token], caller)
	}
}

func GetToken(caller string) string {
	return callersByName[caller]
}

func HasCaller(caller string, token string) bool {
	for _, c := range callersByToken[token] {
		if c == caller {
			return true
		}
	}

	return false
}
