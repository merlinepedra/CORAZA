// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package variables

import (
	"fmt"
	"strings"
)

// This file repeats the same content many times in order to make access
// efficient for seclang and transactions

// RuleVariable is used for the rule to identify information
// Each RuleVariable is unique and represents a variable
type RuleVariable byte

const (
	// Unknown is the default value for a variable
	// it's using for testing and error catching
	Unknown                       RuleVariable = iota
	ResponseContentType           RuleVariable = iota
	UniqueID                      RuleVariable = iota
	ArgsCombinedSize              RuleVariable = iota
	AuthType                      RuleVariable = iota
	FilesCombinedSize             RuleVariable = iota
	FullRequest                   RuleVariable = iota
	FullRequestLength             RuleVariable = iota
	InboundDataError              RuleVariable = iota
	MatchedVar                    RuleVariable = iota
	MatchedVarName                RuleVariable = iota
	MultipartBoundaryQuoted       RuleVariable = iota
	MultipartBoundaryWhitespace   RuleVariable = iota
	MultipartCrlfLfLines          RuleVariable = iota
	MultipartDataAfter            RuleVariable = iota
	MultipartDataBefore           RuleVariable = iota
	MultipartFileLimitExceeded    RuleVariable = iota
	MultipartHeaderFolding        RuleVariable = iota
	MultipartInvalidHeaderFolding RuleVariable = iota
	MultipartInvalidPart          RuleVariable = iota
	MultipartInvalidQuoting       RuleVariable = iota
	MultipartLfLine               RuleVariable = iota
	MultipartMissingSemicolon     RuleVariable = iota
	MultipartStrictError          RuleVariable = iota
	MultipartUnmatchedBoundary    RuleVariable = iota
	OutboundDataError             RuleVariable = iota
	PathInfo                      RuleVariable = iota
	QueryString                   RuleVariable = iota
	RemoteAddr                    RuleVariable = iota
	RemoteHost                    RuleVariable = iota
	RemotePort                    RuleVariable = iota
	ReqbodyError                  RuleVariable = iota
	ReqbodyErrorMsg               RuleVariable = iota
	ReqbodyProcessorError         RuleVariable = iota
	ReqbodyProcessorErrorMsg      RuleVariable = iota
	ReqbodyProcessor              RuleVariable = iota
	RequestBasename               RuleVariable = iota
	RequestBody                   RuleVariable = iota
	RequestBodyLength             RuleVariable = iota
	RequestFilename               RuleVariable = iota
	RequestLine                   RuleVariable = iota
	RequestMethod                 RuleVariable = iota
	RequestProtocol               RuleVariable = iota
	RequestURI                    RuleVariable = iota
	RequestURIRaw                 RuleVariable = iota
	ResponseBody                  RuleVariable = iota
	ResponseContentLength         RuleVariable = iota
	ResponseProtocol              RuleVariable = iota
	ResponseStatus                RuleVariable = iota
	ServerAddr                    RuleVariable = iota
	ServerName                    RuleVariable = iota
	ServerPort                    RuleVariable = iota
	Sessionid                     RuleVariable = iota
	HighestSeverity               RuleVariable = iota
	StatusLine                    RuleVariable = iota
	InboundErrorData              RuleVariable = iota
	Duration                      RuleVariable = iota

	ResponseHeadersNames RuleVariable = iota
	RequestHeadersNames  RuleVariable = iota
	Userid               RuleVariable = iota
	Args                 RuleVariable = iota
	ArgsGet              RuleVariable = iota
	ArgsPost             RuleVariable = iota
	FilesSizes           RuleVariable = iota
	FilesNames           RuleVariable = iota
	FilesTmpContent      RuleVariable = iota
	MultipartFilename    RuleVariable = iota
	MultipartName        RuleVariable = iota
	MatchedVarsNames     RuleVariable = iota
	MatchedVars          RuleVariable = iota
	Files                RuleVariable = iota
	RequestCookies       RuleVariable = iota
	RequestHeaders       RuleVariable = iota
	ResponseHeaders      RuleVariable = iota
	Geo                  RuleVariable = iota
	RequestCookiesNames  RuleVariable = iota
	FilesTmpnames        RuleVariable = iota
	ArgsNames            RuleVariable = iota
	ArgsGetNames         RuleVariable = iota
	ArgsPostNames        RuleVariable = iota
	TX                   RuleVariable = iota
	Rule                 RuleVariable = iota
	XML                  RuleVariable = iota
	JSON                 RuleVariable = iota
	Env                  RuleVariable = iota

	// Persisten storage kepy for compatibility
	IP RuleVariable = iota

	UrlencodedError RuleVariable = iota
)

