package middleware

import (
	"log"
	"net/http"
	"time"
)

type (
	// struct for holding response details
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter // compose original http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	r.responseData.size += size            // capture size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	r.responseData.status = statusCode       // capture status code
}

func ResponseLoggerMid(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lrw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		startTime := time.Now()
		handler.ServeHTTP(&lrw, r)
		endTime := time.Now()

		duration := endTime.Sub(startTime)

		log.Printf("Request: %s %s %s %s %d %d %s",
			startTime.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			r.Proto,
			responseData.status,
			responseData.size,
			duration)
	})
}
