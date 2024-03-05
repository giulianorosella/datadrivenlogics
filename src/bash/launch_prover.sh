#!/bin/bash

echo "Execute launch_prover in: $(pwd)"

source src/bash/vars/vars_prover.sh

prove() {
    local input="$1"
    local output="$2"
    echo "Launching Prover9 with input file: $input"
    echo "Output will be saved to: $output"
    "$prover9" -f $input >$output 2>&1
    echo "done"
}

prove $input_path $output_path
