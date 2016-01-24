# Festival

A minimalistic event page generator written in Go.


## Overview
Festival is a event page generator written in Go. It generates static html page which you can then share with your invitees. It can intelligently guess the event date and title from the event description  

## Screenshots

![Festival](https://raw.githubusercontent.com/gophergala2016/Festival/master/public/scr.png)

Check `screenshots` folder for more

## Features
* Markdown support
* Auto-detect Date and Title
* Fully Static

## Dependencies

    https://github.com/HouzuoGuo/tiedot
    https://github.com/russross/blackfriday

## Hosting your own Festival clone

Clone the repository

    git clone https://github.com/gophergala2016/Festival.git

Install the dependencies

    go get github.com/HouzuoGuo/tiedot
    go get github.com/russross/blackfriday

Build

    go build

Start the process

    ./Festival

Listens on `:3001` by default

