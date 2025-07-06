# Brabus - Machine Metrics Anomaly Detection System

*"Better to detect early than to repair late""*

## 📌 Overview

Brabus is a powerful monitoring application designed to collect machine metrics and analyze them for anomalies. When anomalies are detected, the system sends instant alerts through a Telegram bot, allowing for quick response to potential issues.

**Created by**: Alice088  
**Current Version**: 1.0.0

## 🚀 Key Features

- Real-time CPU monitoring
- Advanced anomaly detection algorithms
- Instant Telegram notifications [Not support now, only CLI yet]
- Lightweight and efficient
- Easy deployment with Docker

## 🔍 Currently Detected Anomalies

| Anomaly Type               | Description                          | Thresholds                          | Severity Level |
|----------------------------|--------------------------------------|-------------------------------------|----------------|
| **High CPU Usage**         | Combined user+system CPU utilization | >90% sustained usage                | ⚠️ Warning    |
| **High IOWait**            | Excessive disk wait time             | >10% of CPU time                    | ⚠️ Warning    |
| **Sustained High Load**    | System load average                  | >70% of CPU cores (warning threshold)<br>>90% of CPU cores (critical threshold) | ⚠️ Warning/🛑 Critical |
| **Load Spike**             | Sudden increase in load              | 1min load > 1.5× 15min load         | 🛑 Critical    |

### Detailed Thresholds:

1. **CPU Usage Anomalies**
   ```go
   totalUsed := cpu.Average.User + cpu.Average.System
   if totalUsed > 90.0 // Triggers warning
   ```

2. **IOWait Anomalies**
   ```go
   if cpu.Average.IOWait > 10.0 // Triggers warning
   ```

3. **Load Average Anomalies**
   ```go
   // Warning threshold (70% of cores)
   t.value > cores*0.7
   
   // Critical threshold (90% of cores)
   t.value > cores*0.9
   ```

4. **Load Spike Detection**
   ```go
   if cpu.Load1 > cpu.Load15*1.5 // Triggers critical alert
   ```

The system provides graduated severity levels:
- ⚠️ **Warning**: Potential issues needing monitoring
- 🛑 **Critical**: Immediate attention required

All anomalies are delivered via Telegram bot with clear severity indicators and specific metric values.

## 🛠 Technologies

- **Backend**: Golang
- **Messaging**: NATS
- **Containerization**: Docker
- **Monitoring**: Custom anomaly detection algorithms

## 📦 Installation

### Prerequisites
- Docker and Docker Compose
- Go 1.23.0+
- Make

### Quick Start

1. Download [release](https://github.com/Alice088/brabus/releases) for your OS architecture OR build own via [scripts](https://github.com/Alice088/brabus/blob/master/scripts/build.sh)

2. Run brabus:
   ```bash
   make run
   ```

3. Stop brabus:
   ```bash
   make stop
   ```

## 🏃‍♂️ Usage

### Running the Application

Start all services:
```bash
make run
```

This will:
1. Launch Docker containers with NATS in background
2. Start Brabus monitoring service silently
3. Launch Banana (analyze and alert service) with console output

### Stopping the Application

To stop all services:
```bash
make stop
```

## 🛑 Command Reference

| Command        | Description                          |
|---------------|--------------------------------------|
| `make run`    | Start all services                   |
| `make stop`   | Stop all services                    |

## 📊 Future Roadmap
- [ ] Notifications via telegram
- [ ] Memory usage monitoring
- [ ] Disk space anomaly detection
- [ ] Network latency monitoring
- [ ] Web dashboard for visualization
- [ ] Slack integration alongside Telegram

## 🤝 Contributing

Contributions are welcome! Please fork the repository and open a pull request with your improvements.

## 📄 License

MIT License - See [LICENSE](LICENSE) for more information.

---

**Happy Monitoring!** 🚨  