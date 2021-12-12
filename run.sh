#!/bin/bash

list_of_days=$(ls ./src)
day_count=$(echo "$list_of_days" | wc -l)
day_choice=$day_count
day_override=$DAY_OVERRIDE
part_override=$PART_OVERRIDE

if [ -z "$day_override" ]; then
	read -p "use latest day?" use_latest
	if [[ "$use_latest" =~ ^[nN]$ ]]; then
		read -p "which day, max $day_count?" day_choice
	fi
else
	use_latest="n"
	day_choice="$day_override"
fi

if [[ "$day_choice" =~ ^[^0-9]*$ ]] || [ "$day_choice" -gt "$day_count" ] || [ "$day_choice" -le "0" ]; then
	echo "bad day choice.." >&2
	exit 1
fi

list_of_parts=$(ls "./src/day$day_choice" | grep '^part\-[0-9]*$')
part_count=$(echo "$list_of_parts" | wc -l)
part_choice=$part_count
if [ -z "$part_override" ]; then
	read -p "use latest part?" use_latest
	if [[ "$use_latest" =~ ^[nN]$ ]]; then
		read -p "which part, max $part_count?" part_choice
	fi
else
	use_latest="n"
	part_choice="$part_override"
fi

if [[ "$part_choice" =~ ^[^0-9]*$ ]] || [ "$part_choice" -gt "$part_count" ] || [ "$part_choice" -le "0" ]; then
	echo "bad part choice.." >&2
	exit 1
fi

echo "running day $day_choice, part $part_choice"
old_dir=$(pwd)

{
	cd "./src/day$day_choice/part-$part_choice"
	go build -o run.o .
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
