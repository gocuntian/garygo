<?php
$conn = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
if(!$conn){
    return false;
}
socket_set_option($conn,SOL_SOCKET, SO_SNDTIMEO,array(
    "sec" => 0,
    "usec" => 50000
));

$result = socket_connect($conn,'127.0.0.1',9002);
$s = time();

$json = array('id'=>time(), 'params'=>[1],'method'=>"Serve.Echo");

socket_write($conn,json_encode($json));
$r = socket_read($conn,1024);
socket_close($conn);
echo $r;

