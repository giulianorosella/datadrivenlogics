#!/bin/bash

echo "Execute launch_prover in: $(pwd)"

source src/bash/vars/vars_prover.sh
source src/bash/verify_prover.sh

prove() {
    local input="$1"
    local output="$2"
    echo "Launching Prover9 with input file: $input"
    echo "Output will be saved to: $output"
    "$prover9" -f $input >$output 2>&1
    echo "done"
}

# Call prover9 modules only if test passes
verify_and_launch() {
    check_prover9
    echo

    test_prover
    echo

    if [ $? -eq 0 ]; then
        echo "Prover9 verification successful. Launching Prover9."
        # Execute the launch.sh script
        prove $input_path $output_path
        
    else
        echo "Error: Prover9 verification failed. Aborting."
        exit 1
    fi
}

verify_and_launch