var rulemap = map[RuleVariable]string{
	Unknown:                       "UNKNOWN",
	UrlencodedError:               "URLENCODED_ERROR",
	ResponseContentType:           "RESPONSE_CONTENT_TYPE",
	UniqueID:                      "UNIQUE_ID",
	ArgsCombinedSize:              "ARGS_COMBINED_SIZE",
	AuthType:                      "AUTH_TYPE",
	FilesCombinedSize:             "FILES_COMBINED_SIZE",
	FullRequest:                   "FULL_REQUEST",
	FullRequestLength:             "FULL_REQUEST_LENGTH",
	InboundDataError:              "INBOUND_DATA_ERROR",
	MatchedVar:                    "MATCHED_VAR",
	MatchedVarName:                "MATCHED_VAR_NAME",
	MultipartBoundaryQuoted:       "MULTIPART_BOUNDARY_QUOTED",
	MultipartBoundaryWhitespace:   "MULTIPART_BOUNDARY_WHITESPACE",
	MultipartCrlfLfLines:          "MULTIPART_CRLF_LF_LINES",
	MultipartDataAfter:            "MULTIPART_DATA_AFTER",
	MultipartDataBefore:           "MULTIPART_DATA_BEFORE",
	MultipartFileLimitExceeded:    "MULTIPART_FILE_LIMIT_EXCEEDED",
	MultipartHeaderFolding:        "MULTIPART_HEADER_FOLDING",
	MultipartInvalidHeaderFolding: "MULTIPART_INVALID_HEADER_FOLDING",
	MultipartInvalidPart:          "MULTIPART_INVALID_PART",
	MultipartInvalidQuoting:       "MULTIPART_INVALID_QUOTING",
	MultipartLfLine:               "MULTIPART_LF_LINE",
	MultipartMissingSemicolon:     "MULTIPART_MISSING_SEMICOLON",
	MultipartStrictError:          "MULTIPART_STRICT_ERROR",
	MultipartUnmatchedBoundary:    "MULTIPART_UNMATCHED_BOUNDARY",
	OutboundDataError:             "OUTBOUND_DATA_ERROR",
	PathInfo:                      "PATH_INFO",
	QueryString:                   "QUERY_STRING",
	RemoteAddr:                    "REMOTE_ADDR",
	RemoteHost:                    "REMOTE_HOST",
	RemotePort:                    "REMOTE_PORT",
	ReqbodyError:                  "REQBODY_ERROR",
	ReqbodyErrorMsg:               "REQBODY_ERROR_MSG",
	ReqbodyProcessorError:         "REQBODY_PROCESSOR_ERROR",
	ReqbodyProcessorErrorMsg:      "REQBODY_PROCESSOR_ERROR_MSG",
	ReqbodyProcessor:              "REQBODY_PROCESSOR",
	RequestBasename:               "REQUEST_BASENAME",
	RequestBody:                   "REQUEST_BODY",
	RequestBodyLength:             "REQUEST_BODY_LENGTH",
	RequestFilename:               "REQUEST_FILENAME",
	RequestLine:                   "REQUEST_LINE",
	RequestMethod:                 "REQUEST_METHOD",
	RequestProtocol:               "REQUEST_PROTOCOL",
	RequestURI:                    "REQUEST_URI",
	RequestURIRaw:                 "REQUEST_URI_RAW",
	ResponseBody:                  "RESPONSE_BODY",
	ResponseContentLength:         "RESPONSE_CONTENT_LENGTH",
	ResponseProtocol:              "RESPONSE_PROTOCOL",
	ResponseStatus:                "RESPONSE_STATUS",
	ServerAddr:                    "SERVER_ADDR",
	ServerName:                    "SERVER_NAME",
	ServerPort:                    "SERVER_PORT",
	Sessionid:                     "SESSIONID",
	HighestSeverity:               "HIGHEST_SEVERITY",
	StatusLine:                    "STATUS_LINE",
	InboundErrorData:              "INBOUND_ERROR_DATA",
	Duration:                      "DURATION",
	ResponseHeadersNames:          "RESPONSE_HEADERS_NAMES",
	RequestHeadersNames:           "REQUEST_HEADERS_NAMES",
	Userid:                        "USERID",
	Args:                          "ARGS",
	ArgsGet:                       "ARGS_GET",
	ArgsPost:                      "ARGS_POST",
	FilesSizes:                    "FILES_SIZES",
	FilesNames:                    "FILES_NAMES",
	FilesTmpContent:               "FILES_TMP_CONTENT",
	MultipartFilename:             "MULTIPART_FILENAME",
	MultipartName:                 "MULTIPART_NAME",
	MatchedVarsNames:              "MATCHED_VARS_NAMES",
	MatchedVars:                   "MATCHED_VARS",
	Files:                         "FILES",
	RequestCookies:                "REQUEST_COOKIES",
	RequestHeaders:                "REQUEST_HEADERS",
	ResponseHeaders:               "RESPONSE_HEADERS",
	Geo:                           "GEO",
	RequestCookiesNames:           "REQUEST_COOKIES_NAMES",
	FilesTmpnames:                 "FILES_TMPNAMES",
	ArgsNames:                     "ARGS_NAMES",
	ArgsGetNames:                  "ARGS_GET_NAMES",
	ArgsPostNames:                 "ARGS_POST_NAMES",
	TX:                            "TX",
	Rule:                          "RULE",
	XML:                           "XML",
	JSON:                          "JSON",
	Env:                           "ENV",
	IP:                            "IP",
}

var rulemapRev = map[string]RuleVariable{}

// Name transforms a VARIABLE representation
// into a string, it's used for audit and logging
func (v RuleVariable) Name() string {
	if name, ok := rulemap[v]; ok {
		return name
	}
	return "UNKNOWN"
}

// ParseRuleVariable returns the byte interpretation
// of a variable from a string
// Returns error if there is no representation
func ParseVariable(v string) (RuleVariable, error) {
	if v, ok := rulemapRev[strings.ToUpper(v)]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("unknown variable %s", v)
}

func init() {
	// we fill the rulemapRev with the reverse of rulemap
	for k, v := range rulemap {
		rulemapRev[v] = k
	}
}