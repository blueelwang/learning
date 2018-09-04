<?php
/**
 * @category    app
 * @package
 * @subpackage
 * @author      sunjianwei@baidu.com
 * @version     2017-11-01 22:07:42
 * @copyright   Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved.
 **/

$strHost    = 'localhost';
$strPort    = '19200';
$strIndex   = 'tashuo';

$arrQueryParams = array(
    'settings' => [
        'index' => [
            'number_of_shards' => 3,
            'number_of_replicas' => 1,
        ],
    ],
);
$arrResult = createIndex($arrQueryParams);
var_export($arrResult);
echo "\n";


function createIndex($arrQueryParams) {
    global $strHost, $strPort, $strIndex;

    $t1 = microtime(true);
    $resCurl = curl_init();
    $arrHeader = array(
        //'Authorization: Basic cm9vdDpyb290',
        'Content-Type: application/json',
        'Expect:',
    );
    $arrOptions = array(
        CURLOPT_URL                 => sprintf("http://$strHost:$strPort/{$strIndex}"),
        CURLOPT_CUSTOMREQUEST       => 'PUT',
        CURLOPT_RETURNTRANSFER      => true,
        CURLOPT_FOLLOWLOCATION      => true,
        CURLOPT_BINARYTRANSFER      => true,
        CURLOPT_TIMEOUT             => 10,
        CURLOPT_CONNECTTIMEOUT      => 10,
        CURLOPT_USERAGENT           => 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1)',
        CURLOPT_HTTPHEADER          => $arrHeader,
        CURLOPT_POSTFIELDS          => json_encode($arrQueryParams),
    );
    curl_setopt_array($resCurl, $arrOptions);
    $strResult = curl_exec($resCurl);
    $intCode = curl_getinfo($resCurl, CURLINFO_HTTP_CODE);
    curl_close($resCurl);

    $t2 = microtime(true);
    $dt = intval(($t2 - $t1 ) * 1000);
    return array(
        'code'  => $intCode,
        'res'   => $strResult,
        'cost'  => $dt,
    );
}
