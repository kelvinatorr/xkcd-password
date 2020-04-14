# What

This is a webapp that generates XKCD style passwords inspired by [XKCD comic 936](https://xkcd.com/936/).
These passwords are supposedly easy to remember for humans but hard to guess for computers. It randomly selects from a list of the most common 
American English words. I created the word list by combining lists from [https://www.wordfrequency.info/](https://www.wordfrequency.info/) and [http://jbauman.com/aboutgsl.html](http://jbauman.com/aboutgsl.html).

# Why

I created this to learn Go. I wanted to learn how to handle HTTP routes, read and write files, use HTML templates, manipulate Go variables, and write unit tests.

# Running with Docker

```
docker build -t xkcd-password .
docker run --restart always --name xkcd-password-0 -p 3002:8080 -d xkcd-password
```

# Running for development

```
go run .
```


# Running the tests

```
go test
```