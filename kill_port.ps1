$envFile = ".env"

if (Test-Path $envFile) {
    Get-Content $envFile | ForEach-Object {
        if ($_ -match "^PORT=(\d+)") {
            $PORT = $matches[1]
            Write-Host "Found PORT=$PORT in .env"
            
            $connections = Get-NetTCPConnection -LocalPort $PORT -ErrorAction SilentlyContinue
            if ($connections) {
                $targetPids = $connections | Select-Object -ExpandProperty OwningProcess -Unique
                foreach ($p in $targetPids) {
                    if ($p -ne 0) {
                        Write-Host "Killing process on port $PORT (PID: $p)..."
                        Stop-Process -Id $p -Force -ErrorAction SilentlyContinue
                        Write-Host "Process $p killed."
                    }
                }
            } else {
                Write-Host "No process found listening on port $PORT."
            }
        }
    }
} else {
    Write-Host ".env file not found."
}
