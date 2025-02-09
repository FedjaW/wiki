# Welcome

Welcome to large file storage! Building a (good) web application almost always involves handling “large” files of some kind - whether its static images and videos for a marketing site, or user generated content like profile pictures and video uploads, it always seems to come up.

In this course we’ll cover strategies for handling files that are kilobytes, megabytes, or even gigabytes in size, as opposed to the small structured data that you might store in a traditional database (integers, booleans, and simple strings).

## Learning Goals

Understand what “large” files are and how they differ from “small” structured data
Build an app that uses AWS S3 and Go to store and serve assets
Learn how to manage files on a “normal” (non-s3) filesystem based application
Learn how to store and serve assets at scale using serverless solutions, like AWS S3
Learn how to stream video and to keep data usage low and improve performance

## AWS Account Required

This course will require an AWS account. We will not go outside of the free tier, so if you do everything properly you shouldn’t be charged. That said, you will need to have a credit card on file, and if you do something wrong you could be charged, so just be careful and understand the risk.

We recommend deleting all the resources that you create when you’re done with the course to avoid any charges. We’ll remind you at the end.

## Tubely

In this course we’ll be building “Tubely”, a SaaS product that helps YouTubers manage their video assets. It allows users to upload, store, serve, add metadata to, and version their video files. It will also allow them to manage thumbnails, titles, and other video metadata.

## Assignment

1. You’ll need both the Go toolchain (version 1.23+) and the Boot.dev CLI installed. If you don’t already have them, here are the installation instructions (https://github.com/bootdotdev/bootdev).
2. Fork the starter repo for this course into your own GitHub namespace, then clone your fork onto your local machine.
3. Copy the `.env.example` file to `.env`. In the future we’ll edit some of the `.env` values to match your configuration, but for now you can leave them as is.

```cli
cp .env.example .env
```

4. Run the Tubely server:

```cli
go run .
```

A URL will be logged to the console, open the URL in a browser to see the Tubely app. The webpage should load, but don’t try to interact with it yet.
