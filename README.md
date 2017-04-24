# vbmatch

Find matching thread titles on VBulletin boards (and generate quicklinks to thread pages)

## Dependencies

```
go get -u
```

## Example

```
# matches.lst
Blog
Version
```

```bash
Usage of ./vbmatch:
  -base-url string
    	Base URL
  -debug
    	Enable debug output.
  -forum-url string
    	(Sub-)Forum URL
  -match-list string
    	Match list (default "matches.lst")


./vbmatch -forum-url "http://forum.vbulletin-germany.com/forumdisplay.php/112-vBulletin-Blog-Fragen-und-Probleme"

http://forum.vbulletin-germany.com/showthread.php/54680-gel%C3%B6schte-Blogeintr%C3%A4ge-bleiben-in-der-Sidebar-sichtbar?s=02df905fa8d054a9b51ce6637243d648&page=1000
...
```

