<?php

$conn_args = array(
    'host' => '127.0.0.1',
    'port' => '5672',
    'login' => 'guest',
    'password' => 'guest',
    'vhost'=>'/'
);
$e_name = 'myex'; //交换机名
//$q_name = 'myq'; //无需队列名
$k_route = 'myr'; //路由key

//创建连接和channel
$conn = new AMQPConnection($conn_args);
if (!$conn->connect()) {
    die("Cannot connect to the broker!\n");
}
$channel = new AMQPChannel($conn);


//创建交换机对象
$ex = new AMQPExchange($channel);
$ex->setName($e_name);
//发送消息
//$channel->startTransaction(); //开始事务
for($i=0; $i<5; ++$i){
    sleep(1);//休眠1秒
    //消息内容
    $message = "TEST MESSAGE!".date("h:i:sa");
    echo "Send Message:".$ex->publish($message, $k_route)."\n";
}
//$channel->commitTransaction(); //提交事务

$conn->disconnect();