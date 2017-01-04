#!/bin/bash
if [ $# -eq 0 ]
  then echo "Please input time."
else
  echo $1 | grep -q '^[0-2][0-3]:[0-5][0-9]:[0-5][0-9]$'
  if [ $? -eq 0 ]
    then echo "ok"
    date -s $1
  else
    echo "Invalid time format! ex: demo.sh <hh:mm:ss>"
  fi
fi
