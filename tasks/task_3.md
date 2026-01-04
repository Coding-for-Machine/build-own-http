In this stage, your server will respond to an HTTP request with a 200 response.

HTTP response
An HTTP response is made up of three parts, each separated by a CRLF (\r\n):

Status line.
Zero or more headers, each ending with a CRLF.
Optional response body.
In this stage, your server's response will only contain a status line. Here's the response your server must send:

```bash
HTTP/1.1 200 OK\r\n\r\n
```
Here's a breakdown of the response:
```bash
// Status line
HTTP/1.1  // HTTP version
200       // Status code
OK        // Optional reason phrase
\r\n      // CRLF that marks the end of the status line

// Headers (empty)
\r\n      // CRLF that marks the end of the headers

// Response body (empty)
```
For more information about HTTP responses, see the MDN Web Docs on HTTP responses or the HTTP/1.1 specification.

