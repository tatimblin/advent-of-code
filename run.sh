#!/bin/bash

test=false
year=0
day=0

while [[ "$#" -gt 0 ]]; do
  case $1 in
    -y|--year)
      year="$2"
      shift
      ;;
    -d|--day)
      day="$2"
      shift
      ;;
    -t|--test)
      test=true
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
  shift
done

if [ -z "$year" ]; then
  echo "Usage: $0 -y <Year> -d <Day> [-t]"
  exit 1
fi

if [ -z "$day" ]; then
  echo "Usage: $0 -y <Year> -d <Day> [-t]"
  exit 1
fi

if ((day < 10)); then
  day="0$day"
fi

cd "$year/$day"

if [ "$test" = true ]; then
  go test
else
  go run main.go
fi