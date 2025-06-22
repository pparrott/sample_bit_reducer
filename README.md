# Sample bit reducer

Used for downsampling wav files to 16 bit so that they're compatible with the MPC 1000. 

To use:
`sample_bit_reducer -path "path to file"`

The program recursively searches all nested directories and converts any wav files above 16 bit down to 16 bit. 