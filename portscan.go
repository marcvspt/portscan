package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
  "sync"
	"syscall"
	"time"
)

func main() {
  // Agregar manejo de señal de interrupción
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
    <-sigc
    fmt.Println("\n[!] Exiting...")
		os.Exit(0)
  }()

  // Definir los parámetros de entrada
  hostPtr := flag.String("h", "", "Host to scan (IPv4 o IPv6)")
  portPtr := flag.String("p", "", "Port or range of ports to be scanned (Example: -p 22 or -p 1-1024)")
  timeoutPtr := flag.Int("timeout", 1000, "Response timeout in milliseconds for each port")
  attemptsPtr := flag.Int("attempts", 1, "Number of connection attempts for each port")
  helpPtr := flag.Bool("help", false, "Show this help panel")
  flag.Parse()

  if *helpPtr || (flag.NFlag() == 0 && flag.NArg() == 0) || *hostPtr == "" || *portPtr == "" {
    fmt.Println("Golang Port Scanner")
    fmt.Println("Usage:")
    fmt.Println("\t./portscan -h example.com -p 80 [options]\n\t./portscan -h 2001:4860:4860::8888 -p 1-1000")
    fmt.Println("Options:")
    flag.PrintDefaults()
    return
  }

  // Parsear el rango de puertos
  ports := make([]int, 0)
  portRange := strings.Split(*portPtr, "-")
  startPort, _ := strconv.Atoi(portRange[0])
  endPort := startPort
  if len(portRange) > 1 {
    endPort, _ = strconv.Atoi(portRange[1])
  }

  for port := startPort; port <= endPort; port++ {
    ports = append(ports, port)
  }

  // Escanear los puertos
  protocol := "tcp"
  fmt.Printf("Scanning %s://%s on port %s\n", protocol, *hostPtr, *portPtr)
  var wg sync.WaitGroup
  lock := sync.Mutex{}
  openPorts := make([]int, 0)

  for _, port := range ports {
    wg.Add(1)
    go func(p int) {
      defer wg.Done()
      address := net.JoinHostPort(*hostPtr, strconv.Itoa(p))

      for i := 0; i < *attemptsPtr; i++ {
        conn, err := net.DialTimeout(protocol, address, time.Duration(*timeoutPtr)*time.Millisecond)

        if err != nil {
          continue
        }

        conn.Close()
        lock.Lock()
        openPorts = append(openPorts, p)
        lock.Unlock()
        break
      }
    }(port)
  }
  wg.Wait()

  if len(openPorts) > 0 {
    for _, port := range openPorts {
      fmt.Printf("[+] Port %d Open\n", port)
    }
  } else {
    fmt.Println("No open ports were found.")
  }
}
