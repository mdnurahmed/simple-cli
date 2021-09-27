# simple-cli

A simple cli written in golang that accepts a string and does the following to it:

- converts the string to uppercase and outputs it to stdout.
- converts the string to alternate upper and lower case and outputs it to stdout.
- creates a CSV file from the string by making each character a column in the CSV and then output
  "CSV created!" to stdout.

# How to Run

## Using GO

```
git clone https://github.com/mdnurahmed/simple-cli
cd simple-cli
go install .
cli
```

press ctrl+C to exit the CLI app .

To run unit tests -

```
go test ./...
```

## Using docker

```
git clone https://github.com/mdnurahmed/simple-cli
cd simple-cli
docker-compose up
docker run -it simple-cli_simple-cli /bin/sh
cli
```

press ctrl+C to exit the CLI app .
To read the created csv file -

```
cat csv.csv
```

To run unit tests -

```
go test ./...
```

Type in 'exit' to exit to the terminal.
