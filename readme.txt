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

iab.txt it is standard OUI database with ranges like oui36 (but can not be use for EUI-64), you can download it from https://standards-oui.ieee.org/iab/iab.txt

Function TestedStruct.Init() accept 4 parameters, first - custom DB filename, stndard DB filename, oui36 DB filename, IAB DB filename
All filename len must be minimum 5.

For example:
TestedStruct.Init("oui2.txt", "oui.txt","oui36.txt","iab.txt")
You can skip parce any file, provide empty string instead filename - example:
TestedStruct.Init("", "oui.txt","oui36.txt","iab.txt")

I found unknown MAC - 00a2.3ca0.2107
I try to check it by: https://www.wireshark.org/tools/oui-lookup.html and got: (no matches)

Few known MAC addresses for example:
58-03-FB-4B-FA-07
0004.a300.e404
8C:8D:28:46:49:C9

Added JS code generator.
It is used for make FireFox extension mactoven: https://addons.mozilla.org/ru/firefox/addon/mactoven/


