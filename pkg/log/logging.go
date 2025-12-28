package log

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imkarthi24/sf-backend/pkg/util"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w *BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *BodyLogWriter) Format() string {
	reader, err := gzip.NewReader(bytes.NewReader(w.Body.Bytes()))
	if err == nil {
		defer reader.Close()
		decompressed, err := io.ReadAll(reader)
		if err == nil {
			return string(decompressed)
		}
	}
	return w.Body.String()
}

func HandleResponseBody(w *gin.ResponseWriter) *BodyLogWriter {
	blw := &BodyLogWriter{
		ResponseWriter: *w,
		Body:           &bytes.Buffer{},
	}
	*w = blw
	return blw
}

func HandleRequestBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Sprintf("Error reading request body: %v", err)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return maskSensitiveData(string(body))
}

func maskSensitiveData(body string) string {
	// Add logic to mask passwords, tokens, etc.
	sensitiveKeys := []string{"password", "token", "secret"}

	var maskedBody = body
	for _, key := range sensitiveKeys {
		if strings.Contains(strings.ToLower(maskedBody), key) {
			maskedBody = "[SENSITIVE DATA MASKED]"
			break
		}
	}
	return maskedBody
}

func FormatRequestAndResponse(
	writer gin.ResponseWriter,
	request *http.Request,
	responseBody string,
	requestID string,
	requestBody string,
) string {
	return fmt.Sprintf(
		"CorrelationId: %s | "+
			"Method: %s | "+
			"Path: %s | "+
			"Host: %s | "+
			"Status: %d | "+
			"RequestBody: %s | "+
			"ResponseBody: %s | "+
			"Timestamp: %s",
		requestID,
		request.Method,
		request.URL.Path,
		request.Host,
		writer.Status(),
		requestBody,
		truncateBody(responseBody, 0),
		util.GetLocalTime().Format(time.RFC3339),
	)
}

func truncateBody(body string, maxLength int) string {
	if maxLength == 0 {
		maxLength = 2048 // Default max length
	}

	if len(body) > maxLength {
		return body[:maxLength] + "... [truncated]"
	}
	return body
}
