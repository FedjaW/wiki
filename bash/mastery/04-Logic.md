# Logic

## Chaining commands with list operators

A list is when oyu put one or more commands on a given line.

List operators:

- ;
- &
- &&
- ||

& lets you run commands asynchronously.

```SHELL
echo 123 & echo 456
# [1] 74 -> echo 123 was send to the background
# 456
# 123
```

`;` let you run commands sequentialy.

```SHELL
echo 123 ; echo 456
# 123
# 456
```

```SHELL
# will be executed even if error appears
ls non_existing_dir ; echo 123
# ls: cannot access...
# 123
```

That brings us to the next operators to chain commands:

The `&&` (and) operator makes it so that the second command only runs if the first one was successful.

The `||` (or) operator makes it so that the second command only runs if the first one failed.

## Test Commands + Conditional Operators

A test command is written using a set of square brackets `[ ]` with a space between the brackets and its content.

```SHELL
# compare numbers
# -eq only comares integer
[ 2 -eq 2 ] ; echo $? # $? will output the exit status of the command, 0 means true

# compare strings
a=hello
b=bye
[[ $a = $b ]] ; echo $?

# -z checks if string is empty (check based on length)
[[ -z = $c ]] ; echo $?

# -n checks if string is not empty (check based on length)
[[ -n = $c ]] ; echo $?

# compare files
# -f checks is file is a file
# -d checks is input is a dir
# -x checks is file has execute permissions
# -e checks is file exists
[[ -e today.txt ]] ; echo $?
```

## If statements

"If" is a compound command [[02-HowBashProcessesCommandsLines.md]]

```SHELL
#! /bin/bash
if [ -1 -gt 2 ]; then
    echo "test passed"
elif [ 1 -eq 1 ]; then
    echo "test second test passed"
else 
    echo "test passed not"
fi
```

Chain conditions:

```SHELL
#! /bin/bash
# && (AND) operator

a=$(cat file1.txt)
b=$(cat file2.txt)
c=$(cat file3.txt)

if [  $a -eq $b ] && [ $a = $c ]; then
    echo "file 1 equals file 2 equals file 3"
    rm file2.txt file3.txt
else 
    echo "files do not match"
fi
```

```SHELL
#! /bin/bash
# || (OR) operator

a=$(cat file1.txt)
b=$(cat file2.txt)
c=$(cat file3.txt)

if [  $a -eq $b ] || [ $a = $c ]; then
    echo "file 1 equals file 2 OR equals file 3"
else 
    echo "files do not match"
fi
```

## Case statements

"Case" is a compound command [[02-HowBashProcessesCommandsLines.md]]

```SHELL
read -p "Please enter a number: " number 
# you must add the $ symbol and use the "". 
# It will safe you from bugs
case "$number" in 
    [0-9]) echo "you entered a single digit number";;
    [0-9][0-9]) echo "you entered a two digit number";;
    *) # default case
esac

# case works with globbing, [0-9] will match any number between 0 and 9
```
