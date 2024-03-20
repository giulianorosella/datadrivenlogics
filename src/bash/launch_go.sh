#!/bin/bash

filename="USER INPUT"

echo "Enter filename you desire to run:" 
echo
read -p "name: " filename
echo

filename="bin/"$filename

"$filename"

