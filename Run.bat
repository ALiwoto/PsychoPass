:: this bat file is created by ALiwoto (woto@kaizoku.cyou)
@echo off
TITLE Building PsychoPass binary file
go mod tidy
go build -o PsychoPass.exe
cd .\docs
.\Build.bat && cd .. && PsychoPass.exe
