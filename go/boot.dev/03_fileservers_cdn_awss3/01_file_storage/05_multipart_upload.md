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

> Q: Now, regarding `10 << 20` in your lesson, can you figure out how that translates to multiplying 10 by a power of 2? What resulting number are we left with?

`10485760` is the answer. Here's how you got there:

`10 << 20` means shifting `10` (in binary: `1010`) to the left by 20 places. Each shift to the left multiplies the number by 2. So the result is:

`10 * 2^20 = 10 * 1,048,576 = 10,485,760`

This gives us 10 megabytes (since a megabyte is `1,048,576` bytes, or `2^20` bytes)!

> Q: isn't mega (M) `10^6` ?

In the world of computers and programming, we typically deal with binary definitions for sizes, and here "mega" (as in 1 Megabyte) is usually `2^20`, or `1,048,576`. This binary definition stems from how computers work internally using powers of 2.

However, in general science or everyday measurements, "mega" is often defined as the decimal value `10^6` (1,000,000). That’s why there’s sometimes this discrepancy.

For something like `10 << 20`, we're working in the binary world, so we use the binary version, hence why `10 << 20` represents 10 megabytes in a programming context.
