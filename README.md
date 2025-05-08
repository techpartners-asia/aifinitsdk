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
    aifinitsdk_device "github.com/techpartners-asia/aifinitsdk/device_manage"
    aifinitsdk_product "github.com/techpartners-asia/aifinitsdk/product_manage"
    aifinitsdk_operation "github.com/techpartners-asia/aifinitsdk/operation"
)

func main() {
    // Initialize client with credentials from environment variables
    credentials := aifinitsdk.Crendetials{
        MerchantCode: os.Getenv("MERCHANT_CODE"),
        SecretKey:    os.Getenv("SECRET_KEY"),
    }
    
    client := aifinitsdk.New(credentials)
    
    // Get signature for authentication
    signature, err := client.GetSignature(time.Now().UnixMilli())
    if err != nil {
        panic(err)
    }
    
    // Initialize device client
    deviceClient := aifinitsdk_device.NewDeviceClient(client, os.Getenv("DEVICE_CODE"))
    
    // Get device information
    device, err := deviceClient.DeviceInfo()
    if err != nil {
        panic(err)
    }
    fmt.Println("Engine Status:", device.Data.EngineOn)
    
    // List devices
    listDeviceRequest := aifinitsdk_device.ListRequest{
        Page:  1,
        Limit: 10,
    }
    listDevice, err := deviceClient.List(&listDeviceRequest)
    if err != nil {
        panic(err)
    }
    
    // Get device details
    detail, err := deviceClient.Detail()
    if err != nil {
        panic(err)
    }
    
    // Get people flow data
    peopleFlowRequest := aifinitsdk_device.PeopleFlowRequest{
        Field:          "visitorCount",
        StartTimeStamp: 1591415672104,
        EndTimeStamp:   1591015672104,
        Codes:          []string{"a2acd5b95919"},
    }
    peopleFlow, err := deviceClient.PeopleFlow(&peopleFlowRequest)
    if err != nil {
        panic(err)
    }
    
    // Control device
    controlRequest := aifinitsdk_device.ControlRequest{
        EngineOn: 1,
    }
    control, err := deviceClient.Control(&controlRequest)
    if err != nil {
        panic(err)
    }
    
    // Initialize product client
    productClient := aifinitsdk_product.NewProductClient(client)
    
    // Get product list
    products, err := productClient.GetProductList(1, 10)
    if err != nil {
        panic(err)
    }
    
    // Get product detail
    productDetail, err := productClient.GetProductDetail(products.Data.Rows[0].ItemCode)
    if err != nil {
        panic(err)
    }
    
    // Check product mutual exclusion
    mutualExclusion, err := productClient.GetProductMutualExclusion(&aifinitsdk_product.MutualExclusionRequest{
        ItemCodes: []string{products.Data.Rows[0].ItemCode},
    })
    if err != nil {
        panic(err)
    }
    
    // List product applications
    listProductApplication, err := productClient.ListProductApplication(&aifinitsdk_product.ListProductApplicationParams{
        Page:        1,
        PageSize:    10,
        ApplyStatus: 1,
        GoodsName:   "",
        QrCodes:     "",
    })
    if err != nil {
        panic(err)
    }
    
    // Initialize operation client
    operationClient := aifinitsdk_operation.NewOperationClientImpl(client, os.Getenv("DEVICE_CODE"))
    
    // Open door
    openDoorRequest := aifinitsdk_operation.OpenDoorRequest{
        Type:           2, // OpenDoorForReplenishment
        RequestID:      "1234567892",
        UserCode:       "1234567892",
        LocalTimeStamp: time.Now().UnixMilli(),
    }
    openDoor, err := operationClient.OpenDoor(&openDoorRequest)
    if err != nil {
        panic(err)
    }
}

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
    credentials := aifinitsdk.Crendetials{
        MerchantCode: os.Getenv("MERCHANT_CODE"),
        SecretKey:    os.Getenv("SECRET_KEY"),
    }
    
    client := aifinitsdk.New(credentials)
    
    // Enable debug mode for the client
    client.SetConfig(aifinitsdk.Config{
        Debug: true,
    })
    
    // ... rest of your code
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

## 📦 Modules

### Device Management (`device_manage/`)
- Device activation and registration
- Device status monitoring
- Device configuration
- Device information retrieval
- People flow monitoring

```go
// Get device information
deviceClient := aifinitsdk_device.NewDeviceClient(client, "VM001")
deviceInfo, err := deviceClient.DeviceInfo()
if err != nil {
    log.Fatal(err)
}

// Get device details
details, err := deviceClient.Detail()
if err != nil {
    log.Fatal(err)
}

// Update device settings
updateReq := &aifinitsdk_device.UpdateRequest{
    Name:          "Vending Machine 1",
    Code:          "VM001",
    ScanCode:      "SCAN001",
    ContactNumber: "1234567890",
    Volume:        80,
    AdVolume:      20,
    Temp:          4,
    EngineOn:      1,
}
updateResp, err := deviceClient.Update(updateReq)
if err != nil {
    log.Fatal(err)
}

// Control device
controlReq := &aifinitsdk_device.ControlRequest{
    Volume:   80,
    AdVolume: 20,
    Temp:     4,
    EngineOn: 1,
}
controlResp, err := deviceClient.Control(controlReq)
if err != nil {
    log.Fatal(err)
}

// Get people flow data
flowReq := &aifinitsdk_device.PeopleFlowRequest{
    Field:          "hour",
    StartTimeStamp: time.Now().Add(-24 * time.Hour).Unix(),
    EndTimeStamp:   time.Now().Unix(),
    Codes:          []string{"VM001"},
}
flowResp, err := deviceClient.PeopleFlow(flowReq)
if err != nil {
    log.Fatal(err)
}
```

