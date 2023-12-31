package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/cpu"
)

// Define the Pushbullet API endpoint and your API token here.
var checkIntervalDuration time.Duration
var logger *log.Logger
var apiKey string

var apiURL string
var threshold float64
var averageLen int

var enableConsole bool
var enableTestNotif bool
var timespanAverage float64
var checkInterval float64

var notifInterval int
var thresholdDuration int

// var enableConsole = true
// var checkInterval = 1.0
// var timespanAverage = 1.0
// var enableTestNotif = true
// var threshold = 80.0
// var apiURL = "https://api.pushbullet.com/v2/pushes"

// apiURL := "ht#tps://api.pushbullet.com/v2/pushes"
// threshold := "80.0"
// averag#eLen := "1"
// enableConsole
// enableTestNotif

func checkCPULoad() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func pushAlert(cpuLoad float64) {
	// Create a data structure for your Pushbullet request.
	timeString := time.Now().Format("2006-01-02 15:04:05")
	data := map[string]interface{}{
		"type":  "note",
		"title": "CPU Alert",
		"body":  fmt.Sprintf("%s \nHigh CPU load detected: %.2f%%", timeString, cpuLoad), //"This is a test message sent via Pushbullet API.",
	}

	// Convert the data to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Create an HTTP request.
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers, including your API key.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", apiKey)

	// Create an HTTP client and send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status.
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Message sent successfully!")
	} else {
		fmt.Println("Error sending message. Status code:", resp.Status)
	}
}

func setEnvs() {
	var err error

	apiURLStr := os.Getenv("PUSHBULLET_ENDPOINT_URL")
	if apiURLStr != "" {
		apiURL = apiURLStr
	}

	apiKey = os.Getenv("PUSHBULLET_API_KEY")
	if apiKey == "" {
		fmt.Println("NO PUSHBULLET_API_KEY given!")
		log.Fatalf("NO PUSHBULLET_API_KEY given!")
		os.Exit(1)
	} else if apiKey == "SKIP" {
		fmt.Println("NO PUSHBULLET_API_KEY given!")
		logger.Println("NO PUSHBULLET_API_KEY given!")
		fmt.Println("PUSHBULLET API SKIPPED")
		logger.Println("PUSHBULLET API SKIPPED")
		apiURL = "https://api.pushbullet.com/v2/pushes"
		threshold = 80.0
		averageLen = 1.0
		enableConsole = true
		enableTestNotif = true
	}

	thresholdStr := os.Getenv("CPU_AVERAGE_MAX_THRESHOLD")
	if thresholdStr != "" {

		// Convert the string to an integer
		threshold, err = strconv.ParseFloat(thresholdStr, 64)
		if err != nil {
			log.Fatalf("Error converting CPU_AVERAGE_MAX_THRESHOLD to float: %v", err)
		}
	}

	checkIntervalStr := os.Getenv("CHECK_INTERVAL_SECONDS")

	if checkIntervalStr != "" {
		// Convert the string to an integer
		checkInterval, err = strconv.ParseFloat(checkIntervalStr, 64)

		if err != nil {
			log.Fatalf("Error converting CHECK_INTERVAL_SECONDS to float: %v", err)
		}
	}

	timespanAvStr := os.Getenv("TIMESPAN_AVERAGE_MINUTES")
	if timespanAvStr != "" {
		// Convert the string to an integer
		timespanAverage, err = strconv.ParseFloat(timespanAvStr, 64)
		if err != nil {
			log.Fatalf("Error converting TIMESPAN_AVERAGE_MINUTES to float: %v", err)
		}
	}

	// Read the ENABLE_CONSOLE_OUTPUT variable from the environment
	enableConsoleOutputStr := os.Getenv("ENABLE_CONSOLE_OUTPUT")
	if enableConsoleOutputStr != "" {
		// Convert the string to a boolean
		enableConsole, err = strconv.ParseBool(enableConsoleOutputStr)
		if err != nil {
			log.Fatalf("Error parsing ENABLE_CONSOLE_OUTPUT: %v", err)
		}
	}

	checkIntervalDuration = time.Second * time.Duration(checkInterval) //checkInterval//time.Minute

	// Calculate the amount of measured points needed for the given time settings
	avergageValsFloat := timespanAverage / checkInterval * 60
	averageLen = int(avergageValsFloat)

	// Read the SEND_TEST_NOTIFICATION_ON_LAUNCH variable from the environment
	enableTestNotifStr := os.Getenv("SEND_TEST_NOTIFICATION_ON_LAUNCH")
	if enableTestNotifStr != "" {
		// Convert the string to a boolean
		enableTestNotif, err = strconv.ParseBool(enableTestNotifStr)
		if err != nil {
			log.Fatalf("Error parsing SEND_TEST_NOTIFICATION_ON_LAUNCH: %v", err)
		}
	}

	notifIntervalStr := os.Getenv("NOTIFICATION_INTERVAL_MINUTES")
	if notifIntervalStr != "" {
		notifInterval, err = strconv.Atoi(notifIntervalStr)
		if err != nil {
			log.Fatalf("Error parsing NOTIFICATION_INTERVAL_MINUTES: %v", err)
		}
	}

	thresholdDurationStr := os.Getenv("THRESHOLD_DURATION_ALARM_MINUTES")
	if thresholdDurationStr != "" {
		thresholdDuration, err = strconv.Atoi(thresholdDurationStr)
		if err != nil {
			log.Fatalf("Error parsing THRESHOLD_DURATION_ALARM_MINUTES: %v", err)
		}
	}

	if enableTestNotif {
		pushAlert(420.69)
	}
}

