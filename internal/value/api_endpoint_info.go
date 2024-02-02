package value

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const (
	ApiUrlPrefix = "/apis/conf/v1"
)

type ApiEndpointInfo struct {
	FunctionCode	string
	UserAuditLogFlag 	bool
	// PiiLogFlag       	bool
}

var apiEndpointFlagsMap = map[string]ApiEndpointInfo{
	"GET"+ApiUrlPrefix+"/local-collect" : ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"GET"+ApiUrlPrefix+"/local-collect-by-key": ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"POST"+ApiUrlPrefix+"/local-collect": ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"PUT"+ApiUrlPrefix+"/update-cheque-item": ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"POST"+ApiUrlPrefix+"/chequeConfirm": ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"GET"+ApiUrlPrefix+"/local-collect-branch-report" : ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"GET"+ApiUrlPrefix+"/summary-cheque-confirm-report" : ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
	"GET"+ApiUrlPrefix+"/summary-cheque-confirm-file" : ApiEndpointInfo{
		FunctionCode:	"",
		UserAuditLogFlag: 	true,
	},
}

func GetApiEndpointFlags(r *http.Request) (ApiEndpointInfo, error) {
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
	return ApiEndpointInfo{}, errors.New("no matching api endpoint")
}

func GetApiEndpointInfoFromCtx(ctx context.Context) ApiEndpointInfo {
	l, isType := ctx.Value(ApiEndpointInfoKey).(ApiEndpointInfo)
	if isType {
		return l
	} else {
		return ApiEndpointInfo{}
	}
}
