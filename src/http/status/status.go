package status

// Informational constants.
// @const int
const (
    CONTINUE                             = 100
    SWITCHING_PROTOCOLS                  = 101
    PROCESSING                           = 102
)

// Success constants.
// @const int
const (
    OK                                   = 200
    CREATED                              = 201
    ACCEPTED                             = 202
    NON_AUTHORITATIVE_INFORMATION        = 203
    NO_CONTENT                           = 204
    RESET_CONTENT                        = 205
    PARTIAL_CONTENT                      = 206
    MULTI_STATUS                         = 207
    ALREADY_REPORTED                     = 208
    IM_USED                              = 226
)

// Redirection constants.
// @const int
const (
    MULTIPLE_CHOICES                     = 300
    MOVED_PERMANENTLY                    = 301
    FOUND                                = 302
    SEE_OTHER                            = 303
    NOT_MODIFIED                         = 304
    USE_PROXY                            = 305
    SWITCH_PROXY                         = 306
    TEMPORARY_REDIRECT                   = 307
    PERMANENT_REDIRECT                   = 308
    RESUME_INCOMPLETE                    = 308
)

// Client error constants.
// @const int
const (
    BAD_REQUEST                          = 400
    UNAUTHORIZED                         = 401
    PAYMENT_REQUIRED                     = 402
    FORBIDDEN                            = 403
    NOT_FOUND                            = 404
    METHOD_NOT_ALLOWED                   = 405
    NOT_ACCEPTABLE                       = 406
    PROXY_AUTHENTICATION_REQUIRED        = 407
    REQUEST_TIMEOUT                      = 408
    CONFLICT                             = 409
    GONE                                 = 410
    LENGTH_REQUIRED                      = 411
    PRECONDITION_FAILED                  = 412
    PAYLOAD_TOO_LARGE                    = 413
    URI_TOO_LONG                         = 414
    UNSUPPORTED_MEDIA_TYPE               = 415
    RANGE_NOT_SATISFIABLE                = 416
    EXPECTATION_FAILED                   = 417
    I_M_A_TEAPOT                         = 418
    AUTHENTICATION_TIMEOUT               = 419
    METHOD_FAILURE                       = 420
    ENHANCE_YOUR_CALM                    = 420
    MISDIRECTED_REQUEST                  = 421
    UNPROCESSABLE_ENTITY                 = 422
    LOCKED                               = 423
    FAILED_DEPENDENCY                    = 424
    UPGRADE_REQUIRED                     = 426
    PRECONDITION_REQUIRED                = 428
    TOO_MANY_REQUESTS                    = 429
    REQUEST_HEADER_FIELDS_TOO_LARGE      = 431
    LOGIN_TIMEOUT                        = 440
    NO_RESPONSE                          = 444
    RETRY_WITH                           = 449
    BLOCKED_BY_WINDOWS_PARENTAL_CONTROLS = 450
    UNAVAILABLE_FOR_LEGAL_REASONS        = 451
    REDIRECT                             = 451
    REQUEST_HEADER_TOO_LARGE             = 494
    CERT_ERROR                           = 495
    NO_CERT                              = 496
    HTTP_TO_HTTPS                        = 497
    TOKEN_EXPIRED_INVALID                = 498
    CLIENT_CLOSED_REQUEST                = 499
    TOKEN_REQUIRED                       = 499
)

// Server error constants.
// @const int
const (
    INTERNAL_SERVER_ERROR                = 500
    NOT_IMPLEMENTED                      = 501
    BAD_GATEWAY                          = 502
    SERVICE_UNAVAILABLE                  = 503
    GATEWAY_TIMEOUT                      = 504
    HTTP_VERSION_NOT_SUPPORTED           = 505
    VARIANT_ALSO_NEGOTIATES              = 506
    INSUFFICIENT_STORAGE                 = 507
    LOOP_DETECTED                        = 508
    BANDWIDTH_LIMIT_EXCEEDED             = 509
    NOT_EXTENDED                         = 510
    NETWORK_AUTHENTICATION_REQUIRED      = 511
    UNKNOWN_ERROR                        = 520
    ORIGIN_CONNECTION_TIME_OUT           = 522
    NETWORK_READ_TIMEOUT_ERROR           = 598
    NETWORK_CONNECT_TIMEOUT_ERROR        = 599
)

// Statuses.
// @var map[int]string
var _statuses = map[int]string{
    // informationals
    100: "Continue",
    101: "Switching Protocols",
    102: "Processing",
    // success
    200: "OK",
    201: "Created",
    202: "Accepted",
    203: "Non-Authoritative Information",
    204: "No Content",
    205: "Reset Content",
    206: "Partial Content",
    207: "Multi-Status",
    208: "Already Reported",
    226: "IM Used",
    // redirection
    300: "Multiple Choices",
    301: "Moved Permanently",
    302: "Found",
    303: "See Other",
    304: "Not Modified",
    305: "Use Proxy",
    306: "Switch Proxy",
    307: "Temporary Redirect",
    308: "Permanent Redirect",
    // client error
    400: "Bad Request",
    401: "Unauthorized",
    402: "Payment Required",
    403: "Forbidden",
    404: "Not Found",
    405: "Method Not Allowed",
    406: "Not Acceptable",
    407: "Proxy Authentication Required",
    408: "Request Timeout",
    409: "Conflict",
    410: "Gone",
    411: "Length Required",
    412: "Precondition Failed",
    413: "Payload Too Large",
    414: "URI Too Long",
    415: "Unsupported Media Type",
    416: "Range Not Satisfiable",
    417: "Expectation Failed",
    418: "I'm a teapot",
    419: "Authentication Timeout",
    420: "Enhance Your Calm",
    421: "Misdirected Request",
    422: "Unprocessable Entity",
    423: "Locked",
    424: "Failed Dependency",
    426: "Upgrade Required",
    428: "Precondition Required",
    429: "Too Many Requests",
    431: "Request Header Fields Too Large",
    440: "Login Timeout",
    444: "No Response",
    449: "Retry With",
    450: "Blocked by Windows Parental Controls",
    451: "Unavailable For Legal Reasons",
    494: "Request Header Too Large",
    495: "Cert Error",
    496: "No Cert",
    497: "HTTP to HTTPS",
    498: "Token Expired/Invalid",
    499: "Client Closed Request",
    // server error
    500: "Internal Server Error",
    501: "Not Implemented",
    502: "Bad Gateway",
    503: "Service Unavailable",
    504: "Gateway Timeout",
    505: "HTTP Version Not Supported",
    506: "Variant Also Negotiates",
    507: "Insufficient Storage",
    508: "Loop Detected",
    509: "Bandwidth Limit Exceeded",
    510: "Not Extended",
    511: "Network Authentication Required",
    520: "Unknown Error",
    522: "Origin Connection Time-out",
    598: "Network Read Timeout Error",
    599: "Network Connect Timeout Error",
}

// Get status code by text.
//
// @param  x string
// @return int
func Code(x string) (int) {
    for code, text := range _statuses {
        if x == text {
            return code
        }
    }
    return 0
}

// Get status text by code.
//
// @param  x int
// @return string
func Text(x int) (string) {
    for code, text := range _statuses {
        if x == code {
            return text
        }
    }
    return ""
}
