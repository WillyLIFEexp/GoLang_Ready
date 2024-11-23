# GoLang lesson

## Description
This is the note for the GoLang class and some of my own notes

## Getting Started

### Dependencies

* Docker

### Installing

* Clone the Repo
* Build the container using the following command
    ```
    docker build -t go-container .
    ```
* Access the Go Environment
    ```
    go version
    # go version go1.23.3 linux/amd64
    ```
* We can reenter the container by the following command. 
    ```
    docker start -ai go_lesson
    # or 
    docker start my-go-container
    docker exec -it my-go-container bash
    ```

## Acknowledgments

Inspiration, code snippets, etc.
* [GoLang Class on Udemy](https://www.udemy.com/course/learn-how-to-code/)