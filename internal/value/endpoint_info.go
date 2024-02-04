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

type EndpointInfo struct {
	FunctionCode     string
	UserAuditLogFlag bool
	// PiiLogFlag       	bool
}

var apiEndpointFlagsMap = map[string]EndpointInfo{
	"GET " + ApiUrlPrefix + "/local-collect": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"GET " + ApiUrlPrefix + "/local-collect-by-key": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"POST " + ApiUrlPrefix + "/local-collect": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"PUT " + ApiUrlPrefix + "/update-cheque-item": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"POST " + ApiUrlPrefix + "/chequeConfirm": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"GET " + ApiUrlPrefix + "/local-collect-branch-report": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"GET " + ApiUrlPrefix + "/summary-cheque-confirm-report": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"GET " + ApiUrlPrefix + "/summary-cheque-confirm-file": {
		FunctionCode:     "",
		UserAuditLogFlag: true,
	},
	"GET " + "/ping": {
		FunctionCode:     "5555",
		UserAuditLogFlag: true,
	},
}

func GetApiEndpointFlags(r *http.Request) (EndpointInfo, string, error) {
	pathSlice := strings.Split(r.RequestURI, "/") //path slice
	// s[len(s)-1] = ""
	lastPath := pathSlice[len(pathSlice)-1] //last element

	if len(lastPath) == 0 {
		//last /path is empty
		pathSlice = pathSlice[:len(pathSlice)-1] //remove '/'
	} else {
		//last /path not empty
		lastPath = strings.Split(lastPath, "?")[0] //remove query string
		_, err := strconv.Atoi(lastPath)           //check last /path is int
		if err == nil {
			//is int, set last /path format
			pathSlice[len(pathSlice)-1] = "{?}"
		}
	}

	pathFormat := strings.Join(pathSlice, "/")
	pathSlice = nil //clear mem

	endpoint := r.Method + " " + pathFormat

	val, exist := apiEndpointFlagsMap[endpoint]
	if exist {
		return val, endpoint, nil
	}
	return EndpointInfo{}, "", errors.New("no matching api endpoint")
}

func GetEndpointInfoFromCtx(ctx context.Context) EndpointInfo {
	l, isType := ctx.Value(EndpointInfoKey).(EndpointInfo)
	if isType {
		return l
	} else {
		return EndpointInfo{}
	}
}
