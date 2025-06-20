
# Light Sequence Memory Game

![Game Concept](./data/media.jpg)

## Overview

This project is an interactive memory game where the system generates randomized light sequences using LEDs connected to an Arduino Uno. Players must observe the sequence and reproduce it accurately by pressing the corresponding push buttons in the correct order.

## Technology Stack

- **Hardware:** Arduino Uno with 3 LEDs and 3 push buttons
- **Backend:** Golang (`net/http`) serving the web interface
- **Utilities:** Python scripts for auxiliary tasks and utilities

## Getting Started

### Prerequisites

- Arduino Uno connected to your computer (ensure it is recognized as `/dev/ttyUSB0` on Linux or the appropriate COM port on Windows)
- Required dependencies for running Golang and Python scripts

## Hardware Setup

<figure>
  <img src="./data/circuit.png" alt="Arduino Circuit" style="width:70%">
  <figcaption>Circuit diagram of the Arduino Uno setup</figcaption>
</figure>

## Running the Game

1. Connect the Arduino Uno to your computer.
2. Verify the device is available at `/dev/ttyUSB0` (or adjust the script accordingly).
3. Run the startup script:
   ```bash
   ./run.sh
   ```
4. Open web browser and navigate to: `http://localhost:8080`
## Web Interface

<figure>
  <img src="./data/media2.png" alt="Web Interface" style="width:70%">
  <figcaption>Screenshot of the web interface to play the game</figcaption>
</figure>

---
