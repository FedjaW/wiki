# Multipart Uploads

So you might already be familiar with simple JSON/HTML form POST requests. That works great for small structured data (small strings, integers, etc.), but what about large files?

We don't typically send massive files as single JSON payloads or forms. Instead, we use a different encoding format called `multipart/form-data`. In a nutshell, it's a way to send multiple pieces of data in a single request and is commonly used for file uploads. It's the "default" way to send files to a server from an HTML form.

Luckily, the Go standard library's `net/http` package has built-in support for parsing `multipart/form-data` requests. The http.Request struct has a method called ParseMultipartForm.

```go
func (cfg *apiConfig) handlerUploadThumbnail(w http.ResponseWriter, r *http.Request) {
    // validate the request

	const maxMemory = 10 << 20 // 10 MB
	r.ParseMultipartForm(maxMemory)

	// "thumbnail" should match the HTML form input name
	file, header, err := r.FormFile("thumbnail")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse form file", err)
		return
	}
	defer file.Close()

	// `file` is an `io.Reader` that we can read from to get the image data
```

## Assignment

The handler for uploading thumbnails is currently a no-op. Let’s get it working. We’re going to keep it simple and store all image data in-memory.

- Notice that in `main.go` there is a global map of video IDs to thumbnail structs called `videoThumbnails`. This is where we’re going to store the thumbnail data.
- Notice the `handlerThumbnailGet` function. It serves the thumbnail file back to the UI, but it assumes that images exist in the `videoThumbnails` map (which they don’t yet!)

### Complete the handlerUploadThumbnail function. It handles a multipart form upload of a thumbnail image and stores it in the videoThumbnails map:

1. Authentication has already been taken care of for you, and the video’s ID has been parsed from the URL path.
2. Parse the form data
   1. Set a const `maxMemory` to 10MB. I just bit-shifted the number `10` to the left `20` times to get an `int` that stores the proper number of bytes.
   2. Use `(http.Request).ParseMultipartForm` with the `maxMemory` const as an argument

---

> Bit shifting is a way to multiply by powers of 2. `10 << 20` is the same as `10 x 1024 x 1024`, which is 10MB.

For example, in Go, `<<` means "shift left," and `>>` means "shift right." Shifting left is equivalent to multiplying by powers of 2, while shifting right is like dividing by powers of 2.

Take this example:

```go
x := 1        // Binary: 00000001
y := x << 3   // Shift left by 3 places: 00001000
fmt.Println(y) // Output: 8 (since 1 * 2^3 = 8)
```

Shifting left (<<) adds zeros to the right of the binary number, effectively multiplying by powers of two. In the example, shifting left by 3 multiplies 1 by 2^3.

Conversely:

```go
x := 16       // Binary: 00010000
y := x >> 2   // Shift right by 2 places: 00000100
fmt.Println(y) // Output: 4 (since 16 / 2^2 = 4)
```

Shifting right `>>` removes bits from the right, effectively dividing by powers of two.

3. Get the image data from the form

   1. Use `r.FormFile` to get the file data. The key the web browser is using is called “thumbnail”
   2. Get the media type from the file’s `Content-Type` header

4. Read all the image data into a byte slice using `io.ReadAll`
5. Get the video’s metadata from the SQLite database. The `apiConfig`’s `db` has a `GetVideo` method you can use

   - If the authenticated user is not the video owner, return a `http.StatusUnauthorized` response

6. Save the thumbnail to the global map

   1. Create a new `thumbnail` struct with the image data and media type
   2. Add the thumbnail to the global map, using the video’s ID as the key

7. Update the database so that the existing video record has a new thumbnail URL by using the `cfg.db.UpdateVideo` function. The thumbnail URL should have this format:

```txt
http://localhost:<port>/api/thumbnails/{videoID}
```

This will all work because the `/api/thumbnails/{videoID}` endpoint serves thumbnails from that global map.

8. Respond with updated JSON of the video’s metadata. Use the provided respondwithJSON function and pass it the updated 1database.Video1 struct to marshal.
9. Test your handler manually by using the Tubely UI to upload the `boots-image-horizontal.png` image. You should see the thumbnail update in the UI!

## Q & A

> Q: Now, regarding `10 << 20` in your lesson, can you figure out how that translates to multiplying 10 by a power of 2? What resulting number are we left with?

`10485760` is the answer. Here's how you got there:

`10 << 20` means shifting `10` (in binary: `1010`) to the left by 20 places. Each shift to the left multiplies the number by 2. So the result is:

`10 * 2^20 = 10 * 1,048,576 = 10,485,760`

This gives us 10 megabytes (since a megabyte is `1,048,576` bytes, or `2^20` bytes)!

> Q: isn't mega (M) `10^6` ?

In the world of computers and programming, we typically deal with binary definitions for sizes, and here "mega" (as in 1 Megabyte) is usually `2^20`, or `1,048,576`. This binary definition stems from how computers work internally using powers of 2.

However, in general science or everyday measurements, "mega" is often defined as the decimal value `10^6` (1,000,000). That’s why there’s sometimes this discrepancy.

For something like `10 << 20`, we're working in the binary world, so we use the binary version, hence why `10 << 20` represents 10 megabytes in a programming context.
