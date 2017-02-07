<?php
$context = new ZMQContext();

$requester = new ZMQSocket($context,ZMQ::SOCKET_REQ);
$requester->connect("tcp://localhost:5559");

for($request_nbr = 0; $request_nbr < 10 ;$request_nbr++){
   $requester->send("Hello");
   $string = $requester->recv();
   printf ("Received reply %d [%s]%s", $request_nbr, $string, PHP_EOL);
}
