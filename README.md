# AudioLoom: MIDI-OSC Integration Microservice

## Overview

**AudioLoom: MIDI-OSC Integration** is a lightweight and modular microservice designed to facilitate seamless communication between Digital Audio Workstations (DAWs) and custom applications using the Open Sound Control (OSC) protocol and MIDI messages. The microservice acts as a bridge to collect, process, and send playback, position, and control data to enhance the interoperability of your audio production workflows.

---

## Features

- **OSC Server and Client Integration**
  - Supports receiving playback and position data from DAWs like Ardour and Reaper.
  - Provides feedback to DAWs via OSC commands for transport and other control functionalities.

- **MIDI Integration**
  - Handles MIDI data such as note events, control changes, and channel-specific data.

- **Multi-Client Support**
  - Connect multiple OSC clients and DAWs simultaneously.
  - Provides data collection and control for different audio tracks or sessions.

- **Customizable Configuration**
  - Fully configurable OSC server and client settings for ports, addresses, and paths.
  - Supports protocol versions for fault tolerance (e.g., MIDI 1.0, MIDI 2.0).

---

## Installation

### Prerequisites

1. **Go Language**: Install [Go](https://go.dev/dl/) version 1.20 or higher.
2. **Dependencies**: Ensure the following Go libraries are installed:
   - `github.com/hypebeast/go-osc/osc`
3. **Digital Audio Workstation (DAW)**:
   - Enable OSC communication in your DAW (e.g., Ardour or Reaper).


