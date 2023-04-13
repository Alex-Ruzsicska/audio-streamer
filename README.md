<h1 align="center">
<br>
Golang Audio Streamer
</h1>

<p align="center">An API that streams audio files. Build with Golang.</p>

<p align="center">
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License MIT">
  </a>
</p>

## Features
[//]: # (Add the features of your project here:)

- **[Go](https://go.dev/)** —  An open-source programming language supported by Google.
- **[Gorilla/Mux](https://pkg.go.dev/github.com/gorilla/mux)** - Implements a request router and dispatcher.
- **[go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3#section-readme)** - Provides interface to SQLite3 databases.

## Getting started

- Run "git clone git@github.com:Alex-Ruzsicska/white-border.git"
- Navigate to the App's folder
- Run "go run main.go"*

  
 \* Go has to be installed: [GO!](https://go.dev/)


## Routes
- **/folders** - Get all floders.
- **/folders/genre/{id}** - Get floders by genre.

- **/audios** - Get all audios.
- **/audios/folder/{id}** - Get audios by folder.

- **/genres** - Get all genres.

- **/authors** - Get all authors.

- **/stream/{audio_file}** - Get an audio stream.

## Database Schema
<div>
  <img src="https://i.imgur.com/vFduMdx.png" alt="demo" height="425">
</div>

<hr />







## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/licenses/MIT) page for details.
