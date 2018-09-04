SET mapred.job.priority=VERY_HIGH;

SELECT
    tashuo_id as id,
    vote as vote,
    create_time as creaetTime,
    publish_time as publishTime,
    update_time as update_time,
    title as title,
    abstract as abstract,
    content as content
FROM
    wiki_tashuo_tbltashuocontent
WHERE
    dt='${hiveconf:dt}'
    and status=1
