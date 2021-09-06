# gnark-bench

Command line tool to benchmark and profile execution of `plonk` and `groth16` zkSNARKs with `gnark`. 

If running benchmarks on a `amd64` machine with `adx` instructions, be sure to build `gnark-bench` with `-tags=amd64_adx`. 

## Usage

`go build [-tags=amd64_adx]`

`./gnark-bench` will output self explanatory help:

```
runs benchmarks and profiles using gnark

Usage:
  gnark-bench [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  groth16     runs benchmarks and profiles using Groth16 proof system
  help        Help about any command
  plonk       runs benchmarks and profiles using PlonK proof system

Flags:
      --algo string      name of the algorithm to benchmark. must be compile, setup, prove or verify (default "prove")
      --circuit string   name of the circuit to use (default "expo")
      --count int        bench count (time is averaged on number of executions) (default 2)
      --curve string     curve name. must be [bn254 bls12_377 bls12_381 bw6_761 bls24_315 bw6_633] (default "bn254")
  -h, --help             help for gnark-bench
      --profile string   type of profile. must be none, trace, cpu or mem (default "none")
      --size int         size of the circuit, parameter to circuit constructor (default 10000)

```


## Notes
* It is possible to benchmark a specific circuit, need to fork and extend `circuit` subpackage with a struct that implements `BenchCircuit` and register it in the global map. 
