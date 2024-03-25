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
   - Removes special meaning from all except dollar signs ($) and backticks (`)  

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

There are two types of commands:

1. Simple commands (most common)
2. Compound commands
    - bash programming constructs like for loop..

**Simple command:**

The first word of a simple command is interpreted as command name.
Following words are interpreted as arguments.

```SHELL
echo 1 2 3
# Command name: echo
# Individual Arguments: 1 2 3
```

The end of a command is given by a control operator
In the upper case the end of the command is defined by the invisible command newline at the end.

```SHELL
echo 1 2 3; echo 1 2 3
# Here the end of a command is defined by ; and for the second part by a newline
# There are two commands here.
```

## Step 3: Expansion

There are 4 stages of shell expansions (executed in the following order):

  1. Brace expansion
  2. Expansions
    - Parameter Expansion
    - Arithmetic Expansion
    - Command substitution
    - Tilde Expansion
  3. Word Splitting
  4. Globbing

**Word splitting:**

Word splitting is only performed on the results of unquoted:

- Parameter expansions
- Command substitutions
- Arithmetic expansions

The character used to split words are governed by the IFS (Internal Field Seperator) variable.

- Space, Tab and newline (by default)

```SHELL
# Inspect the IFS variable
echo "${IFS@Q}"
```

Example:

```SHELL
numbers="1 2 3 4 5"
touch $numbers # will create 5 folders, not 1 folder with name '1 2 3 4 5'
touch "$numbers" # will create 1 folder with name '1 2 3 4 5'
# Again: Word splitting is only performed on the results of unquoted:
# - Parameter expansions
# - Command substitutions
# - Arithmetic expansions
```

**Globbing:**

- Globbing is only performed on words (not operators)
- Globbing Patterns are words that contain unquoted Special Pattern Characters:
  - \*
  - ?
  - []

```SHELL
ls *.txt # (* match anything) fileABC.txt
ls file?.txt # (? matches exactly one character) fileX.txt
ls file[ab].txt # ([] matches any file has either an a or b in the placeholder) filea.txt fileb.txt
```

## Step 4: Quote Removal

During quote removal, the shell removes all unquoted backslashes, single quote characters, and double quote characters that did NOT result from a shell expansion.

```SHELL
echo \$HOME # $HOME
# The backslash was removed during quote removal
```

## Step 5: Redirection

Data Streams:

- Stream 0 = Standard Input (stdin)
- Stream 1 = Standard Output (stdout)
- Stream 2 = Standard Error (stderr)

```SHELL
1> # redirects standard output
```

```SHELL
>> # appends data
```

```SHELL
2> # redirects standard error
```

```SHELL
&> # will connect the standard output and the standard error datastreams to the same place
```

```SHELL
echo hello &> /dev/null # /dev/null is the bit bucket, data gets deleted
```
