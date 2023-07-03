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
