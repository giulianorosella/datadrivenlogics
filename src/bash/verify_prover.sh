#!/bin/bash

echo "Execute verify_prover in: $(pwd)"

source src/bash/vars/vars_prover.sh

# Function to check if prover9 variable points to a valid directory
check_prover9() {
    if [ -z "$prover9" ]; then
        echo "Error: prover9 variable is empty. Please set the prover9 variable in vars_prover.sh."
        exit 1
    fi

    if [ ! -x "$prover9" ]; then
        echo "Error: prover9 executable does not exist or is not executable: $prover9"
        exit 1
    fi
}

#launch prover9 with test input
test_prover() {
    echo "Launching Prover9 with input test file: $input_test"
    echo "Output will be saved to: $output_test"
    "$prover9" -f "$input_test" >"$output_test" 2>&1
    echo "Check $output_test if input test correctly present with current date time. If so, prover9 is correctly working"
}
