<?php

$context = new ZMQContext();

$responder = new ZMQSocket($context,ZMQ::SOCKET_REP);
$responder->connect("tcp://localhost:5560");

while(true){
    $string = $responder->recv();
    printf ("Received request: [%s]%s", $string, PHP_EOL);
    // Do some 'work'
    sleep(1);
    //  Send reply back to client
    $responder->send("World");
}
