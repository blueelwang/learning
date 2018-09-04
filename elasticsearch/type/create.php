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
$strType    = 'tashuocontent';

$arrQueryParams = [
    'properties' => [
        'title' => [
            'type'      => 'string',
            'index'     => 'analyzed',
            "analyzer"  => "ik_max_word",
        ],
        'abstract' => [
            'type'      => 'string',
            'index'     => 'analyzed',
            "analyzer"  => "ik_max_word",
        ],
        'content' => [
            'type'      => 'string',
            'index'     => 'analyzed',
            "analyzer"  => "ik_max_word",
        ],
        'vote' => [
            'type'      => 'long',
            'index'     => 'not_analyzed',
        ],
        'createTime' => [
            'type'      => 'long',
            'index'     => 'not_analyzed',
        ],
        'publishTime' => [
            'type'      => 'long',
            'index'     => 'not_analyzed',
        ],
        'updateTime' => [
            'type'      => 'long',
            'index'     => 'not_analyzed',
        ],
    ],
];

$arrResult = createType($arrQueryParams);
var_export($arrResult);
echo "\n";


function createType($arrQueryParams) {
    global $strHost, $strPort, $strIndex, $strType;

    $t1 = microtime(true);
    $resCurl = curl_init();
    $arrHeader = array(
        //'Authorization: Basic cm9vdDpyb290',
        'Content-Type: application/json',
        'Expect:',
    );
    $arrOptions = array(
        CURLOPT_URL                 => sprintf("http://$strHost:$strPort/{$strIndex}/_mapping/{$strType}"),
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
