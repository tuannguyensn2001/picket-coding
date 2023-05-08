compiler=$1
file=$2
output=$3

$compiler $file > $output

START=$(date +%s.%4N)

$compiler $file > $output

END=$(date +%s.%4N)

runtime=$(echo "$END - $START" | bc)

echo $runtime