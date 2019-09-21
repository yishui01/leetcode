# Read from the file file.txt and output all valid phone numbers to stdout.

cat file.txt |  grep -E '^(\([0-9]{3}\) [0-9]{3}-[0-9]{4}|([0-9]{3}-){2}[0-9]{4})$'


#awk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' file.txt

