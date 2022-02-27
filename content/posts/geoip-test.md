---
title: "Geoip Test"
date: 2022-02-27T02:43:28Z
tags: []
categories: []
syndication: []
nocomment: false
draft: false
---
{{ $visitor := getJSON "./netlify/functions/geolocate"}}

<p>{{ visitor }}</p>