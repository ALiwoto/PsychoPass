@echo off
git pull && powershell -command "Stop-service -Force -name "Psychopass" -ErrorAction SilentlyContinue; go build; Start-service -name "Psychopass""
:: Hail Hydra
