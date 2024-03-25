# How Bash Processes Commands Lines

## Overview

Bash goes through a 5 step process to read the command line input:

- Step 1: Tokenisation
  - A sequence of characters that is considered as a single unit by the shell
  - Metacharacters: | & ; () <> space tab newline
  - Tokens will be classified into:
    - Words (a token that does not contain an unquoted metacharacter)
    - Operators (a token that contains at least one unquoted metacharacter)
      - Control operators
      - Redirection operators
- Step 2: Command Identification
  - Simple commands
    - are just a bunch of individual words, and each simple command is terminated by a control operator
  - Compound commands
    - provide bash with its programming constructs, such as if statements, for loops, while loops, etc.
- Step 3: Expansions
- Step 4: Quote Removal
- Step 5: Redirection

## Quoting

Quoting is about **Removing Special Meanings**.

There are three types of quoting in bash.

1. Backslash (\\)
   - Removes the special meaning from the next character
2. Single Quotes (' ')
   - Removes special meaning from all character inside
3. Double Quotes (" ")
   - Removes special meanignn from all except dollar signs ($) and backticks (`)  

## Step 1: Tokenisation

- The shell identifies unquoted metacharacters and uses them to divide up the command line into tokens
- It then characterises the tokens into words and operators

Two types of Operators:

- Control operators
  - Newline
  - |
  - ||
  - &
  - &&
  - ;
  - ;;
  - ;&
  - ;;&
  - |&
  - (
  - )

- Redirection Operators (without single quotes)
  - '<'
  - '>'
  - '<<'
  - '>>'
  - '<&'
  - '>|'
  - '<<-'
  - '<>'
  - '>&'

Control operators and Redirection operators only matter if they are unquoted!

### Example

```SHELL
# this is just a sequence of characters
echo $name > out.txt

# Tokenisation:
# - Identifying the metacharacters for this characters
#    Here the metacharacter are the spaces and the >
#    There are 4 metachacters
# - Break the command in Words and Operators
#    Words are here (3): echo, $name, out.txt
#    Will look for an unquoted metachacter and there for an Operator, here: >
# The result are tokens that are divided into words and operators.
```

## Step 2: Command Identification

todo