### Product Management (`product_manage/`)
- Product listing and details
- Product applications
- Price updates and management
- Inventory tracking
- Product mutual exclusion rules

```go
// Initialize product client
productClient := aifinitsdk_product.NewProductClient(client)

// Get product list
products, err := productClient.GetProductList(1, 10)
if err != nil {
    log.Fatal(err)
}

// Get product detail
productDetail, err := productClient.GetProductDetail("ITEM001")
if err != nil {
    log.Fatal(err)
}

// Check product mutual exclusion
mutualExclusionReq := &aifinitsdk_product.MutualExclusionRequest{
    ItemCodes: []string{"ITEM001", "ITEM002"},
}
mutualExclusion, err := productClient.GetProductMutualExclusion(mutualExclusionReq)
if err != nil {
    log.Fatal(err)
}

// Create new product application
newProductReq := &aifinitsdk_product.NewProductApplicationRequest{
    Product: &aifinitsdk_product.Product{
        Name:           "New Product",
        Price:          200,
        Weight:         100,
        WeightVariance: 5,
        ItemCode:       "ITEM003",
        CollType:       1,
        Status:         1,
    },
}
newProduct, err := productClient.NewProductApplication(newProductReq)
if err != nil {
    log.Fatal(err)
}

// List product applications
listParams := &aifinitsdk_product.ListProductApplicationParams{
    Page:        1,
    PageSize:    10,
    ApplyStatus: 1,
    GoodsName:   "Product",
    QrCodes:     "QR001",
}
applications, err := productClient.ListProductApplication(listParams)
if err != nil {
    log.Fatal(err)
}
```

### Operations (`operation/`)
- Door control operations
- Order video retrieval
- Sold goods management
- Door search and status
- Error handling for operations

```go
// Initialize operation client
operationClient := aifinitsdk_operation.NewOperationClientImpl(client, "VM001")

// Open door for shopping
openDoorReq := &aifinitsdk_operation.OpenDoorRequest{
    Type:           int(aifinitsdk_operation.OpenDoorForShopping),
    RequestID:      "REQ001",
    UserCode:       "USER001",
    LocalTimeStamp: time.Now().Unix(),
}
openDoorResp, err := operationClient.OpenDoor(openDoorReq)
if err != nil {
    log.Fatal(err)
}

// Get sold goods
soldGoods, err := operationClient.GetSoldGoods()
if err != nil {
    log.Fatal(err)
}

// Update sold goods
updateSoldGoodsReq := &aifinitsdk_operation.UpdateSoldGoodsRequest{
    {
        ItemCode:      "ITEM001",
        ActualPrice:   200,
        OriginalPrice: 250,
    },
}
updateSoldGoodsResp, err := operationClient.UpdateSoldGoods(updateSoldGoodsReq)
if err != nil {
    log.Fatal(err)
}

// Search open door status
searchReq := &aifinitsdk_operation.SearchOpenDoorRequest{
    Type:      aifinitsdk_operation.OpenDoorForShopping,
    RequestID: "REQ001",
}
searchResp, err := operationClient.SearchOpenDoor(searchReq)
if err != nil {
    log.Fatal(err)
}

// Get order video
videoReq := &aifinitsdk_operation.GetOrderVideoRequest{
    RequestID: "REQ001",
    Type:      aifinitsdk_operation.OpenDoorForShopping,
}
videoResp, err := operationClient.GetOrderVideo(videoReq)
if err != nil {
    log.Fatal(err)
}

// Update product price
priceUpdateReq := &aifinitsdk_operation.ProductPriceUpdateRequest{
    VmCodes: []string{"VM001"},
    Items: []aifinitsdk_operation.Good{
        {
            ItemCode:      "ITEM001",
            ActualPrice:   200,
            OriginalPrice: 250,
        },
    },
}
priceUpdateResp, err := operationClient.ProductPriceUpdate(priceUpdateReq)
if err != nil {
    log.Fatal(err)
}
```

### Constants (`constants/`)
- Error codes and messages
- API endpoints
- Configuration constants
- Status codes
- Common values

```go
// Using API endpoints
deviceActivationURL := aifinitsdk_constants.Post_DeviceActivation
vendingMachineListURL := aifinitsdk_constants.Get_VendingMachineList
productListURL := aifinitsdk_constants.Get_ProductList
openDoorURL := aifinitsdk_constants.Put_OpenDoor

// Using error constants
if err != nil {
    switch err.(type) {
    case aifinitsdk_operation.OpenDoorError:
        if err.(aifinitsdk_operation.OpenDoorError) == aifinitsdk_operation.ErrOpenDoorTimeout {
            // Handle timeout error
        }
    case aifinitsdk_operation.ProductPriceUpdateError:
        if err.(aifinitsdk_operation.ProductPriceUpdateError) == aifinitsdk_operation.ErrProductPriceUpdateNoOperatingPermissions {
            // Handle permission error
        }
    }
}

// Using video status constants
if videoResp.Data.VideoStatus == aifinitsdk_operation.VideoStatusUploadComplete {
    // Video is ready
} else if videoResp.Data.VideoStatus == aifinitsdk_operation.VideoStatusPendingUpload {
    // Video is still uploading
}
```

### Core (`./`)
- Client initialization and configuration
- Authentication and encryption
- Base models and interfaces
- Utility functions
- Common operations

```go
// Initialize client with credentials
credentials := aifinitsdk.Crendetials{
    MerchantCode: "your_merchant_code",
    SecretKey:    "your_secret_key",
}
client := aifinitsdk.New(credentials)

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
