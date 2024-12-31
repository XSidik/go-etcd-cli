package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcdClient *clientv3.Client

func main() {
	var err error
	// Load .env file
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv("etcdEndpoints")},
		Username:    os.Getenv("etcdUsername"),
		Password:    os.Getenv("etcdPassword"),
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("Failed to connect to etcd: %v\n", err)
		return
	}

	defer etcdClient.Close()
	fmt.Println("Connected to etcd")

	fmt.Println("Interactive Console with etcd (Type 'help' for commands)")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		handleCommand(input)
	}
}

func handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("Invalid command. Type 'help' for a list of commands.")
		return
	}

	command := parts[0]
	switch command {
	case "help":
		showHelp()
	case "set":
		if len(parts) < 3 {
			fmt.Println("Usage: set <key> <value>")
			return
		}

		key := parts[1]
		value := strings.Join(parts[2:], " ")
		setKey(key, value)

	case "get":
		if len(parts) < 2 {
			fmt.Println("Usage: get <key>")
			return
		}

		key := parts[1]
		getKey(key)

	case "delete":
		if len(parts) < 2 {
			fmt.Println("Usage: delete <key>")
			return
		}

		key := parts[1]
		deleteKey(key)

	case "list":
		listKeys()

	default:
		fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", command)
	}
}

func showHelp() {
	fmt.Print(`
Commands:
  help                Show this help message
  set <key> <value>   Set a key-value pair in etcd
  get <key>           Get the value of a key from etcd
  delete <key>        Delete a key from etcd
  list                List all keys stored in etcd
  exit                Exit the console
`)
}

func setKey(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := etcdClient.Put(ctx, key, value)
	if err != nil {
		fmt.Printf("Failed to set key '%s': %v\n", key, err)
		return
	}
	fmt.Printf("Key '%s' set to '%s'\n", key, value)
}

func getKey(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := etcdClient.Get(ctx, key)
	if err != nil {
		fmt.Printf("Failed to get key '%s': %v\n", key, err)
		return
	}

	if len(resp.Kvs) == 0 {
		fmt.Printf("Key '%s' not found\n", key)
		return
	}

	fmt.Printf("Key '%s' = '%s'\n", key, string(resp.Kvs[0].Value))
}

func deleteKey(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := etcdClient.Delete(ctx, key)
	if err != nil {
		fmt.Printf("Failed to delete key '%s': %v\n", key, err)
		return
	}
	fmt.Printf("Key '%s' deleted\n", key)
}

func listKeys() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := etcdClient.Get(ctx, "", clientv3.WithPrefix()) // Fetch all keys
	if err != nil {
		fmt.Printf("Failed to list keys: %v\n", err)
		return
	}
	if len(resp.Kvs) == 0 {
		fmt.Println("No keys found.")
		return
	}
	fmt.Println("Stored Keys:")
	for _, kv := range resp.Kvs {
		fmt.Printf("%s = %s\n", kv.Key, kv.Value)
	}
}
