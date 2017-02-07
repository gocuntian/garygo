<?php

$context = new ZMQContext();

//  Socket to talk to server
$requester = new ZMQSocket($context, ZMQ::SOCKET_REQ);
$requester->connect("tcp://localhost:5555");

for ($request_nbr = 0; $request_nbr < 10; $request_nbr++) {
    $requester->send("Hello");
    $string = $requester->recv();
    printf ("Received reply %d [%s]%s", $request_nbr, $string, PHP_EOL);
}
