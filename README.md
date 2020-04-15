# Polybar: Corona Stats

Polybar information script/tool that will show you local statistics from the global datasets (updated once daily at UTC 00:00)

## Installation

For convenience, you can either use `mkpkg` for Arch-based linux distributions or you can build from source. Makepkg will also compile from source, but it is an alternative go using the go toolchain.

### Method 1: `makepkg`

1. `makepkg -i`

### Method 2: from source

2. `go build -o /usr/bin/mycorona .` or `go install .`

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

Make sure the command, `` icons and hexadecimal below are changed to your liking.

### Module configuration

```
[module/coronavirus-stats-active]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -o "Alternative Location String" -a -g
interval = 3600
format-underline = #3333FF

[module/coronavirus-stats-confirmed]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -o "Alternative Location String" -c
interval = 3600
format-underline = #3333FF

[module/coronavirus-stats-dead]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -d
interval = 3600
format-underline = #FF3333

[module/coronavirus-stats-recovered]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -r
interval = 3600
format-underline = #33FF33
```

### Extras

An extra button I have which simply points to common locations where data can be found. It uses a red biohazard icon, but you can change this up to meet your needs/taste.

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

### Coronavirus Bar

Optionally, you could create a dedicated polybar like I have. In your launch script, just the following:

```
polybar -r coronavirus &
```

The bar configuration (including the extra) button described under [extras](#Extras): 

```
[bar/coronavirus]
width = 100%
height = 40
radius = 6.0
fixed-center = false
bottom = true

background = ${colors.background}
foreground = ${colors.foreground}

line-size = 5
line-color = #f00
module-margin-left = 1
module-margin-right = 1

border-size = 4
border-color = #000000

font-0 = Fira\ Code:style=Bold
font-1 = FontAwesome:style=Regular
font-2 = GLYPHICONS\ Basic\ Set:style=Regular

modules-center = coronavirus coronavirus-stats-confirmed coronavirus-stats-recovered coronavirus-stats-active coronavirus-stats-dead

cursor-click = pointer
cursor-scroll = ns-resize
```