# General topics about BASH

## what is a shell?

"A `program` that `interprets` the commands you type in your terminal and `passes` them on to the `operating system`" - unknown.

"The `purpose` of the shell is to make it more `convenient` for ypu to `issue commands` to your computer". - unknown.

## what it bash?

BASH - Bourne Again SHell

**Shebang #!**

`#!/bin/bash` is a comment read by the shell and tells the os where the path to the binary is that must execute the script.

in fact `#!/bin/bash` is not the best way to define bash as the interpreter of the script, because on some system bash isn't at that given path.

a better way to specify bash is to use `#!/usr/bin/env bash`. with this the system will look up the correct path of bash.

run a file command on a file to see the difference:
`file myscript.sh`
output:
`myscript.sh: Bourne-Again shell script, ASCII text executable`

## Permissions

To give the scipt execute permissions run:
`chmod +x myscript.sh` or  `chmod 744 myscript.sh`

`744` is a `octonumber`

see: https://chmod-calculator.com/

Definition:
`-rwxrwxrwx`

`owner|group|other`
`r: read`
`w: write`
`x: execute`

## PATH

In Linux the systems path is a variable that tell the shell which directories to search for executable files in the directories defined in its path variable.

`echo $PATH` to inspect the path variable.

One way to add a script to the path variable:

1. open `vim ~/.profile`
2. add `export PATH="$PATH:$HOME/myscripts"`

## Types of shell parameters

- Variables
  - User defined variables
  - Shell variables
    - Bourne Shell Variables
        - <span style="color:brown">Bourne</span> shell (created by Stephen Bourne 1979)
        - <span style="color:brown">10</span>  in total
    - Bash Shell Variables
        - <span style="color:brown">Bash</span> shell (based on the Bourne shell)
        - <span style="color:brown">B</span> ourne <span style="color:brown">A</span>gain <span style="color:brown">SH</span>ell (BASH)
        - Around <span style="color:brown">95</span> in total
- Positional parameters
- Special parameters

### User defined variables

```SHELL
# User defined variable
student="Sarah" # no space around the equals sign, otherwise bash will think it is a command.
echo "Hello ${student}" # This is called paramater expansion
```

### Shell variables

```SHELL
# The PATH variable contains the list of folders 
# that the shell will search for executable files to run as command names
$PATH # A Bourne Shell variable
echo ${PATH}
```

```SHELL
# The HOME variable is used to store 
# the absolute path to the current users's home directory
$HOME # A Bourne Shell variable
echo ${HOME}
```

```SHELL
# The USER variable containts the username of the current user
$USER # A Bourne Shell variable
echo ${USER}
```

```SHELL
# The PS1 variable containts the promt string
# shown in the terminal before each command
$PS1 # A Bourne Shell variable
echo ${PS1}
```

### Parameter expansion

```SHELL
name=ZiYaD
# parameter expansion
echo ${name} # ZiYaD
echo ${name,} # ziYaD (lower case first letter)
echo ${name,,} # ziyad (all lower case)
echo ${name^} # ZiYaD (upper case first letter)
echo ${#name} # 5 (number of letters)

numbers=0123456789
# ${parameter:offset:length}
${numbers:0:7} # 0123456
${numbers:3} # 3456789
```

### Command Substitution

```SHELL
# $(command)
# the shell will run the command and substitute the $(command) with the standard output of the command.
time=$(date +%H:%m:%S)
echo "Hello $USER, the time right now is $time"
```

### Arithmetic Expansion

```SHELL
# $((expression))
echo $(( 3 + 4)) # 7
```

### Tilde Expansion

```SHELL
# ~ expands to the home directory of the user
echo ~ # /home/hans

# every system has the user root
# which is the account of the system administrator
echo ~root  # /root

# if the user does not exist
echo ~abc  # ~abc
```

## Inspect processes

execute: `zsh`\
execute: `ps -o pid,ppid= -C zsh`\
execute: `zsh`\
execute: `ps -o pid,ppid= -C zsh`\
execute: `exit`\
execute: `ps -o pid,ppid= -C zsh`\
execute: `exit`\
execute: `ps -o pid,ppid= -C zsh`
