# Virtual Tabletop (VTT) API - Go Implementation

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![WebSocket](https://img.shields.io/badge/WebSocket-supported-brightgreen)
![REST API](https://img.shields.io/badge/REST_API-supported-blue)
![License](https://img.shields.io/badge/License-MIT-green)

A lightweight Virtual Tabletop (VTT) API built with Go that combines RESTful endpoints with real-time WebSocket functionality for tabletop gaming applications.

## Features

- **Dice Rolling Engine**: Supports complex dice notation (e.g., `3d6+2`, `4d10k3`, `2d20!>15`)
- **Real-time Updates**: WebSocket connections for instant game state synchronization
- **Room System**: Multiple isolated game sessions with separate states
- **RESTful API**: Standard HTTP endpoints for game management
- **Stateless Design**: No external dependencies required (runs on pure Go)

## Installation

1. Ensure you have Go 1.20 or later installed
2. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/go-vtt-api.git
   cd go-vtt-api
   ```

3. Build and run:

   ```bash
   go build
   ./go-vtt-api
   ```

The server will start on port 8080.
