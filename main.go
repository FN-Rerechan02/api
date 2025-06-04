package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	logFilePath    = "/var/log/api_go.log" // Ubah nama file log agar tidak bentrok
	tokenFilePath  = "/etc/api/key"
	scriptBasePath = "/usr/local/sbin/api/"
	serverPort     = "9001" // Port berbeda agar tidak bentrok jika Python masih jalan
)

var (
	validTokens []string
	logger      *log.Logger
)

// Struct untuk respons error JSON
type ErrorResponse struct {
	Error string `json:"error"`
}

// Struct untuk respons pesan JSON
type MessageResponse struct {
	Message string `json:"message"`
}

func init() {
	// Setup direktori log jika belum ada
	logDir := filepath.Dir(logFilePath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	// Setup file log
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	//defer logFile.Close(); // Ditutup saat program berakhir, tapi logger akan tetap terbuka

	// Logger akan menulis ke file dan stdout
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(multiWriter, "API_SERVER: ", log.LstdFlags)

	// Baca token dari file
	file, err := os.Open(tokenFilePath)
	if err != nil {
		logger.Fatalf("Failed to open token file '%s': %v", tokenFilePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		if token != "" {
			validTokens = append(validTokens, token)
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Fatalf("Failed to read tokens: %v", err)
	}
	if len(validTokens) == 0 {
		logger.Warning("No valid tokens loaded.")
	} else {
		logger.Printf("Loaded %d token(s).", len(validTokens))
	}
}

func logRequestInfo(r *http.Request, additionalInfo string) {
	clientIP := r.RemoteAddr
	// Upaya untuk mendapatkan IP asli jika di belakang proxy
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		clientIP = forwardedFor
	}
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" {
		userAgent = "User-Agent not provided"
	}
	logger.Printf("Access from IP: %s, User-Agent: %s, Path: %s, Method: %s, %s",
		clientIP, userAgent, r.URL.Path, r.Method, additionalInfo)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Printf("Error encoding JSON response: %v", err)
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.Header().Set("WWW-Authenticate", `Bearer realm="Authentication required"`)
			sendJSONResponse(w, http.StatusUnauthorized, MessageResponse{Message: "Unauthorized: Missing or invalid Bearer token"})
			logRequestInfo(r, "Unauthorized access attempt: Missing or malformed Bearer token")
			logger.Warning("Unauthorized access attempt: Missing or malformed Bearer token")
			return
		}

		providedToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		isValid := false
		for _, token := range validTokens {
			if providedToken == token {
				isValid = true
				break
			}
		}

		if !isValid {
			w.Header().Set("WWW-Authenticate", `Bearer realm="Authentication required"`)
			sendJSONResponse(w, http.StatusUnauthorized, MessageResponse{Message: "Unauthorized: Invalid Bearer token"})
			logRequestInfo(r, "Unauthorized access attempt: Invalid token")
			logger.Warning("Unauthorized access attempt: Invalid token")
			return
		}
		next.ServeHTTP(w, r)
	}
}

func executeScript(w http.ResponseWriter, r *http.Request) {
	scriptName := strings.TrimPrefix(r.URL.Path, "/")
	scriptPath := filepath.Join(scriptBasePath, scriptName)

	// Periksa apakah skrip ada dan merupakan file
	fileInfo, err := os.Stat(scriptPath)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		sendJSONResponse(w, http.StatusNotFound, MessageResponse{Message: "Script not found"})
		logRequestInfo(r, fmt.Sprintf("Script not found: %s", scriptPath))
		logger.Errorf("Script not found: %s", scriptPath) // Menggunakan Errorf jika logger mendukungnya (mis. logrus) atau Printf
		return
	}

	var cmd *exec.Cmd
	var postData []byte

	// Hanya baca body untuk metode yang relevan
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" || r.Method == "DELETE" {
		var readErr error
		postData, readErr = io.ReadAll(r.Body)
		if readErr != nil {
			sendJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to read request body"})
			logRequestInfo(r, fmt.Sprintf("Error reading request body for script %s: %v", scriptPath, readErr))
			logger.Errorf("Error reading request body for script %s: %v", scriptPath, readErr)
			return
		}
		defer r.Body.Close()
		cmd = exec.Command(scriptPath)
		cmd.Stdin = bytes.NewReader(postData)
	} else {
		cmd = exec.Command(scriptPath)
	}

	startTime := time.Now()
	output, err := cmd.CombinedOutput() // CombinedOutput menangkap stdout dan stderr
	duration := time.Since(startTime)

	if err != nil {
		// Coba interpretasi error sebagai exit code non-zero
		if exitError, ok := err.(*exec.ExitError); ok {
			 errorMsg := fmt.Sprintf("Script execution failed with exit code %d. Output: %s", exitError.ExitCode(), string(output))
			sendJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: errorMsg})
			logRequestInfo(r, fmt.Sprintf("Error executing script: %s, ExitCode: %d, Duration: %s, Error: %s, Output: %s", scriptPath, exitError.ExitCode(), duration, err.Error(), string(output)))
			logger.Errorf("Error executing script: %s, ExitCode: %d, Duration: %s, Error: %s, Output: %s", scriptPath, exitError.ExitCode(), duration, err.Error(), string(output))

		} else {
			// Error lain (mis. script tidak ditemukan oleh exec, permission denied)
			sendJSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: fmt.Sprintf("Error executing script: %v", err)})
			logRequestInfo(r, fmt.Sprintf("Error executing script: %s, Duration: %s, Error: %v", scriptPath, duration, err))
			logger.Errorf("Error executing script: %s, Duration: %s, Error: %v", scriptPath, duration, err)
		}
		return
	}

	// Sukses
	w.Header().Set("Content-Type", "application/json") // Asumsikan output skrip adalah JSON
	// Jika skrip mungkin mengeluarkan non-JSON, atur Content-Type sesuai atau biarkan kosong
	w.WriteHeader(http.StatusOK)
	w.Write(output) // Tulis output skrip langsung
	logRequestInfo(r, fmt.Sprintf("Successfully executed script: %s, Duration: %s", scriptPath, duration))
	logger.Printf("Successfully executed script: %s, Duration: %s, Output: %s", scriptPath, duration, string(output))
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET, POST, DELETE, PUT, PATCH, CONNECT, TRACE, HEAD, OPTIONS")
	sendJSONResponse(w, http.StatusOK, MessageResponse{Message: "OPTIONS request received"})
	logRequestInfo(r, "OPTIONS request handled")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Log dasar untuk setiap permintaan yang masuk ke handler utama
	logRequestInfo(r, "Incoming request")

	switch r.Method {
	case "GET", "CONNECT", "TRACE", "HEAD":
		executeScript(w, r)
	case "POST", "DELETE", "PUT", "PATCH":
		executeScript(w, r) // executeScript akan menangani pembacaan body
	case "OPTIONS":
		optionsHandler(w, r)
	default:
		sendJSONResponse(w, http.StatusMethodNotAllowed, MessageResponse{Message: "Method not allowed"})
		logRequestInfo(r, "Method not allowed")
	}
}

func main() {
	// Setup handler dengan middleware otentikasi
	// Middleware diterapkan ke semua path yang ditangani oleh mainHandler
	http.HandleFunc("/", authMiddleware(mainHandler))

	logger.Printf("Starting Go httpd server on port %s", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
