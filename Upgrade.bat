@echo off
git pull && powershell -command "Stop-service -Force -name "Psychopass" -ErrorAction SilentlyContinue; go mod vendor; go build; Start-service -name "Psychopass""
:: Hail Hydra
