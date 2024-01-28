package value

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type ApiEndpointFlags struct {
	Auth         bool
	UserAuditLog bool
	PiiLog       bool
	SecurityLog  bool
}

var apiEndpointFlagsMap = map[string]ApiEndpointFlags{
	"GET/dogs/{?}": ApiEndpointFlags{
		Auth:         false,
		UserAuditLog: false,
		PiiLog:       false,
		SecurityLog:  false,
	},
	"GET/dogs": ApiEndpointFlags{
		Auth:         true,
		UserAuditLog: true,
		PiiLog:       true,
		SecurityLog:  true,
	},
	"POST/dogs/{?}": ApiEndpointFlags{
		Auth:         true,
		UserAuditLog: true,
		PiiLog:       true,
		SecurityLog:  true,
	},
	"PATCH/dogs/{?}": ApiEndpointFlags{
		Auth:         true,
		UserAuditLog: true,
		PiiLog:       true,
		SecurityLog:  true,
	},
	"DELETE/dogs/{?}": ApiEndpointFlags{
		Auth:         true,
		UserAuditLog: true,
		PiiLog:       true,
		SecurityLog:  true,
	},
}

func GetApiEndpointFlags(r *http.Request) (ApiEndpointFlags, error) {
	pathSlice := strings.Split(r.RequestURI, "/") //path slice
	// s[len(s)-1] = ""
	lastPath := pathSlice[len(pathSlice)-1] //last element

	if len(lastPath) == 0 {
		//last /path is empty
		pathSlice = pathSlice[:len(pathSlice)-1]	//remove '/'
	} else {
		//last /path not empty
		lastPath = strings.Split(lastPath, "?")[0] //remove query string
		_, err := strconv.Atoi(lastPath) //check last /path is int
		if err == nil {
			//is int, set last /path format
			pathSlice[len(pathSlice)-1] = "{?}"
		}
	}

	pathFormat := strings.Join(pathSlice, "/")
	pathSlice = nil //clear mem

	apiPath := r.Method + pathFormat

	val, exist := apiEndpointFlagsMap[apiPath]
	if exist {
		return val, nil
	}
	return ApiEndpointFlags{}, errors.New("no matching api endpoint")
}

func FromCtx(ctx context.Context) ApiEndpointFlags {
	l, isType := ctx.Value(ApiEndpointFlagsKey).(ApiEndpointFlags)
	if isType {
		return l
	} else {
		return ApiEndpointFlags{}
	}
}
