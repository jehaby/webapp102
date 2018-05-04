<?php

$filename = 'bikemarkt';

$cont = file_get_contents($filename);
if ($cont == null) {
    $cont = file_get_contents('https://bikemarkt.mtb-news.de/en/categories');
    file_put_contents($filename, $cont);    
}

$dom = new DOMDocument();
$dom->loadHTML($cont);

$xpath = new DOMXPath($dom);

$classname = 'categoryContainer';
$mainCats = $xpath->query("//*[contains(@class, '$classname')]//div");

$res = [];
foreach ($mainCats as $mainCat) {
    $classname = 'categoryMain';
    $catMain = formatName($xpath->query("div[1]/div/strong/a", $mainCat)->item(0)->textContent, $mainCat);
    $res[$catMain] = [];

    $secs = $xpath->query('div[position()>1]', $mainCat);

    foreach ($secs as $sec) {
        $secName = formatName($xpath->query('div/strong/a', $sec)->item(0)->textContent);
        $res[$catMain][$secName] = array();

        $leafs = $xpath->query('div/ul/li/a', $sec);
        if ($leafs) {
            foreach ($leafs as $leaf) {
                $res[$catMain][$secName][] = formatName($leaf->textContent);
            }
        }
    }
}

$id = 1;
$values = [];
foreach ($res as $cat => $subcats) {
    $curID = $id; $id += 20;

    $values[] = path($curID, "$curID", $cat);

    foreach ($subcats as $subcat => $leafs) {
        $curID2 = $id; $id += 20;
        $values[] = path($curID2, "$curID, $curID2", $subcat);

        foreach ($leafs as $leaf) {
            $curID3 = $id; $id += 20;
            $values[] = path($curID3, "$curID, $curID2, $curID3", $leaf);
        }
    }
}

function path($id, $path, $name) {
    return sprintf("(%d, '{%s}', '%s')", $id, $path, $name);
}

// var_dump($res);
// var_dump($values);

file_put_contents('res.sql', implode(",\n", $values));

printf("Written %d categories \n", count($values));


function formatName($name) {
    return str_replace(' ', '_', strtolower($name));
}

// var_dump($mainCats);
