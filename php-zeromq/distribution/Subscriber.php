<?php
$hostname = $_SERVER['argc'] > 1 ? $_SERVER['argv'][1] : "s1";

$context = new ZMQContext(2);

$sub = new ZMQSocket($context,ZMQ::SOCKET_SUB);

$sub->connect("tcp://localhost:5561");

$sub->setSockOpt(ZMQ::SOCKOPT_SUBSCRIBE,"");

$client = $context->getSocket (ZMQ::SOCKET_REQ);

$client->connect("tcp://localhost:5562");

while(1){
    
    $client->send($hostname);
    $version = $client->recv();
    echo $version."\r\n";
    if(!empty($version)){
        $recive = $sub->recv();
        $vars = json_decode($recive);
        var_dump($vars);
    }
}