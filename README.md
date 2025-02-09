# DockMon
Backend for Docker monitoring service

# Features
 - Receives data from /api/machine enpoint (PUT method) and stores it in PostgreSQL database
 - Sends data via /api/machines endpoint (GET method)

# Domain model
 - ip key represents IPv4 address of a Docker container
 - ping_time defines time of ping execution in seconds
 - success defines result of ping execution
 - last_success defines date of last successful ping (in RFC3339)
```json
{
  "ip": "172.18.0.4",
  "ping_time": 3,
  "success": true,
  "last_success": "2025-02-09T14:32:05.49081Z"
}
```

# Standalone launch
 1. Specify host and port in [config.yaml](config/config.yaml), that will be listened after launch
 2. Specify options of your PostgreSQL instance and provide names environment variables in fields ending with EnvKey
 3. Specify environment variables in .env file
 4. Build and launch via docker 

