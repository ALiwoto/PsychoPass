@echo off
TITLE Build latest version
echo.
cd.. && cd psychopass && git pull && powershell -command "Stop-service -Force -name "Psychopass" -ErrorAction SilentlyContinue; go build; Start-service -name "Psychopass""
timeout 4
exit
:: Hail Hydra
