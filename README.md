# Set environment variables from Terraform output

This repository contains the code for a unix shell helper command that sets
environment variables from a terraform project.

## Usage

`source <(terraform output -json | tfout2env)`

### create a .env shell file

`terraform output -json | tfout2env > .env`
