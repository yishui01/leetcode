# Read from the file file.txt and print its transposed content to stdout.

numc=$(($(head -n 1 "file.txt" | grep -o ' ' | wc -l)+1))
for ((i=1; i<="$numc"; i++)); do
cut -d' ' -f"$i" "file.txt" | paste -s -d ' '
done


# awk '{
#     for (i = 1; i <= NF; ++i) {
#         if (NR == 1) s[i] = $i;
#         else s[i] = s[i] " " $i;
#     }
# } END {
#     for (i = 1; s[i] != ""; ++i) {
#         print s[i];
#     }
# }' file.txt
