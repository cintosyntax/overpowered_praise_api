# Overpowered Praise API

[![Build Status](https://travis-ci.org/cintosyntax/overpowered_praise_api.svg)](https://travis-ci.org/cintosyntax/overpowered_praise_api.svg)

A ridiculous API that serves insane compliments. This utilizes the Internet Chuck Norris Database (http://www.icndb.com/) API as a foundation. This API proxies the call and replaces the name of 'Chuck Norris' to the input provided by the user.

# Results

I have deployed this on Heroku using Docker here(https://overpoweredpraiseapi.herokuapp.com). It just responds back with a JSON that accomodates the format expected by slack.

I define a few slash commands on slack and then presto!
![alt text](https://raw.githubusercontent.com/cintosyntax/overpowered_praise_api/travis-test/slack_example.png)


## Wish list
- Build a seperate Go package that handles API calls to ICNDB.
- Build a slack friendly endpoint mode.
