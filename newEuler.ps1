param(
    [Parameter(Mandatory=$true)]
    [int]$num
)

$folder = "$num"
if (Test-Path -Path $folder) {
    Write-Host "Folder $folder already exists. Exiting."
    exit
}
New-Item -ItemType Directory -Path $folder -Force | Out-Null
Set-Location $folder
go mod init $folder
@'
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
'@ | Set-Content "$num.go"
Set-Location ..