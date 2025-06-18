#!/bin/bash

echo "build web interface..."
cd ./web_backend
go build
./led_game &
RUNNING_PID=$!
cd ..
echo "compile and run arduino sketch"
arduino-cli compile -b arduino:avr:uno guess_game.ino
arduino-cli upload -b arduino:avr:uno -p /dev/ttyUSB0 .

if [ ! -d "utils/.venv" ];
then
   echo "create .venv for python utils"
   python -m venv ./utils/.venv
   echo "enter .venv"
   source ./utils/.venv/bin/activate
   pip install --no-cache-dir -r ./utils/requirements.txt
else
   echo "enter .venv"
   source ./utils/.venv/bin/activate
fi

python ./utils/update_res.py 
echo "killing web server..."
kill ${RUNNING_PID}
