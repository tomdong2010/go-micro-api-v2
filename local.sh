#!/bin/sh

go build -o $1.exe app/$1/main.go

echo "build finish!!!"

./$1.exe $2 $3 $4 $5 $6 $7 $8 $9
