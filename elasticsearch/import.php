<?php
/**
 * @category    app
 * @package
 * @subpackage
 * @author      sunjianwei@baidu.com
 * @version     2017-10-30 14:47:29
 * @copyright   Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved.
 **/

$strHost    = 'localhost';
$strPort    = '19200';
$strIndex   = 'tashuo';
$strType    = 'tashuocontent';

$fp = fopen($argv[1], 'rb') or die('open file failed');
while($strLine = fgets($fp)) {
    $strLine = trim($strLine);
    if (empty($strLine)) {
        continue;
    }
    processLine($strLine);
}
fclose($fp);

function processLine($strLine) {
    global $strHost, $strPort, $strIndex, $strType;

    list($id, $vote, $createTime, $publishTime, $updateTime, $title, $abstract, $content) = explode("\t", $strLine);

    $strUrl = "http://$strHost:$strPort/{$strIndex}/{$strType}/" . $id;
    $arrHeader = array(
        //'Authorization: Basic cm9vdDpyb290',
        'Content-Type: application/json',
        'Expect:',
    );
    $arrData = array(
        'vote'          => intval($vote),
        'createTime'    => intval($createTime),
        'publishTime'   => intval($publishTime),
        'updateTime'    => intval($updateTime),
        'title'         => strip_tags(strval($title)),
        'abstract'      => strip_tags(strval($abstract)),
        'content'       => strip_tags(strval($content)),
    );

    file_put_contents('log/log', "$id\t", FILE_APPEND);
    post($strUrl, $arrHeader, $arrData);
}

function post($strUrl, $arrHeader, $arrPostData) {
    $t1 = microtime(true);

    $resCurl = curl_init();
    $arrOptions = array(
        CURLOPT_URL                 => $strUrl,
        CURLOPT_CUSTOMREQUEST       => 'PUT',
        CURLOPT_RETURNTRANSFER      => true,
        CURLOPT_FOLLOWLOCATION      => true,
        CURLOPT_BINARYTRANSFER      => true,
        //CURLOPT_TIMEOUT             => 60,
        CURLOPT_CONNECTTIMEOUT      => 60,
        CURLOPT_USERAGENT           => 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1)',
        CURLOPT_HTTPHEADER          => $arrHeader,
        CURLOPT_POSTFIELDS          => json_encode($arrPostData),
    );
    curl_setopt_array($resCurl, $arrOptions);
    $strResult = curl_exec($resCurl);
    $arrInfo = curl_getinfo($resCurl);
    $intCode = $arrInfo['http_code'];
    curl_close($resCurl);

    $t2 = microtime(true);
    $dt = intval(($t2 - $t1 ) * 1000);
    if (200 == $intCode || 201 == $intCode) {
        file_put_contents('log/log', "succ\t$dt\t$intCode\n", FILE_APPEND);
        return true;
    }
    file_put_contents('log/log', "fail\t$dt\t$intCode\t$strResult\t" . json_encode($arrInfo) . "\n", FILE_APPEND);
    return false;
}
