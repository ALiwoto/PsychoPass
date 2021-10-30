@echo off
del pull_log.txt /F
del build_log.txt /F
TITLE PsychoPass
echo.
nssm stop PsychoPass
git pull > pull_log.txt
go build > build_log.txt
cls
nssm start PsychoPass
:: Hail Hydra
