# Go-Svelte Link Shortener

This is a URL shortener project developed with Go (backend) and Svelte (frontend). It was created as a learning project and includes some modifications compared to the original version available on the NerdCademy YouTube channel and I would like to acknowledge the valuable contribution of the author.

click [here], to access the original project.

# Table of Contents

1. [Backend](#backend)
   - [Database Configuration (Docker Compose)](#database-configuration-docker-compose)
   - [Environment Variables](#environment-variables)
   - [Scripts for Linux and Windows](#scripts-for-linux-and-windows)
   - [Backend Dependencies](#backend-dependencies)

2. [Frontend](#frontend)
_________________________________________________________________________________________________________________________________________

## Backend

The backend is developed in Go and uses PostgreSQL as the database. To run the backend, you have the option to use a Docker Compose setup. However, feel free to choose another method if you prefer.

### Database Configuration (Docker Compose)

If you have chosen to use Docker Compose, you can set up the database by creating a docker-compose.yml file with the following content:

```yaml
version: '2.17.3'
services:
  dev-db:
    container_name: link-shortener-dev
    image: postgres:14.7
    ports:
      - 5434:5432
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 123
      POSTGRES_DB: go
    networks:
      - goserver

networks:
  goserver:
```
You can save this content to a file named docker-compose.yml. This configuration sets up a PostgreSQL database container using the specified image and environment variables. The container will be accessible on port 5434 of your local machine, and the database name, username, and password are set as go, postgres, and 123, respectively.

### Environment variables 

You need to create a file named .env in the root directory of the backend (./go-app). This file will be used by the backend server to load the environment variables during runtime.

```env
PORT=:8000
DATABASE_URL="postgresql://postgres:123@localhost:5434/go"
```

In this file, you can modify the values of the PORT and DATABASE_URL variables according to your needs. The PORT variable specifies the port number on which the backend server will listen. The DATABASE_URL variable contains the connection URL for the PostgreSQL database, including the username, password, host, port, and database name.

If you decide to change the port, please ensure that you also update the corresponding configuration in other relevant files to ensure proper functionality.

### Scripts for Linux and Windows

To simplify the process of running the project, two scripts have been provided: one for Windows (init.ps1) and another for bash (init.sh). You can choose the script that corresponds to your operating system and execute it to start the project. To do this, you just need to create one of the two following file:

#### Linux
```bash
#!/bin/bash

processFilePath="./link-shortener"
pidFilePath="./process.pid"
dockerComposeFilePath="./docker-compose.yml"
envFilePath="./.env"

# Function to check if the Docker container is running
function isDockerContainerRunning {
    containerName="link-shortener-dev"
    containerStatus=$(docker inspect --format='{{.State.Status}}' "$containerName" 2> /dev/null)
    [ "$containerStatus" == "running" ]
}

# Command to build the Go code
go build
if [ $? -eq 0 ]; then
    if ! isDockerContainerRunning; then
        echo "Starting the Docker container..."
        docker-compose -f "$dockerComposeFilePath" up -d
        # Wait for a few seconds for the container to fully start
        sleep 2
    fi

    # Start the application process
    "$processFilePath" "$envFilePath" &
    echo "Process started with PID $!" 

    trap 'cleanup' SIGINT

    # Function to perform cleanup actions
    function cleanup {
        echo "Choose the container action:"
        echo -e "1. Stop\n2. Remove\n3. Keep running"
        read -p "Enter the corresponding number: " containerAction

        case $containerAction in
            1)
                echo "Stopping the Docker container..."
                docker-compose -f "$dockerComposeFilePath" stop
                ;;
            2)
                echo "Removing the Docker container..."
                docker-compose -f "$dockerComposeFilePath" down
                ;;
            3)
                echo "Keeping the Docker container running."
                ;;
            *)
                echo "Invalid option. No action taken."
                ;;
        esac

        exit
    }

    # Wait for the process to finish
    wait
else
    echo "Build failed"
fi
```

#### Windows

```ps1

$processFilePath = "$PSScriptRoot\link-shortener.exe"
$pidFilePath = "$PSScriptRoot\process.pid"
$dockerComposeFilePath = "$PSScriptRoot\docker-compose.yml"
$envFilePath = "$PSScriptRoot\.env"

# Function to check if the Docker container is running
function IsDockerContainerRunning($containerName) {
    $containerStatus = docker inspect --format='{{.State.Status}}' $containerName 2> $null
    return ($containerStatus -eq "running")
}

# Command to build the Go code
go build
if ($LASTEXITCODE -eq 0) {
    if (-not (IsDockerContainerRunning "link-shortener-dev")) {
        Write-Host "Starting the Docker container..."
        docker-compose -f $dockerComposeFilePath up -d
        # Wait for a few seconds for the container to fully start
        Start-Sleep -Seconds 2
    }

    # Start the application process
    $process = Start-Process -FilePath $processFilePath -NoNewWindow -PassThru -ArgumentList $envFilePath
    $process.Id | Out-File $pidFilePath
    Write-Host "Process started with PID $($process.Id)" -ForegroundColor Magenta

    $job = Start-Job -ScriptBlock {
        param($processId)
        Wait-Process -Id $processId
    } -ArgumentList $process.Id
    Write-Host "Job started with ID $($job.Id)" -ForegroundColor DarkGray

    try {
        Wait-Job -Job $job
    } finally {
        # Prompt when pressing Ctrl + C to choose the container action
        Write-Host "Choose the container action:"
        Write-Host "1. Stop" -ForegroundColor Yellow -NoNewline
        Write-Host " || " -NoNewline
        Write-Host "2. Remove" -ForegroundColor Red -NoNewline
        Write-Host " || " -NoNewline
        Write-Host "3. Keep running" -ForegroundColor Green
        $containerAction = Read-Host "Enter the corresponding number"

        switch ($containerAction) {
            1 {
                Write-Host "Stopping the Docker container..."
                docker-compose -f $dockerComposeFilePath stop
            }
            2 {
                Write-Host "Removing the Docker container..."
                docker-compose -f $dockerComposeFilePath down
            }
            3 {
                Write-Host "Keeping the Docker container running."
            }
            default {
                Write-Host "Invalid option. No action taken."
            }
        }

        Stop-Job -Job $job | Remove-Job
    }
} else {
    Write-Host "Build failed" -ForegroundColor Red
}
```

### Backend dependencies 

To run the github.com/RodBarenco/link-shortener module, you will need to install the following dependencies:

Go 1.20 or a compatible version.
gorm.io/gorm v1.25.2
github.com/andybalholm/brotli v1.0.5 (indirect)
github.com/gofiber/fiber/v2 v2.47.0
github.com/google/uuid v1.3.0
github.com/jackc/pgpassfile v1.0.0 (indirect)
github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a (indirect)
github.com/jackc/pgx/v5 v5.3.1 (indirect)
github.com/klauspost/compress v1.16.3 (indirect)
github.com/mattn/go-colorable v0.1.13 (indirect)
github.com/mattn/go-isatty v0.0.19 (indirect)
github.com/mattn/go-runewidth v0.0.14 (indirect)
github.com/philhofer/fwd v1.1.2 (indirect)
github.com/rivo/uniseg v0.2.0 (indirect)
github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 (indirect)
github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee (indirect)
github.com/tinylib/msgp v1.1.8 (indirect)
github.com/valyala/bytebufferpool v1.0.0 (indirect)
github.com/valyala/fasthttp v1.47.0 (indirect)
github.com/valyala/tcplisten v1.0.0 (indirect)
golang.org/x/crypto v0.8.0 (indirect)
golang.org/x/sys v0.9.0 (indirect)
golang.org/x/text v0.9.0 (indirect)
github.com/jinzhu/inflection v1.0.0 (indirect)
github.com/jinzhu/now v1.1.5 (indirect)
github.com/joho/godotenv v1.5.1
gorm.io/driver/postgres v1.5.2
You can use the following command to install these dependencies:

```bash
go get github.com/RodBarenco/link-shortener
```

## Frontend

The frontend is developed in Svelte. To run the frontend, follow the instructions below:
Make sure you have Node.js installed on your system.
Install the project dependencies by running the following command in the ./svelte-viwe directory:

```bash
npm install
```

If you prefer using Yarn, run the following command instead:

```bash
yarn install
```

Start the frontend development server:

```bash
npm run dev
```

If you're using Yarn, use the following command:

```bash
yarn dev
```
This will start the Svelte development server, and you can access the frontend application in your browser at http://localhost:8080.
Feel free to explore and modify the project according to your needs and preferences.

Please let me know if there's anything else I can assist you with!

[here]: ttps://www.youtube.com/watch?v=bTLQT7W12dQ&t=1601