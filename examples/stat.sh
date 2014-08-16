#1/bin/bash

stat=""
for x in {1..4}; do
	acc=$(strings $1 | cut -f$x -d / | awk '{x+=$0}END{print x}')
	stat=$stat$acc/
done

echo $stat
