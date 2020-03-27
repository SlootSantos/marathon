package marathon

import "net/http"

const (
	headerContentType  = "text/event-stream"
	headerCacheControl = "no-cache"
	headerConnection   = "keep-alive"
	headerAccesControl = "*"
)

func setSSEHeaders(w http.ResponseWriter) http.ResponseWriter {
	h := w.Header()
	h.Set("Content-Type", "text/event-stream")
	h.Set("Cache-Control", "no-cache")
	h.Set("Connection", "keep-alive")
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("X-Accel-Buffering", "no")

	return w
}
