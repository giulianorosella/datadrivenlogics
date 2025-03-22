# DataDrivenLogics

This project aims to develop quantitative methods to investigate logics and their relationships.

## Running Prover9 from the Terminal:

```
prover9/bin/LADR-2009-11A/bin/prover9 -f prover9/conf/default_input.in prover9/inputs_in/CHOSENINPUT > prover9/outputs_out/CHOSENNAMEOUTPUTFILE
```

```
prover9/bin/LADR-2009-11A/bin/mace4 -c -f prover9/conf/default_input.in prover9/inputs_in/CHOSENINPUT > prover9/outputs_out/CHOSENNAMEOUTPUTFILE
```



## Go Setup Instructions:

- Navigate to the development directory:
  ```bash
  cd src
  ```
- Ensure everything is okay with the dependencies:
```bash
  go mod tidy
  ```
- Build your binary:
```bash
  go build -o ../bin/FILENAME main.go
  ```
- Navigate to the bin directory and run the executable:
```bash
 cd ../bin
./FILENAME
  ```
