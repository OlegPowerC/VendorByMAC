oui2.txt it is a custom oui database wtich has format: (You can add new data):
"XXXXXX":"Vendor name" where XXXXXX first 6 digit (in hex) of MAC address
example:
"0019B9":"Dell Inc."

oui.txt it is standard OUI database, you can download it from https://standards-oui.ieee.org/oui/oui.txt
example:
00-22-72   (hex)		American Micro-Fuel Device Corp.
002272     (base 16)		American Micro-Fuel Device Corp.
				2181 Buchanan Loop
				Ferndale  WA  98248
				US

oui36.txt it is standard OUI database with ranges, you can download it from https://standards-oui.ieee.org/oui36/oui36.txt
example:
70-B3-D5   (hex)		2M Technology
719000-719FFF     (base 16)		2M Technology
				802 Greenview Drive
				Grand Prairie  TX  75050
				US

Function TestedStruct.Init() accept 3 parameters, first - custom DB filename, second - stndard DB filename
and last - oui36 DB filename.
All filename len must be minimum 5.

For example:
TestedStruct.Init("oui2.txt", "oui.txt","oui38.txt")
You can skip parce any file, provide empty string instead filename - example:
TestedStruct.Init("", "oui.txt","oui28.txt")

I found unknown MAC - 00a2.3ca0.2107
I try to check it by: https://www.wireshark.org/tools/oui-lookup.html and got: (no matches)

