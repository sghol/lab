<html>

<head>
    <title>First php page</title>
</head>

<body>
    <h1>Calculate two integers</h1>

    <?php
    $name = "John Doe";
    $age = 42;
    $message = "My name is " . $name . " and I am " . $age . " years old.";
    echo $message;
    // data types value could be get by var_dump
    var_dump(1);
    var_dump(3.14);
    var_dump([1, 2, 3]);
    ?>
    <p>After running some page</p>
</body>

</html>
