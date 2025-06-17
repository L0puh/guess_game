#!/bin/bash

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
