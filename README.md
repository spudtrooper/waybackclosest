# waybackclosest

Finds the closest matching URl ini archive.org.

Usage:

```
$ waybackclosest <url>
```

To install:

```
go get github.com/spudtrooper/waybackclosest
```

## Example

By default this returns the actual URL:

```
$ waybackclosest http://www.donaldjtrump.com/images/site/banner.jpg

http://web.archive.org/web/20151030190157/http://www.donaldjtrump.com/images/site/banner.jpg
```


Pass `--raw` to get the actual image path if the URL is an image:

```
$ waybackclosest --raw http://www.donaldjtrump.com/images/site/banner.jpg

http://web.archive.org/web/20151030190157if_/http://www.donaldjtrump.com/images/site/banner.jpg
```
