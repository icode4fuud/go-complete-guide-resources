# Go - The Complete Guide

# How to search file content in Linux/Bash
1. grep -R "ig4llc.com/db" .

# How to search file content in Windows/Powershell
1. Get-ChildItem -Recurse -Filter *.go | Select-String "ig4llc.com/db"


# To run the app and the migrations use the following command
go run ./cmd/api
go run ./cmd/migrate --action=up
go run ./cmd/migrate --action=down
go run ./cmd/migrate --action=status


