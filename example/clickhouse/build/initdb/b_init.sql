CREATE TABLE logs
(
    id        UInt64,
    date      Date default toDate(date_time),
    date_time DateTime64(3, 'Europe/Moscow'),
    level     String,
    name      String,
    message   String,
    caller    String,
    function  String,
    stack     String,
    fields    String
) ENGINE = TinyLog;