export INTERVAL=6

rm -f *.png
rm -f out*.gif
for VAR in {1..10}
do
  echo $VAR;./glitch -seed=$RANDOM -interval=$INTERVAL; mv blank.png $VAR.png; convert $VAR.png out-$VAR.gif
done
gifsicle --loop --colors 128 --delay 10 out*.gif > anim.gif 
