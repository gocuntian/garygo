<?php

$context = new ZMQContext();

$receiver = new ZMQSocket($context,ZMQ::SOCKET_PULL);
$receiver->bind("tcp://*:5558");

$string = $receiver->recv();

$tstart = microtime(true);

$total_msec = 0;
for($task_nbr = 0;$task_nbr < 100;$task_nbr++){
   $string = $receiver->recv();
   if($task_nbr%10 == 0){
     echo ":";
   }else{
     echo ".";
   }
}

$tend = microtime(true);

$total_msec = ($tend - $tstart) * 1000;
echo PHP_EOL;
printf ("Total elapsed time: %d msec", $total_msec);
echo PHP_EOL;
