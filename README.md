# redis

## Objectives:

I will be attempting to make a redis clone from scratch using Go programming language.
Before we make our own redis i want use to get familiar with how redis works.

## What is redis?:

Think of redis like a super-fast, organized storage box that lives in your computer memory. It allows you to quickly store and retrieve data. 
Imagine a bookshelf where you can instantly grab or place a book because you know exactly where it goes and what it's called.

## Installing redis(original):
I will be using Ubuntu for this project.

1. update your package list:
```
	$ sudo apt-get update
```

2. install redis after updating
```
	$ sudo apt-get install redis-server
```

3. start the redis server
```
	$ sudo systemctl start redis-server
```

4. allow redis to start automatically when you boot your system
```
	$ sudo systemctl enable redis-server
```

5. if you want to confirm if redis is working
```
	$ redis-cli ping
```
"PONG" should be displayed if everything is ok.


## Getting used to redis.
Well we will have to get used to redis first in order to make our own redis.

Redis stores data in key-value format, like a dictionary or a map in some programming languages.

1. Open Redis Command Line Interface:

```
	$ redis-cli
```