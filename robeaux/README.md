# Robeaux

Robeaux (/rō-bō/) is a universal dashboard to your robotic devices. Like a router admin
page, but for robots.

Robeaux is powered by [AngularJS](http://angularjs.org/), and provides
a front-end to the API interface offered by [Artoo](http://artoo.io),
[Cylon.js](http://cylonjs.com) and [Gobot](http://gobot.io). It looks like this:

![Robeaux Interface](http://i.imgur.com/Uxb21j0.png)

## How It Works

Thanks to using a structured interface based on the RESTful and websockets APIs
of Artoo, Cylon.js and Gobot, Robeaux can query or set values on one or more
robots, and all of the devices connected to each robot. As long as they are
connected to a compatible API server.

### Robots

You can see each robot that your are connected to, and then drill in to view or
edit it.

![Robots View](http://i.imgur.com/Uxb21j0.png)

### Connections

How are you connected? Serial port? WiFi? View the details for each of your
robots.

![Connections View](http://i.imgur.com/Kof6g50.png)

### Devices

Each robot has one or more devices connected. You can view the device status and
send commands all via the web interface. You can even stream the real-time
websockets data for your buttons, switches, LEDs, sensors, and more.

![Devices View](http://i.imgur.com/AMJZCe9.png)

## Themes

You can change the theme! Pick from one of the presets, or write your own. It'll
be persisted in `localStorage` and ready to go for the next time you want to use
Robeaux.

We have three default themes:

**Default Theme**:

![Default Theme](http://i.imgur.com/AMJZCe9.png)

**Dark Theme**:

![Dark Theme](http://i.imgur.com/SNSLROf.png)

**Flat Theme**:

![Flat Theme](http://i.imgur.com/AVzMd7j.png)

You can also use the theme editor to write your own custom themes:

![Custom Theme Editor](http://i.imgur.com/pxxmhfD.png)

## Pushing new versions

Before you push a new release, make sure to create a tag:

    git tag -m "[VERSION]" [VERSION]

And push it to GitHub

    git push --tags

### RubyGems:

    gem build robeaux.gemspec
    gem push robeaux-[VERSION].gem

### NPM

    npm publish ./

## LICENSE

Copyright (c) 2014 The Hybrid Group

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
