<?php

$context = new ZMQContext();

$sender = new ZMQSocket($context,ZMQ::SOCKET_PUSH);
$sender->bind("tcp://*:5557");
echo "Press Enter when the workers are ready:";
$fp = fopen('php://stdin','r');

$line = fgets($fp,512);

fclose($fp);

echo "Sendng tasks to workers...",PHP_EOL;

$sender->send(0);

$total_msec = 0;

for($task_nbr=0;$task_nbr < 100;$task_nbr++){
   $workload = mt_rand(1,100);
   $total_msec+=$workload;
   $sender->send($workload);
}

printf ("Total expected cost: %d msec\n", $total_msec);
sleep (1); // Give 0MQ time to deliver


