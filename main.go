package main

import (
	"fmt"
	"net"
	"strings"
)

type cfmHTTP struct {
}

func (c *cfmHTTP) New() *cfmHTTP {
	return &cfmHTTP{}
}

const PORT string = ":8181"

type Request struct {
	// start line
	Method  string
	Path    string
	Version string
	// headers
	Headers map[string]string
	// body
	body string
}

var Status = map[int16]string{
	// Axborotga oid javoblar
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	103: "Early Hints",
	// Muvaffaqiyatli javoblar
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
	// Yo'naltirish xabarlari
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	306: "unused",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	// Mijoz xatolariga javoblar
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
	413: "Content Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Content",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	428: "Precondition Required",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	// Server xatolariga javoblar
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}

type Response struct {
	// start line
	Version    string
	StatusCode map[uint16]string
	// headers
	Headers map[string]string
	// body
	body string
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Listen Error: ", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Connection error: ", err)
		}
		// http -> request va response

		// request
		buf := make([]byte, 4096) //4 kb
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client Request read error: ", err)
		}
		message := string(buf[:n])
		fmt.Println("Request: ", message)
		full_request := strings.Split(message, "\r\n")   // ["GET /user-agent HTTP/1.1", "Host: localhost:4221", "Accept: */*", ..."", {bosy}]
		StartLine := strings.Split(full_request[0], " ") // "GET /user-agent HTTP/1.1" --> ["GET", "/", "HTTP/1,1"]
		if len(StartLine) != 3 {
			fmt.Println("Bad Request")
			continue
		}
		request := Request{
			Method:  StartLine[0],
			Path:    StartLine[1],
			Version: StartLine[2],
		}
		fmt.Println("Request: ", request)
		if request.Path == "/" {
			// HTTP Response
			body := `<!DOCTYPE html>
					<html>
					<head>
						<title> CfM </title>
					</head>
					<body>
						<h1> This is TCP Server</h1>
					</body>
					</html>
					`
			headers := "Connection: close\r\n" + "Content-Type: text/html; charset=UTF-8\r\n" + fmt.Sprintf("Content-Length: %d\r\n", len(body))
			response := "HTTP/1.1 200 OK\r\n" + headers + "\r\n" + body
			conn.Write([]byte(response))
		} else {
			// HTTP Response
			body := `<!DOCTYPE html>
					<html>
						<head>
						<title> Not Found </title>
						</head>
						<body>
						<h1> Not Found 404</h1>
						</body>
					</html>
					`
			headers := "Connection: close\r\n" + "Content-Type: text/html; charset=UTF-8\r\n" + fmt.Sprintf("Content-Length: %d\r\n", len(body))
			response := "HTTP/1.1 404 Not Found\r\n" + headers + "\r\n" + body
			conn.Write([]byte(response))
		}

	}

}

// -------------------------------------------------
// Request

// Request line
// GET /user-agent HTTP/1.1\r\nHost: localhost:4221\r\nAccept: */*\r\n\r\n

// Request body (empty)
/*
example: request method
GET, POST, PUT/PATCH, DELETE
*/
// -------------------------------------------------
// Response

// Status line
// HTTP/1.1 200 OK\r\n

// Headers
// Accept-Ranges: bytes\r\n
// Age: 294510\r\n
// Cache-Control: max-age=604800\r\n
// Content-Type: text/html; charset=UTF-8\r\n
// Date: Fri, 21 Jun 2024 14:18:33 GMT\r\n
// Etag: "3147526947"\r\n
// Expires: Fri, 28 Jun 2024 14:18:33 GMT\r\n
// Last-Modified: Thu, 17 Oct 2019 07:18:26 GMT\r\n
// Server: ECAcc (nyd/D10E)\r\n
// X-Cache: HIT\r\n
// Content-Length: 1256\r\n
// \r\n

// Response body
// <!doctype html>
// <!-- HTML content follows here -->
