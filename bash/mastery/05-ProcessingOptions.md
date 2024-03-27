# Processing Options & Reading Files

## While loops

```SHELL
#! /bin/bash
read -p "Please enter a number: " num
while [ $num -gt 10 ]; do
    echo "$num"
    num=$(( $num - 1 ))
done
```

## Handling Command Line Options

```SHELL
#! /bin/bash
# cal farenheit from celcius
while getopts "f:c:"  opt; do
    case "$opt" in
        c) result=$(echo "scale=2; ($OPTARG * (9 / 5)) + 32" | bc);;
        f) result=$(echo "scale=2; ($OPTARG - 32") * (5/9) | bc);;
        \?);; # interprete ? literaly by escaping, getops will put a ? if no value is provided
    esac
done

# ./myscript.sh -f 32
# ./myscript.sh -c 0
```

## Read-while loops

We have the following scenario:

```SHELL
cat file1.txt
This is line 1
This is line 2
This is line 3
```

```SHELL
#! /bin/bash
while read line; do
    echo "$line"
done < "$1" # read the file that i refere to in the first positional parameter

# < redirection from standard input

# ./myscript.sh file1.txt
```

Using process substitution

```SHELL
#! /bin/bash
while read line; do
    echo "$line"
done < <(ls $HOME) # process substitution

# < redirection from standard input

# ./myscript.sh file1.txt
```
