# sleeptimer

sleeptimer is a Windows command line version of a program I wrote in college to make my computer shut down, restart or quit certain programs after a set amount of time.

# usage

```bat
> sleeptimer [-action (shutdown|restart)] [-mute] [-blank] [-ttl duration]
```

* `ttl` is required and a duration > 0 is required. Use a string parseable by [`time.ParseDuration`](https://golang.org/pkg/time/#ParseDuration), i.e. `10s` or `1h30m`.
* If `action` is set to `shutdown` or `restart` the computer will perform that action when the ttl expires. At most, one of these actions can be performed.
* If `mute` or `blank` is set, the sound will be muted or the screen will go blank (respectively). You can use both of these flags at the same time.
