<?php

$context = new ZMQContext(1);
$responder = new ZMQSocket($context,ZMQ::SOCKET_REP);
$responder->bind("tcp://*:5555");

while(true){
   $request = $responder->recv();
   printf("Received request:[%s]\n",$request);
   // Do some 'work'
   sleep(1);

   $responder->send("World");
}
