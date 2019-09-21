# Read from the file file.txt and output the tenth line to stdout.

#head -n +10 file.txt | tail -n +10

#awk "NR==10" file.txt

sed -n 10p file.txt
