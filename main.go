package main

import (
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ServerURL  string        `mapstructure:"server_url"`
	NumClients int           `mapstructure:"num_clients"`
	Delay      time.Duration `mapstructure:"delay"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // or specify the path

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
func main() {
	var rootCmd = &cobra.Command{
		Use:   "test",
		Short: "Run a WebSocket load test",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := LoadConfig()
			if err != nil {
				log.Fatalf("Error loading configuration: %v", err)
			}

			runLoadTest(config)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func runLoadTest(config Config) {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i < config.NumClients; i++ {
		wg.Add(1)
		time.Sleep(config.Delay)
		go simulateCompleteFlow(i, &wg, config)
	}

	wg.Wait()
	duration := time.Since(start)
	log.Printf("Load test completed in %v", duration)
}

func simulateCompleteFlow(clientID int, wg *sync.WaitGroup, config Config) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/spicy-dice"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("failed to connect: %w", err)
		return
	}
	defer conn.Close()
	defer wg.Done()
	if err := simulateWalletRequest(conn, clientID); err != nil {
		log.Printf("Client %d: Wallet request failed: %v", clientID, err)
		return
	}
	time.Sleep(config.Delay)

	if err := simulatePlayRequest(conn, clientID); err != nil {
		log.Printf("Client %d: Play request failed: %v", clientID, err)
		return
	}
	time.Sleep(config.Delay)

	if err := simulateEndPlayRequest(conn, clientID); err != nil {
		log.Printf("Client %d: End play request failed: %v", clientID, err)
		return
	}

	log.Printf("Client %d: Complete flow successful", clientID)
}
