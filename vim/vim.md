# Vim

> 1 Exit vim with `ZZ / ZQ`

Exit vim WITH saving changes

```vim
ZZ
```

Exit vim WITHOUT saving changes

```vim
ZQ
```

> 2 Select () and {} with `vib/viB`

```
to select the content within the (): `vib`
to select the content within the {} `viB`

Dummy Text for testing:
("/api/route/1", 200, { "message": "Hello"})
```

> 3 Edit multiple lines at once

```
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

> 9 Open URL in browser with typing `gx` when cursor is under an URL, same works for file paths