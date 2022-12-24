cd the-final-cl-test/mystery/

clue1=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2)

echo $clue1

cat interviews/interview-$clue1

echo $MAIN_SUSPECT