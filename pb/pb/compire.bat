@echo off
setlocal enabledelayedexpansion

set "output_file=all.proto"

(
    echo syntax = "proto3";
    echo option go_package =".";
    echo message pb {
) > %output_file%

for %%f in (*.proto) do (
    if "%%f" neq "%output_file%" (
        more +1 "%%f" >> %output_file%
    )
)

echo } >> %output_file%

protoc --go_out=../msg all.proto
pause