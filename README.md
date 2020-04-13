# Polybar: Corona Stats

WIP Polybar script that will show you local statistics from the global datasets

## Example Usage

Examples below make use of Glyphicons because I have a license, feel free to change it as required.

```
[module/coronavirus-stats-confirmed]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -d=false -r=false
interval = 900
format-underline = #3333FF

[module/coronavirus-stats-dead]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -c=false -r=false
interval = 900
format-underline = #FF3333

[module/coronavirus-stats-recovered]
type = custom/script
label =  %output%
exec = /path/to/mycorona -l "Location String" -c=false -d=false
interval = 900
format-underline = #33FF33

``` 