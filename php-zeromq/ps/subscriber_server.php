<?php

$context = new ZMQContext();

echo "Connecting updates from weather server...",PHP_EOL;

$subscriber = new ZMQSocket($context,ZMQ::SOCKET_SUB);
$subscriber->connect("tcp://localhost:5556");
$filter = $_SERVER['argc'] > 1 ? $_SERVER['argv'][1] : "10001";

$subscriber->setSockOpt(ZMQ::SOCKOPT_SUBSCRIBE,$filter);

$total_temp = 0;
for($update_nbr=0;$update_nbr<100;$update_nbr++){
   $string = $subscriber->recv();
   sscanf($string,"%d %d %d",$zipcode,$temperature,$relhumidity);
   $total_temp+=$temperature;
}
printf ("Average temperature for zipcode '%s' was %dF\n",
$filter, (int) ($total_temp / $update_nbr));
