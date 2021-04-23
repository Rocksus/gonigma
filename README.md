# Gonigma

A golang implementation of the Enigma machine.

## Overview

The Enigma machine is a cipher device used to transmit messages securely during the early to mid 20th century. It consists of several components, such as:

- Rotors
- Reflector
- Plugboard (only the army's enigma machines have these)

The amazing feat of the Enigma machine is its superiority compared to a basic caesar cipher.
For example, when you want to encrypt a sequence of letters:
`aaaaaaaaaaaaaaaaaaaa`

using a caesar cipher will result into:

`jjjjjjjjjjjjjjjjjjjj`

The Enigma machine uses the rotating rotors to create a shifting sequence, resulting in an encryption like:

`jnettodpzbibjjnettbl`

which is very hard to solve if we don't know the right configuration.

## Running the Example

You can run the example by running

```bash
go run cmd/main.go

```

## Improvements

There are many attempts in trying to crack the Enigma machine using brute force, thanks to our massive technological advancements. It might be beneficial to add them here and compare how they fare.
