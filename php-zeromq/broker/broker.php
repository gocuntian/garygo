<?php
//Simple request-reply broker

$context = new ZMQContext();

$frontend = new ZMQSocket($context,ZMQ::SOCKET_ROUTER);
$backend = new ZMQSocket($context,ZMQ::SOCKET_DEALER);

$frontend->bind("tcp://*:5559");
$backend->bind("tcp://*:5560");

// Initialize poll set
$poll = new ZMQPoll();

$poll->add($frontend,ZMQ::POLL_IN);
$poll->add($backend,ZMQ::POLL_IN);
$readable = $writeable = array();

//Switch message between sockets
while(true){
 $events = $poll->poll($readable,$writeable);
 
  foreach($readable as $socket){
      if($socket === $frontend){
         while(true){
           $message = $socket->recv();
           $more = $socket->getSockOpt(ZMQ::SOCKOPT_RCVMORE);
           $backend->send($message,$more ? ZMQ::MODE_SNDMORE : null);
           if(!$more){
              break;
            }
         }
      }else if($socket === $backend){
         $message = $socket->recv();
         $more = $socket->getSockOpt(ZMQ::SOCKOPT_RCVMORE);
         $frontend->send($message,$more ? ZMQ::MODE_SNDMORE : null);
         if(!$more){
           break;
         }
      } 
  }

}
