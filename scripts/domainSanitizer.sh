#!/bin/bash
function methodOne() {
while IFS= read -r line; do
    keep=$(echo $line | sed 's|.*://||' | sed 's|/.*$||'); echo ${keep%.} |tr '[:upper:]' '[:lower:]'| rev | cut -d '.' -f 1,2 | rev
done < domains.txt > tmp.txt
cat tmp.txt |sort|uniq
}

function methodTwo() {
while IFS= read -r line; do
    echo ${line%.} |awk '{gsub(/.*:\/\//,""); gsub(/\/.*$/,""); print tolower($NF)}' | awk -F'.' '{print $(NF-1)"."$NF}'
done < domains.txt > tmp.txt
cat tmp.txt |sort|uniq
}

methodOne
echo "----------------------------------------"
methodTwo

rm tmp.txt
