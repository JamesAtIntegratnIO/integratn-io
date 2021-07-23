---
draft: true
layout: post
title: My Hugo+Netlify setup
date: 2021-07-23T22:17:46.983Z
categories:
  - tutorial
tags:
  - hugo
  - netlify
  - howTo
nocomment: false
---


So here is how I got my blog all set up. It was a lot of back and forth on how I set this up. So this post is for those that read it. But also for me. It was a lot of work and I want to capture what I did so I can reproduce it later.

## Setting up Hugo

First thing I did was install hugo. You can find instructions for it [here](https://gohugo.io/getting-started/installing/). After I had it installed I used `hugo new site ${siteName}` to create my site.

## Version Control is Key

After that I wanted to have git setup so that I had solid version control incase I broke anything. `git init` solved that really quickly. Then I created my repo on [Github](https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-on-github/creating-a-new-repository). After you do that it will give you instructions on how to create your first commit and change your origin for the remote and push your code. 

## Gotta Make it Pretty

Next up is adding the theme. You can find a lot of them on the [Hugo site](https://themes.gohugo.io/). Use `git submodule add ${url to theme repo} themes/{theme-name}` to put it in place.

## Deploy It

Now I have everything installed. Lets see it run. To run it locally to make sure everything works you can use `hugo serve`. But if you really want to get going then you will need to deploy it somewhere online. I put mine on netlify as I knew the name already and the process was simple enough. 

Create your account on [netlify](https://app.netlify.com/signup). After your account is created you can create a new site with this magic button ![new site button](/images/newsitebutton.png)
