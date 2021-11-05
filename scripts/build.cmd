:: Zestaw znak¢w: OEM 852
:: Kodowanie nowej linii: Windows (RC LF)
:: Narz©dzie "rsrc" https://github.com/akavel/rsrc
:: Kompilator "go" https://golang.org/dl/
@ECHO OFF
setlocal
set "SCRIPTS_DIR=%~dp0"
set "PROJECT_DIR=%SCRIPTS_DIR%.."
set "BUILD_DIR=%PROJECT_DIR%\build"
set "MAIN_DIR=%PROJECT_DIR%\cmd\cvc"
cd "%MAIN_DIR%"
rsrc -ico icon.ico -manifest manifest.xml -o cvc.syso && go build && move "cvc.exe" "%BUILD_DIR%"
cd "%PROJECT_DIR%"
endlocal
