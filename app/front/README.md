# Doc
https://templ.guide/
https://www.magicbell.com/blog/setting-up-htmx-and-templ-for-go

# API stack
The front end app will be built using htmx, templ and tailwind.

# How to deploy and build ?
### Generate templ files:
cd templ && templ generate
### Refresh one liner:
cd template && templ generate && cd .. && go build main.go && ./main

## TODO
Make a nice looking login page
Setup htmx
Setup tailwind

## DONE
Init htmx and templ