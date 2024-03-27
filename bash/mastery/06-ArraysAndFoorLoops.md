# Arrays and For Loops

## Indexed Arrays

```SHELL
#! /bin/bash
numbers=(1 2 3 4)

echo "$numbers" # 1 (will give you the index 0 by default)
echo "${numbers[2]}" # 3 (will give you the index 2)
echo "${numbers[@]}" # 1 2 3 4 (@ same as in positional parameter expansion)
echo "${numbers[@]:1}" # 2 3 4 (can use all things from parameter expansion)
echo "${numbers[@]:1:2}" # 2 3 (can use all things from parameter expansion)

# add a number to the array
numbers+=(5)
echo "${numbers[@]}" # 1 2 3 4 5

# remove a number to the array
unset numbers[2]
echo "${numbers[@]}" # 1 2 4 5

# replace a value
numbers[0]=a
echo "${numbers[@]}" # a 2 4 5
```

## The readarray command

```SHELL
#! /bin/bash
# each line in the file will be an element in the array
# -t is good for not having the \n on every element
readarray -t days < days.txt
echo "${days[@]}" 
echo "${days[@]@Q}" # @Q shows also invisible characters like newlines
# without -t flag => $'Monday\n' ...
# with-t flag => 'Monday' ...
```

## For Loops

```SHELL
#! /bin/bash
for element in first secon third; do
    echo "This is $element"
done
```

```SHELL
#! /bin/bash
readarray -t files < files.txt
for file in "$files[@]"; do
    if [ -f "$file" ]; then
        echo "$file already exists"
    else
        touch "$file" 
        echo "$file was created"
    fi
done
```