func pushArray(item float64, arr []float64) []float64 {

	if len(arr) < averageLen {
		return append([]float64{item}, arr...)
	} else {

		return append([]float64{item}, arr[:len(arr)-1]...)
	}
}

func averageArray(arr []float64) float64 {
	var sum float64 = 0
	for _, i := range arr {
		sum += i
	}
	return sum / float64(len(arr))
}

// func timestamp() string {
// 	t := time.Now()
// 	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d : ",
// 		t.Year(), t.Month(), t.Day(),
// 		t.Hour(), t.Minute(), t.Second())
// 	return formatted
// }

func main() {

	// Open the file for appending or create it if it doesn't exist
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	defer func() {
		logger.Println("Closing Application")
		logFile.Close() // Close the file when done
	}()
	// Create a logger using the opened file
	logger = log.New(logFile, "CPU: ", log.Ldate|log.Ltime)

	// Capture Ctrl+C signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		logger.Println("Closing Application")
		logFile.Close()
		os.Exit(0)
	}()

	// Print the formatted time
	logger.Println("CPU watcher launched.")

	if os.Getenv("DISABLE_ENV_FILE") != "true" && os.Getenv("DISABLE_ENV_FILE") == "" {
		fmt.Println("Reading Environment variables from file .env")
		logger.Println("Reading Environment variables from file .env")
		if err := godotenv.Load(".env"); err != nil {
			fmt.Printf("Error loading .env file: %v", err)
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
	setEnvs()
	logger.Println("Environment variables loaded.")

	array := []float64{0}

	// array := make([]float64, 25)
	var averageCPU float64

	now := time.Now()
	lastAlertT := now.Add(-time.Duration(notifInterval) * time.Minute)

	//timestamp := now.Add(-time.Duration(notifInterval))

	thresholdCounter := 0
	thresholdCheckedT := time.Now()

	thresholdMax := int(thresholdDuration)
	// fmt.Println(time.Duration(notifInterval) * time.Minute)
	// fmt.Println(time.Now().After(lastAlertT.Add(time.Duration(notifInterval) * time.Minute)))

	for {

		cpuLoad := checkCPULoad()
		array = pushArray(cpuLoad, array)
		averageCPU = averageArray(array)
		// fmt.Println(thresholdCounter)

		// fmt.Printf("Current CPU load: %.2f%%\n", averageCPU)
		if enableConsole {
			fmt.Printf("Average CPU load: %.2f%% - momentary %.2f%%\n", averageCPU, cpuLoad)
		}

		if time.Now().After(thresholdCheckedT.Add(time.Minute)) {
			thresholdCheckedT = time.Now()
			if averageCPU > threshold {
				if enableConsole {
					fmt.Println(time.Now().Format("2006-02-13 15:04:05"))
					fmt.Printf("Warning! CPU load: %.2f%% - momentary %.2f%%\n", averageCPU, cpuLoad)
				}
				if thresholdCounter == thresholdMax {
					// fmt.Println("Max reached")
					if time.Now().After(lastAlertT.Add(time.Duration(notifInterval) * time.Minute)) {
						lastAlertT = time.Now()
						pushAlert(averageCPU)
						logger.Println("Pushed Notification")

					}
				} else {
					thresholdCounter += 1
				}
			} else {
				if thresholdCounter > 0 {
					thresholdCounter -= 1
				}
			}
		}
		// if averageCPU > threshold {
		// 	logger.Printf("WARNING average CPU load: %.2f%% - momentary %.2f%%\n", averageCPU, cpuLoad)
		// 	timestamp := time.Now().Add(- time.Minute * time.Duration(thresholdDuration))
		// 	if time.Now().After(timestamp.Add(time.Duration(notifInterval) * time.Minute)) { // timestamp.Add + time.Duration(notifInterval).Minutes() {
		// 		pushAlert(averageCPU)
		// 		fmt.Println("A:", time.Now().Second())
		// 		fmt.Println("B:", timestamp.Add(time.Duration(notifInterval)).Second())
		// 		timestamp = time.Now()
		// 		fmt.Println("T:", timestamp)
		// 	}

		// }

		time.Sleep(checkIntervalDuration)
	}
}
