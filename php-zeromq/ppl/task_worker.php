<?php

$context = new ZMQContext();

$receiver = new ZMQSocket($context,ZMQ::SOCKET_PULL);
$receiver->connect("tcp://localhost:5557");

$sender = new ZMQSocket($context,ZMQ::SOCKET_PUSH);
$sender->connect("tcp://localhost:5558");

while(true){
   $string = $receiver->recv();
   echo $string.PHP_EOL;
   usleep($string*1000);
   $sender->send("");
}
