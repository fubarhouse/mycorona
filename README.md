# Polybar: Corona Stats

WIP Polybar script that will show you local statistics from the global datasets

## Flags

```
$ mycorona -h
  -a    Show 'active' data.
  -c    Show 'confirmed' data.
  -d    Show 'dead' data.
  -g    Show global location data
  -l string
        Specify the primary location
  -o string
        Specify the secondary/alternative location
  -r    Show 'recovered' data.
```

## Polybar Configuration

Examples below make use of Glyphicons because I have a license, feel free to change it as required.

### Module configuration

```
[module/coronavirus-stats-active]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -o "Alternative Location String" -a -g
interval = 900
format-underline = #3333FF

[module/coronavirus-stats-confirmed]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -o "Alternative Location String" -c
interval = 900
format-underline = #3333FF

[module/coronavirus-stats-dead]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -d
interval = 900
format-underline = #FF3333

[module/coronavirus-stats-recovered]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -r
interval = 900
format-underline = #33FF33
```

### Coronavirus Bar

```
[bar/coronavirus]
width = 100%
height = 40
radius = 6.0
fixed-center = false
bottom = false

background = ${colors.background}
foreground = ${colors.foreground}

line-size = 5
line-color = #f00
module-margin-left = 1
module-margin-right = 1

border-size = 4
border-color = #000000

font-0 = Fira\ Code:style=Bold
font-1 = Fira\ Mono:style=Bold
font-2 = FontAwesome:style=Regular
font-3 = GLYPHICONS\ Basic\ Set:style=Regular

modules-center = coronavirus coronavirus-stats-confirmed coronavirus-stats-recovered coronavirus-stats-active coronavirus-stats-dead

cursor-click = pointer
cursor-scroll = ns-resize
```

### Extras

```
[module/coronavirus]
type = custom/text
content = 
content-padding = 1
content-foreground = #FF3333
content-underline = #FF3333
; change the URL's below to your taste
click-left = /usr/bin/firefox https://gisanddata.maps.arcgis.com/apps/opsdashboard/index.html#/bda7594740fd40299423467b48e9ecf6
click-right = /usr/bin/firefox https://www.covid19.act.gov.au
```