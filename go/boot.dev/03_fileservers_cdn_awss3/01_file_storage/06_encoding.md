# Encoding

As you know, we use a SQLite database to power the majority of the web app. SQLite is a traditional relational database that works out of a single flat file, meaning it doesn’t need a separate server process to run.

Let’s talk about the elephant in the room: Our current solution for video thumbnails (storing the media in-memory) is a terrible solution. If the server is restarted, all the thumnbnails are lost!

But we can’t store an image in a SQLite column?.. Right?

To do so, we can actually encode the image as a base64 string and shove the whole thing into a `text` column in SQLite. Base64 is just a way to encode binary (raw) data as text. It’s not the most efficient way to do it, but it will work for now.

## Assignment

Update the code to store the image data in the thumbnail_url column in the database.

1. Use `base64.StdEncoding.EncodeToString` from the `encoding/base64` package to convert the image data to a base64 string.
2. Create a data URL with the media type and base64 encoded image data. The format is:

```txt
data:<media-type>;base64,<data>
```

3. Store the URL in the `thumbnail_url` column in the database.
4. Because the `thumbnail_url` has all the data we need, delete the global thumbnail map and the `GET` route for thumbnails.
5. Restart the server and re-upload the `boots-image-horizontal.png` thumbnail image to ensure it’s working.
