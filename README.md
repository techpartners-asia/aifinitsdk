# 🤖 Aifinit SDK

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/aifinitsdk)](https://goreportcard.com/report/github.com/yourusername/aifinitsdk)

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
go get git.techpartners.asia/mtm/thirdparty/aifinitsdk
```

### Basic Usage

```go
package main

import (
    "log"
    "time"
    
    "git.techpartners.asia/mtm/thirdparty/aifinitsdk"
)

func main() {
    // Initialize client with your credentials
    credentials := aifinitsdk.Crendetials{
        MerchantCode: "your_merchant_code",
        SecretKey:    "your_secret_key",
    }
    
    client := aifinitsdk.New(credentials)
    
    // Get signature for authentication
    signature, err := client.GetSignature(time.Now().UnixMilli())
    if err != nil {
        log.Fatal(err)
    }
    
    // Use the signature for API calls
    // ...
}
```

## 📚 Core Components

### Vending Machine Management
- Device activation
- Machine listing and details
- Device information
- People flow monitoring
- Machine control and settings

### Product Management
- Product listing and details
- Product applications
- Price updates
- Inventory tracking
- Mutual exclusion rules

### Operations
- Door control
- Order management
- Video recording
- Sold goods tracking
- Real-time monitoring

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
