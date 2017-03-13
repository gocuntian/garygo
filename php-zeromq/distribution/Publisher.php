<?php
$CONFIG["TAOKE_BTS"]["ENABLE"] = true;
 
$CONFIG["QP_BTS"]["ENABLE"] = true;
 
$CONFIG["QP_BTS"]["TK_TEST"] = 13;
 
$string = json_encode ($CONFIG);

$clients = array("s2","s1","s3");

$context = new ZMQContext (10);
 
//Socket talk to clients
 
$publisher = new ZMQSocket ($context,ZMQ::SOCKET_PUB);
 
$publisher->bind ("tcp://*:5561");

$server = new ZMQSocket ($context,ZMQ::SOCKET_REP);
 
$server->bind ("tcp://*:5562");
 
while(count($clients)!=0) {
 
        $client_name = $server->recv ();
 
        echo "{$client_name} is connect!\r\n";
 
if (in_array($client_name, $clients)) { //coming one client
 
        $key = array_search($client_name, $clients);
 
        unset($clients[$key]);
 
        echo "$client_name has come in!\r\n";
        
        $server->send ("Version is 2.0");
 
} else {
 
        $server->send ("You are a stranger!");
 
}

}
 $publisher->send ($string);
