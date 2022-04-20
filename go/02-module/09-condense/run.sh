#!/bin/bash

go run condense.go < input.txt > input.dot
dot -Tpng input.dot -o graph.png