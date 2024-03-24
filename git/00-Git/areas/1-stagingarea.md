# Staging area

- What files are going to be part of the next commit
- The staging area is how git knows what will change between the current and the next commit
--
- A "clean" staging are isn't empty
- Consider the baseline staging area as being an exact copy of the latest commit
- The staging area contains a list of files that where in your last commit and the sha1's to that files
- The staging area (also called index) is a binary
--
- When you change a file git detects this by knowing that the sha1 of the file has changes compared to the file in the repo

To inspect the index

```
git ls-files -s

Mode   SHA1                                         filename
100644 b512c09d476623ff4bf8d0d63c29b784925dbdf8 0	.gitignore
100644 fa51da29e7fabbc7d3e3f9a076bf20ba06b0fbe1 0	.prettierrc
100644 a55da3cfb38885448bcd8e9069cf8ed5d3e7b159 0	README.md
100644 79ce8e979c9e2381814212c1bdce9f8a2de61486 0	README_DEV.md
100644 b819c602060d56fa62d07fc84387c39800c2e0b8 0	buzzwords.js
100644 140dd43d0e815250b9390defc7d622892cd1b78f 0	index.js
100644 c3c3500a67dc18a5dd59a8b83a7f84488b45310b 0	package-lock.json
100644 55d24a96d57aa6ee3f225241dc3b6df784e93c91 0	package.json
```