SELECT `app_brand`.`id`, `app_brand`.`name`, `app_brand`.`icon`, `app_brand`.`app_id`, `app_brand`.`storefont_type`,
 `app_brand`.`industry`, `app_brand`.`genres`, `app_brand`.`deleted`, `app_brand`.`company_id`, `app_brand`.`modify_time`,
  `app_brand`.`create_time`,
  (
  SELECT COUNT(*) AS `count`
  FROM `app_collection` U0
  WHERE U0.`app_brand_id` = (`app_brand`.`id`)
  GROUP BY U0.`app_brand_id` ORDER BY NULL LIMIT 1
  ) AS `apps_count`,
   (
        SELECT COUNT(*) AS `count` FROM `app_brand_property` U0 WHERE
        (U0.`app_brand_id` = (`app_brand`.`id`) AND U0.`deleted` = 0 AND U0.`property_type` = 2) GROUP BY
        U0.`app_brand_id` ORDER BY NULL LIMIT 1
     ) AS `industry_list`


      FROM `app_brand` WHERE
      (
      `app_brand`.`deleted` = 0 AND

       NOT (`app_brand`.`modify_time` = '2019-01-01 00:00:00') AND
       (
           (
               SELECT COUNT(*) AS `count` FROM `app_brand_property` U0 WHERE (U0.`app_brand_id` = (`app_brand`.`id`) AND
                U0.`deleted` = 0 AND U0.`property_type` = 2) GROUP BY U0.`app_brand_id` ORDER BY NULL LIMIT 1
            ) IS NULL OR

             (
                 SELECT COUNT(*) AS `count` FROM `app_brand_property` U0 WHERE (U0.`app_brand_id` = (`app_brand`.`id`) AND
                  U0.`deleted` = 0 AND U0.`property_type` = 2) GROUP BY U0.`app_brand_id` ORDER BY NULL LIMIT 1
             ) = 0
        )
    )

    ORDER BY `app_brand`.`id` ASC;


    24761
    23942