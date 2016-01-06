// Copyright (c) 2015-2016 Kerem Güneş
//   <http://qeremy.com>
//
// GNU General Public License v3.0
//   <http://www.gnu.org/licenses/gpl-3.0.txt>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// @package    ohttp
// @subpackage ohttp.response.status
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package status

// Informational constants.
// @const int
const (
    CONTINUE                               = 100
    SWITCHING_PROTOCOLS                    = 101
    PROCESSING                             = 102
)

// Success constants.
// @const int
const (
    OK                                     = 200
    CREATED                                = 201
    ACCEPTED                               = 202
    NON_AUTHORITATIVE_INFORMATION          = 203
    NO_CONTENT                             = 204
    RESET_CONTENT                          = 205
    PARTIAL_CONTENT                        = 206
    MULTI_STATUS                           = 207
    ALREADY_REPORTED                       = 208
    IM_USED                                = 226
)

// Redirection constants.
// @const int
const (
    MULTIPLE_CHOICES                       = 300
    MOVED_PERMANENTLY                      = 301
    FOUND                                  = 302
    SEE_OTHER                              = 303
    NOT_MODIFIED                           = 304
    USE_PROXY                              = 305
    SWITCH_PROXY                           = 306
    TEMPORARY_REDIRECT                     = 307
    PERMANENT_REDIRECT                     = 308
)

// Client error constants.
// @const int
const (
    BAD_REQUEST                            = 400
    UNAUTHORIZED                           = 401
    PAYMENT_REQUIRED                       = 402
    FORBIDDEN                              = 403
    NOT_FOUND                              = 404
    METHOD_NOT_ALLOWED                     = 405
    NOT_ACCEPTABLE                         = 406
    PROXY_AUTHENTICATION_REQUIRED          = 407
    REQUEST_TIMEOUT                        = 408
    CONFLICT                               = 409
    GONE                                   = 410
    LENGTH_REQUIRED                        = 411
    PRECONDITION_FAILED                    = 412
    PAYLOAD_TOO_LARGE                      = 413
    URI_TOO_LONG                           = 414
    UNSUPPORTED_MEDIA_TYPE                 = 415
    RANGE_NOT_SATISFIABLE                  = 416
    EXPECTATION_FAILED                     = 417
    I_M_A_TEAPOT                           = 418
    AUTHENTICATION_TIMEOUT                 = 419
    ENHANCE_YOUR_CALM                      = 420
    MISDIRECTED_REQUEST                    = 421
    UNPROCESSABLE_ENTITY                   = 422
    LOCKED                                 = 423
    FAILED_DEPENDENCY                      = 424
    UPGRADE_REQUIRED                       = 426
    PRECONDITION_REQUIRED                  = 428
    TOO_MANY_REQUESTS                      = 429
    REQUEST_HEADER_FIELDS_TOO_LARGE        = 431
    LOGIN_TIMEOUT                          = 440
    NO_RESPONSE                            = 444
    RETRY_WITH                             = 449
    BLOCKED_BY_WINDOWS_PARENTAL_CONTROLS   = 450
    UNAVAILABLE_FOR_LEGAL_REASONS          = 451
    REQUEST_HEADER_TOO_LARGE               = 494
    CERT_ERROR                             = 495
    NO_CERT                                = 496
    HTTP_TO_HTTPS                          = 497
    TOKEN_EXPIRED_INVALID                  = 498
    CLIENT_CLOSED_REQUEST                  = 499
)

// Server error constants.
// @const int
const (
    INTERNAL_SERVER_ERROR                  = 500
    NOT_IMPLEMENTED                        = 501
    BAD_GATEWAY                            = 502
    SERVICE_UNAVAILABLE                    = 503
    GATEWAY_TIMEOUT                        = 504
    HTTP_VERSION_NOT_SUPPORTED             = 505
    VARIANT_ALSO_NEGOTIATES                = 506
    INSUFFICIENT_STORAGE                   = 507
    LOOP_DETECTED                          = 508
    BANDWIDTH_LIMIT_EXCEEDED               = 509
    NOT_EXTENDED                           = 510
    NETWORK_AUTHENTICATION_REQUIRED        = 511
    UNKNOWN_ERROR                          = 520
    ORIGIN_CONNECTION_TIME_OUT             = 522
    NETWORK_READ_TIMEOUT_ERROR             = 598
    NETWORK_CONNECT_TIMEOUT_ERROR          = 599
)

