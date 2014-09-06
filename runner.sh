export INTERVAL=1
export COLORS=256

./glitch
for VAR in {0..9}
do
  #echo $VAR;./glitch -seed=$RANDOM -interval=$INTERVAL; mv blank.png $VAR.png; convert $VAR.png out-$VAR.gif
  convert out$VAR.png out$VAR.gif
done

gifsicle --loop --colors $COLORS --delay 10 out*.gif > anim.gif 
