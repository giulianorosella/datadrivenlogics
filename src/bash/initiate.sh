#!/bin/bash

echo "Execute initiate.sh in: $(pwd)"
# usefull vars
requirements_py="src/python/requirements.txt"
python_script="src/python/betterinputs.py"

# Function to activate virtual environment
activate_venv() {
    # Check if virtual environment exists and activate it
    if [ -d "venv" ]; then
        source venv/bin/activate
    else
        echo "Virtual environment not found. Creating a new one..."
        python -m venv venv
        source venv/bin/activate
    fi
}

# Activate virtual environment
activate_venv

# Install dependencies from requirements.txt
# only i fnot already there
if [ -f "$requirements_py" ]; then
    if ! pip show -q -r "$requirements_py"; then
        echo "Installing dependencies from $requirements_py..."
        pip install -r "$requirements_py"
    else
        echo "Dependencies are already installed."
    fi
else
    echo "No $requirements_py file found."
fi

# Check if the Python script exists
if [ ! -f "$python_script" ]; then
    echo "Error: Python script '$python_script' not found."
    exit 1
fi

# Run the Python script
python "$python_script"

# Check if the Python script exited successfully
if [ $? -ne 0 ]; then
    echo "Error: Python script '$python_script' failed to execute."
    exit 1
fi

# Deactivate virtual environment
deactivate
