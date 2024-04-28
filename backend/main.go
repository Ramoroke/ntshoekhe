package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var (
	filePath = "cluster.txt" // Path to the text file containing the list of ports
	mutex    sync.Mutex      // Mutex for concurrent access to the map
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: programname command [args...]")
		os.Exit(1)
	}

	go monitorPorts()
	command := os.Args[1]

	switch command {
	case "createNode":
		if len(os.Args) != 4 {
			fmt.Println("Usage: programname createNode <name> <port>")
			os.Exit(1)
		}
		name := os.Args[2]
		portStr := os.Args[3]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("Invalid port number:", err)
			os.Exit(1)
		}
		createNode(name, port)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}

	// Wait indefinitely
	select {}

}

// N O D E S

func createNode(name string, port int) {

	// Database Initialize
	dbFileName := fmt.Sprintf("%s.db", name)
	fmt.Println("Database name is", dbFileName)
	file, err := os.Create(dbFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("Database file was created successfully.")
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		fmt.Println("Error opening database file:", err)
		os.Exit(1)
	}
	defer db.Close()

	// Create or join a cluster
	cluster := "cluster.txt"

	// Check if cluster exists
	_, err = os.Stat(cluster)

	// Handle cluster existence
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Cluster", cluster, "does not exist.")
			createCluster(cluster, port)
		} else {
			fmt.Println("Error checking file:", err)
		}
	} else {
		// join existing cluster
		err := joinCluster(cluster, port)
		if err != nil {
			fmt.Println("Error writing port", port, "to file:", err)
			return
		}
		fmt.Println("Successfully added a node to the cluster")
	}

	// Server/Node Initialize
	router := gin.Default()
	hostAddress := fmt.Sprintf("%s%d", "localhost:", port)
	router.Run(hostAddress)
}

// C L U S T E R

func createCluster(cluster string, port int) {
	// Open the file for writing, creating it if it doesn't exist
	file, err := os.Create(cluster)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed even on errors

	// Write some content to the file
	portString := strconv.Itoa(port)
	_, err = file.Write([]byte(portString + "\n"))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully created the cluster")
}

func joinCluster(cluster string, port int) error {
	// Write some content to the file
	file, err := os.OpenFile(cluster, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert port number to string
	portString := strconv.Itoa(port)

	// Write the port number followed by a newline character
	_, err = file.Write([]byte(portString + "\n"))
	return err
}

// P O R T S

func removePortFromFile(port int) error {
	mutex.Lock()
	defer mutex.Unlock()

	// Read the list of ports from the file
	portsFromFile, err := readPortsFromFile()
	if err != nil {
		return err
	}

	// Remove the specified port from the list
	var updatedPorts []int
	for _, p := range portsFromFile {
		if p != port {
			updatedPorts = append(updatedPorts, p)
		}
	}

	// Write the updated list of ports back to the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, p := range updatedPorts {
		_, err := fmt.Fprintf(file, "%d\n", p)
		if err != nil {
			return err
		}
	}

	return nil
}

func readPortsFromFile() ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ports []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line) // Trim leading and trailing whitespaces
		if line == "" {
			continue // Skip empty lines
		}
		port, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}
	return ports, scanner.Err()
}

func monitorPorts() {
	ticker := time.NewTicker(5 * time.Second) // Check every 5 seconds
	for range ticker.C {
		// Read the list of ports from the text file
		portsFromFile, err := readPortsFromFile()
		if err != nil {
			log.Printf("Error reading ports from file: %v", err)
			continue
		}

		// Iterate over the ports in the file and check if they are still running
		for _, port := range portsFromFile {
			if !isPortListening(port) {
				// Port is not running, remove it from the file
				removePortFromFile(port)
				log.Printf("Port %d terminated and removed from the file", port)
			}
		}
	}
	fmt.Println("Successfully created the cluster")
}

func isPortListening(port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// R A F T
