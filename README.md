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

## Dice Rolling System

This system implements a powerful syntax for rolling dice in various RPG systems, supporting a wide range of modifiers and specific rules.

## Basic Syntax

The basic dice rolling syntax is:

``` text
NdX
```

Where:

- N is the number of dice to roll
- X is the die type (e.g., d6, d20, f for Fudge dice)

## Available Modifiers

### Exploding Dice (ex)

```text
NdXexY
```

- When a die rolls Y or higher, roll it again and add the result  
- Example: `2d6ex4` - two 6-sided dice that explode on rolls of 4+

### Re-Roll (re)

```text
NdXreY
```

- Re-roll when the result is Y or lower  
- Example: `2d6re2` - two 6-sided dice that re-roll on 2 or lower

### Keep Highest (kh)

```text
NdXkhY
```

- Keep the Y highest results  
- Example: `5d6kh3` - roll 5 dice and keep the 3 highest

### Keep Lowest (kl)

```text
NdXklY
```

- Keep the Y lowest results  
- Example: `5d6kl3` - roll 5 dice and keep the 3 lowest

### Success Counting (su)

```text
NdXsuY
```

- Count as success when result is Y or higher  
- Example: `2d6su4` - two 6-sided dice where 4+ counts as success

### Failure Counting (fa)

```text
NdXfaY
```

- Count as failure when result is Y or lower  
- Example: `2d6fa2` - two 6-sided dice where 2 or lower counts as failure

### Yin Yang (yy)

```text
NdXyy
```

- Alternate between Yin and Yang for valid dice  
- Example: `2d6yy` - two dice where first is Yin and second is Yang

### Red or Blue (rb)

```text
NdXrb
```

- Compare two dice (red vs blue)  
- Example: `2d6rb` - two dice where highest determines winning color

### Critical by Value (csv/cfv)

```text
NdXcsvY
NdXcfvY
```

- Critical success when value is Y or higher (csv)  
- Critical failure when value is Y or lower (cfv)  
- Example: `2d6csv5` - two dice where 5+ is critical success

### Critical by Equality (cse/cfe)

```text
NdXcseY
NdXcfeY
```

- Critical when all dice are equal and ≥ Y (cse)  
- Critical when all dice are equal and ≤ Y (cfe)  
- Example: `2d6cse4` - two dice where equal rolls of 4+ are critical success

## Combined Examples

```text
3d6ex5re2kh2
```

- Roll 3d6, explode on 5+, re-roll on 2-, keep 2 highest

```text
2d6yycsv4
```

- Roll 2d6 with alternating Yin/Yang, critical success on 4+

```text
4d6kl2cfv1
```

- Roll 4d6, keep 2 lowest, critical failure on 1

## Special Dice Support

The system supports Fudge dice (f) that can roll -1, 0, or +1.

Example: `2df` - roll two Fudge dice.
