# Requesting user input

## Positional Parameters

Call the script like that
myscript.sh a b c d e f g h i j k

```SHELL
#! /bin/bash
echo "My name is $1" # first argument (a)
echo "My home dir is $2" # second argument (b)
echo "The 10th argument is ${10}" #  (j)
echo "The 11th argument is $11" # (a1) <- expected k, but shell did $1 and then 1
```

## Special Parameters

Parameters that bash gives special meaning.
Value of a special parameter is calculated for us based on our current script.

- $#
- $0
- $*
- $@

```SHELL
$# # will give us the number of parameters the shell script received
```

```SHELL
$0 # when used in a script, $0 expands to the name of the script
```

```SHELL
$@ # access all positional parameters passed to the script
# it seperates all parameters with a space
# $1 $2 $3 ... $N

"$@" # "$1" "$2" "$3" ... "$N"
# takeaway: there is a difference between quoted and unquoted $@
```

```SHELL
$* # works with the IFS, it concats the positional parameters with the IFS value

# script myscript.sh
IFS=,
echo $*

# execute
myscript.sh 1 2 3 # -> 1,2,3
```

## Read Command

```SHELL
# in the command line (not as script)
echo $myvar # empty output, REPLY has no value
read # shell waits for input
hello
echo $myvar # hello
```

```SHELL
#! /bin/bash
read name age
echo "My name is $name"
echo "My age is $age"
```

```SHELL
#! /bin/bash
read -p "Input your name: " name 
read -p "Input your age: " age 
read -s -p "Input your password: " password 
echo "My name is $name"
echo "My age is $age"
echo "My password is $password"
```

## Select Command

The `PS1` variable was to control the prompt string. The `PS3` controls the promt string of the select command.

```SHELL
#! /bin/bash
PS3="What is the day of the week?: "
select day in mon tue wed thu fri sat sun; do
    echo "The day is $day"
    break
done
```
