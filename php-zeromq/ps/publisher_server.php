<?php

$context = new ZMQContext();

$publisher = $context->getSocket(ZMQ::SOCKET_PUB);
$publisher->bind("tcp://*:5556");

while(true){
   $zipcode = mt_rand(100000,999999);
   $temperature = mt_rand(-80,136);
   $relhumidity = mt_rand(10,60);
   $update = sprintf("%05d %d %d",$zipcode,$temperature,$relhumidity);
   $publisher->send($update);
}
