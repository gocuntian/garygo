<?php

function step1(){
   $context = new ZMQContext();
   $sender = new ZMQSocket($context,ZMQ::SOCKET_PAIR);
   $sender->connect("ipc://step2.ipc");
   $sender->send("hello, i am step1");
}

function step2(){
   $pid = pcntl_fork();
   if($pid == 0){
      step1();
      exit();
   }
   
 $context = new ZMQContext();
 $receiver = new ZMQSocket($context,ZMQ::SOCKET_PAIR);
 $receiver->bind("ipc://step2.ipc");
 
 sleep(10);
 
 $strings = $receiver->recv();

 echo "step2 receiver is $strings".PHP_EOL;

 sleep(10);
 
 $sender = new ZMQSocket($context,ZMQ::SOCKET_PAIR);
 $sender->connect("ipc://step3.ipc");
 $sender->send($strings);  
   
}

$pid = pcntl_fork();

if($pid == 0){
   step2();
   exit();
}

$context = new ZMQContext();

$receiver = new ZMQSocket($context,ZMQ::SOCKET_PAIR);
$receiver->bind("ipc://step3.ipc");

$sr = $receiver->recv();

echo "the result is {$sr} ".PHP_EOL;
