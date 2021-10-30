@echo off
TITLE PsychoPass
echo.
:: stop service, delete existing pull and build log files, git pull, build, then start service again. 
nssm stop PsychoPass && del pull_log.txt /F && del build_log.txt /F && git pull > pull_log.txt && go build > build_log.txt && nssm start PsychoPass && echo All Done, check once. 
:: Hail Hydra
