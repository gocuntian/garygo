<?php
$context = new ZMQContext();

//socket facing clients
$frontend = $context->getSocket(ZMQ::SOCKET_ROUTER);
$frontend->bind("tcp://*:5559");

//socket facing services;

$backend = $context->getSocket(ZMQ::SOCKET_DEALER);
$backend->bind("tcp://*:5560");

//start built-in device

$device = new ZMQDevice($frontend,$backend);
$device->run();
