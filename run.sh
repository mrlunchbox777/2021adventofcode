#!/bin/bash

list_of_days=$(ls ./src)
day_count=$(echo "$list_of_days" | wc -l)
day_choice=$day_count

read -p "use latest?" use_latest

if [[ "$use_latest" =~ ^[nN]$ ]]; then
	read -p "which day, max $day_count?" day_choice
fi

if [[ "$day_choice" =~ ^[^0-9]*$ ]] || [ "$day_choice" -gt "$day_count" ] || [ "$day_choice" -le "0" ]; then
	echo "bad day choice.." >&2
	exit 1
fi

echo "running day $day_choice"
old_dir=$(pwd)

{
	cd "./src/day$day_choice"
	go build -o run.o run.go
	./run.o
	rm run.o
} || {
	error="$?"
	failed="true"
}

cd "$old_dir"

if [[ "$failed" == "true" ]]; then
	echo "Failure detected, check logs, exiting...">&2
	echo "exception code - $exception">&2
	exit $exception
fi
