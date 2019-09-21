# Read from the file words.txt and output the word frequency list to stdout.
#还有考shell的？
grep -oE '[a-z]+' words.txt | sort | uniq -c | sort -nr | awk '{print $2" "$1}' 


# for i in `cat words.txt`
# do
# echo $i 
# done | sort | uniq -c | sort -nr | awk '{print $2 " " $1}' 


# awk '{
#     for (i = 1; i <= NF; ++i) ++s[$i];
# } END {
#     for (i in s) print i, s[i];
# }' words.txt | sort -nr -k 2

