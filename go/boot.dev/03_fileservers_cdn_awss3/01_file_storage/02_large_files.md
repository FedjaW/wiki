# Large Files

So, what are “large files” anyway?

You’re probably already familiar with small structured data; the stuff that’s usually stored in a relational database like Postgres or MySQL. I’m talking about simple, primitive data types like:

- `user_id` (integer)
- `is_active` (boolean)
- `email` (string)

Large files, or “large assets”, on the other hand, are giant blobs of data encoded in a specific file format and measured in kilo, mega, or gigabytes. As a simple rule:

- If it makes sense to go into an excel spreadsheet, it probably belongs in a traditional database
- If it would normally be stored on your hard drive as its own file, it probably is a “large file”

Large files are interesting because:

- They’re large in size (duh) and are thus more performance-sensitive
- They’re usually accessed frequently, and this combined with their size can quickly lead to performance bottlenecks

## Assignment

1. In the root of your repo there is a script called `samplesdownload.sh`. Run it from the root of the repo to download some sample images and videos into the `samples` directory:

```cli
./samplesdownload.sh
```

2. Take a look at the `boots-image-horizontal.png` file in the `samples` directory: it’s a PNG image file. You can open it in an image viewer to see what it looks like.
3. Use `xxd` to view a hexdump of the `samples/boots-image-horizontal.png` file:

```cli
xxd <file>
```

`xxd` converts the binary content of the file into a human-readable hexadecimal and ASCII formats. You should see a bunch of gobbledegook - that’s what a PNG file looks like in its raw form.

4. Inspect the first 8 bytes of the file more closely. Use `xxd` with `-l` (length) to limit the output:

```cli
xxd -l 8 <file>
```

You’ll see that the first 8 bytes are 89 50 4e 47 0d 0a 1a 0a, which is the PNG file signature. It tells the reader that this is a PNG file, in fact, the characters PNG are present in bytes 2-4.
