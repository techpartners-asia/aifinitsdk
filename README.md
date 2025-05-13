# 🤖 Aifinit SDK

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/ainfinitsdk)](https://goreportcard.com/report/github.com/yourusername/ainfinitsdk)

A powerful Go SDK for integrating with the Aifinit vending machine platform. Manage vending machines, products, and operations with ease.

</div>

## ✨ Features

- 🔐 Secure authentication and encryption
- 🏪 Vending machine management
- 📦 Product management
- 🚪 Door control operations
- 📊 Real-time monitoring
- 🔄 Inventory management
- 📱 Mobile-friendly API

## 🚀 Quick Start

### Installation

```bash
go get github.com/techpartners-asia/aifinitsdk
```

### Basic Usage

```go
package main

import (
    "fmt"
    "os"
    "time"
    
    "github.com/techpartners-asia/aifinitsdk"
)

func main() {
    // Initialize client with credentials from environment variables
    credentials := ainfinitsdk.Crendetials{
        MerchantCode: os.Getenv("MERCHANT_CODE"),
        SecretKey:    os.Getenv("SECRET_KEY"),
    }
    
    client := ainfinitsdk.New(credentials)
    
    // Get signature for authentication
    signature, err := client.GetSignature(time.Now().UnixMilli())
    if err != nil {
        panic(err)
    }
}
```

### Debug Mode

To enable debug mode and see detailed logs, you need to configure both the client and logrus. Here's how to do it:

```go
import (
    "github.com/sirupsen/logrus"
    "github.com/techpartners-asia/aifinitsdk"
)

func main() {
    // Set logrus to debug level
    logrus.SetLevel(logrus.DebugLevel)
    
    // Initialize client with credentials
    credentials := ainfinitsdk.Crendetials{
        MerchantCode: os.Getenv("MERCHANT_CODE"),
        SecretKey:    os.Getenv("SECRET_KEY"),
    }
    
    client := ainfinitsdk.New(credentials)
    
    // Enable debug mode for the client
    client.SetConfig(ainfinitsdk.Config{
        Debug: true,
    })
    
    // ... rest of your code
}
```

## 📚 Core Components

### Core (`./`)
- Client initialization and configuration
- Authentication and encryption
- Base models and interfaces
- Utility functions
- Common operations

```go
// Initialize client with credentials
credentials := ainfinitsdk.Crendetials{
    MerchantCode: "your_merchant_code",
    SecretKey:    "your_secret_key",
}
client := ainfinitsdk.New(credentials)

// Get signature for authentication
timestamp := time.Now().UnixMilli()
signature, err := client.GetSignature(timestamp)
if err != nil {
    log.Fatal(err)
}
```

## 🔒 Security

The SDK implements secure authentication using:
- AES encryption
- Base64 encoding
- Timestamp-based signatures
- Secure key management

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
