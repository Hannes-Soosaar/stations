## Overview

This program is designed to find the most efficient paths for trains to travel from a start station to an end station on a given network map. It is implemented in Go and can be executed via the command line interface (CLI). The program also includes built-in tests to ensure functionality and correctness.

## Features

- Efficient Pathfinding: Determines the most efficient routes for trains between specified stations.
- CLI Execution: Easily run the program from the command line.
- Automated Testing: Includes a Makefile for running tests to verify the program's correctness.
- Error Handling: If an error occurs, detailed information will be output to the console.

## Prerequisites

- Go programming language installed
- Make utility installed

## Running the program

To execute the program, use the following command in your terminal:

```go run . [path to file containing network map] [start station] [end station] [number of trains]```

Example:

```go run . beginning-terminus.map beginning terminus 20```

## Running tests

The program includes built-in tests to verify its functionality.

Run ```make init``` in the parent directory

Run all the tests with ```make run```