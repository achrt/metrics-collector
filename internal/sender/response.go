package sender

import "io"

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string
	Body       io.ReadCloser
}
