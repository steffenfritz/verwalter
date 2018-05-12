# verwalter
A tool to manage the boring IT sec book-keeping tasks.

# Installation

## Arch Linux

yaourt -S verwalter

## FreeBSD

pkg install verwalter

## Source

    go get github.com/steffenfritz/verwalter


After installation create the database with `verwalter -init`

Database is located in the user's directory in ~/.verwalter.

# Usage

1. Start the server with `verwalter`
2. Point your browser to 127.0.0.1:8666

If you don't like the poor web gui, use a sqlite tool to access and manipulate the db \\\_(^_^)\_/

# How and when to commit 
There are probably process holes in this tool for your needs and some bugs. Clone verwalter, hack it. If you think your changes might be useful for others, create a pull request with a short, descriptive commit message!