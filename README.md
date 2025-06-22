# Sample bit reducer

Used for downsampling wav files to 16 bit so that they're compatible with the MPC 1000. 

## Installation

Install with `go install github.com/pparrott/sample_bit_reducer/cmd/sample_bit_reducer@latest`

## How to use

To use, run:
`sample_bit_reducer -path "folder path"`

The program recursively searches all nested directories and converts any wav files above 16 bit down to 16 bit. 