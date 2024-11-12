# Vim

> 1 Exit vim with `ZZ` WITH saving changes and `ZQ` WITHOUT saving changes

> 2 Select () and {} with `vib/viB`

```txt
to select the content within the (): `vib`
to select the content within the {} `viB`

Dummy Text for testing:
("/api/route/1", 200, { "message": "Hello"})
```

> 3 Edit multiple lines at once

```txt
1. to put text to the beginning of all lines:
ctrl+v to get into visual block mode
select vertically
insert mode with I
then type and press Enter
all changes are applied at once to the beginning

2. to add to the end of all lines:
ctrl+v to get into visual block mode
select vertically
press $ to go to the end of each line
then type and press Enter
all changes are applied at once to the end


Dummy Text for testing:
Hello World
Bye Bye Germany
What a great day
```

> 4 Highlight last highlighted text with `gv`

> 5 Toggle case with `~` and `g~`

> 6 Re-indent the whole file with `gg=G`

> 7 Jump between (), [], {}, etc. with `%`

> 8 Put Vim to the background with `Ctrl-z` and return to Vim with typing `fg` in the console

> 9 Open URL in browser with typing `gx` when cursor is under an URL

> 10 Open file in folder with typing `gf` when cursor is under an file path

> 11 Mark a location and return to it by setting a mark with `ma`, a is here any letter. Then press `'a` to return to the marked position. When typing a uppercase `A` (or any letter) you can jump between files to the mark. You can set multiple marks at a time.

> 12 Jump to a specific line like a pro with `42G`. Will jump to line 42

> 14 Join two lines with `J` (with space) and `gJ` (without space) when cursor placed on the upper line

```text
hallo
welt

J -> hallo welt
gJ -> hallowelt
```