// Statuses.
// @var map[int]string
var statuses = map[int]string{
    // informationals
    CONTINUE                               : "Continue",
    SWITCHING_PROTOCOLS                    : "Switching Protocols",
    PROCESSING                             : "Processing",
    // success
    OK                                     : "OK",
    CREATED                                : "Created",
    ACCEPTED                               : "Accepted",
    NON_AUTHORITATIVE_INFORMATION          : "Non-Authoritative Information",
    NO_CONTENT                             : "No Content",
    RESET_CONTENT                          : "Reset Content",
    PARTIAL_CONTENT                        : "Partial Content",
    MULTI_STATUS                           : "Multi-Status",
    ALREADY_REPORTED                       : "Already Reported",
    IM_USED                                : "IM Used",
    // redirection
    MULTIPLE_CHOICES                       : "Multiple Choices",
    MOVED_PERMANENTLY                      : "Moved Permanently",
    FOUND                                  : "Found",
    SEE_OTHER                              : "See Other",
    NOT_MODIFIED                           : "Not Modified",
    USE_PROXY                              : "Use Proxy",
    SWITCH_PROXY                           : "Switch Proxy",
    TEMPORARY_REDIRECT                     : "Temporary Redirect",
    PERMANENT_REDIRECT                     : "Permanent Redirect",
    // client error
    BAD_REQUEST                            : "Bad Request",
    UNAUTHORIZED                           : "Unauthorized",
    PAYMENT_REQUIRED                       : "Payment Required",
    FORBIDDEN                              : "Forbidden",
    NOT_FOUND                              : "Not Found",
    METHOD_NOT_ALLOWED                     : "Method Not Allowed",
    NOT_ACCEPTABLE                         : "Not Acceptable",
    PROXY_AUTHENTICATION_REQUIRED          : "Proxy Authentication Required",
    REQUEST_TIMEOUT                        : "Request Timeout",
    CONFLICT                               : "Conflict",
    GONE                                   : "Gone",
    LENGTH_REQUIRED                        : "Length Required",
    PRECONDITION_FAILED                    : "Precondition Failed",
    PAYLOAD_TOO_LARGE                      : "Payload Too Large",
    URI_TOO_LONG                           : "URI Too Long",
    UNSUPPORTED_MEDIA_TYPE                 : "Unsupported Media Type",
    RANGE_NOT_SATISFIABLE                  : "Range Not Satisfiable",
    EXPECTATION_FAILED                     : "Expectation Failed",
    I_M_A_TEAPOT                           : "I'm a teapot",
    AUTHENTICATION_TIMEOUT                 : "Authentication Timeout",
    ENHANCE_YOUR_CALM                      : "Enhance Your Calm",
    MISDIRECTED_REQUEST                    : "Misdirected Request",
    UNPROCESSABLE_ENTITY                   : "Unprocessable Entity",
    LOCKED                                 : "Locked",
    FAILED_DEPENDENCY                      : "Failed Dependency",
    UPGRADE_REQUIRED                       : "Upgrade Required",
    PRECONDITION_REQUIRED                  : "Precondition Required",
    TOO_MANY_REQUESTS                      : "Too Many Requests",
    REQUEST_HEADER_FIELDS_TOO_LARGE        : "Request Header Fields Too Large",
    LOGIN_TIMEOUT                          : "Login Timeout",
    NO_RESPONSE                            : "No Response",
    RETRY_WITH                             : "Retry With",
    BLOCKED_BY_WINDOWS_PARENTAL_CONTROLS   : "Blocked by Windows Parental Controls",
    UNAVAILABLE_FOR_LEGAL_REASONS          : "Unavailable For Legal Reasons",
    REQUEST_HEADER_TOO_LARGE               : "Request Header Too Large",
    CERT_ERROR                             : "Cert Error",
    NO_CERT                                : "No Cert",
    HTTP_TO_HTTPS                          : "HTTP to HTTPS",
    TOKEN_EXPIRED_INVALID                  : "Token Expired/Invalid",
    CLIENT_CLOSED_REQUEST                  : "Client Closed Request",
    // server error
    INTERNAL_SERVER_ERROR                  : "Internal Server Error",
    NOT_IMPLEMENTED                        : "Not Implemented",
    BAD_GATEWAY                            : "Bad Gateway",
    SERVICE_UNAVAILABLE                    : "Service Unavailable",
    GATEWAY_TIMEOUT                        : "Gateway Timeout",
    HTTP_VERSION_NOT_SUPPORTED             : "HTTP Version Not Supported",
    VARIANT_ALSO_NEGOTIATES                : "Variant Also Negotiates",
    INSUFFICIENT_STORAGE                   : "Insufficient Storage",
    LOOP_DETECTED                          : "Loop Detected",
    BANDWIDTH_LIMIT_EXCEEDED               : "Bandwidth Limit Exceeded",
    NOT_EXTENDED                           : "Not Extended",
    NETWORK_AUTHENTICATION_REQUIRED        : "Network Authentication Required",
    UNKNOWN_ERROR                          : "Unknown Error",
    ORIGIN_CONNECTION_TIME_OUT             : "Origin Connection Time-out",
    NETWORK_READ_TIMEOUT_ERROR             : "Network Read Timeout Error",
    NETWORK_CONNECT_TIMEOUT_ERROR          : "Network Connect Timeout Error",
}

// Get status code by text.
//
// @param  x string
// @return (int)
func Code(x string) (int) {
    for c, t := range statuses {
        if x == t {
            return c
        }
    }
    return 0
}

// Get status text by code.
//
// @param  x int
// @return (string)
func Text(x int) (string) {
    for c, t := range statuses {
        if x == c {
            return t
        }
    }
    return ""
}

// @object ohttp.response.status.Status
type Status struct {
    status       string
    code         int
    text         string // reason phrase
}

// Constructor.
//
// @param  c int
// @param  t string
// @param  s string
// @return (*ohttp.status.Status)
func New(c int, t, s string) (*Status) {
    return &Status{
        status: s,
          code: c,
          text: t,
    }
}

// Get: status line.
//
// @return (string)
func (this *Status) Status() (string) {
    return this.status
}

// Get: code.
//
// @return (int)
func (this *Status) Code() (int) {
    return this.code
}

// Get: text.
//
// @return (string)
func (this *Status) Text() (string) {
    return this.text
}

// Set: status-line.
//
// @param  s string
// @return (*ohttp.status.Status)
func (this *Status) SetStatus(s string) (*Status) {
    this.status = s
    return this
}

// Set: code.
//
// @param  c int
// @return (*ohttp.status.Status)
func (this *Status) SetCode(c int) (*Status) {
    this.code = c
    return this
}

// Set: text.
//
// @param  t string
// @return (*ohttp.status.Status)
func (this *Status) SetText(t string) (*Status) {
    this.text = t
    return this
}
