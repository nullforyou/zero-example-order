/*Table structure for table `greet_address` */

DROP TABLE IF EXISTS `greet_address`;

CREATE TABLE `greet_address` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `member_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
    `contact_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系人姓名',
    `contact_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系人手机号',
    `province_code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区编码 省',
    `city_code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区编码 市',
    `county_code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区编码 县',
    `province_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 省',
    `city_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 市',
    `county_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 县',
    `detailed_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '详细地址',
    `first_boot` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认地址1：是；0：否；',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户地址';

/*Table structure for table `greet_category` */

DROP TABLE IF EXISTS `greet_category`;

CREATE TABLE `greet_category` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `category_id` int NOT NULL DEFAULT '0' COMMENT '第三方Id',
    `category_pid` int NOT NULL DEFAULT '0' COMMENT '第三方父级Id',
    `category_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
    `category_picture` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类图片',
    `diff_hash` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '值hash',
    `is_recommend` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否推荐1:已推荐;0:未推荐;',
    `category_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '分类状态1：使用；0：禁用；',
    `operation_user` int NOT NULL DEFAULT '0' COMMENT '操作人id',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `goods_sort` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`),
    KEY `idx_category_pid` (`category_pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品分类';

/*Table structure for table `greet_goods` */

DROP TABLE IF EXISTS `greet_goods`;

CREATE TABLE `greet_goods` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `category_id` int NOT NULL DEFAULT '0' COMMENT '商品分类id',
    `goods_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品名称',
    `goods_picture` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品图片',
    `goods_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
    `goods_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '商品状态1:上架;0:下架;',
    `goods_is_recommend` tinyint(1) NOT NULL DEFAULT '0' COMMENT '商品是否推荐1:已推荐;0:未推荐;',
    `goods_sort` smallint NOT NULL DEFAULT '1' COMMENT '商品排序',
    `goods_sales_volume` int NOT NULL DEFAULT '0' COMMENT '商品销量',
    `actual_sales_volume` int NOT NULL DEFAULT '0' COMMENT '商品实际销量',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品';

/*Table structure for table `greet_order` */

DROP TABLE IF EXISTS `greet_order`;

CREATE TABLE `greet_order` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_serial_number` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '销巴订单号',
    `member_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
    `client` tinyint NOT NULL DEFAULT '0' COMMENT '下单源 10：APP安卓；11：APPIOS; 12：WAP；13：小程序；20：PC',
    `order_status` smallint NOT NULL DEFAULT '1' COMMENT '订单状态 -30:支付异常;-10:取消订单;10:新订单待支付;20:已支付;40:已完成待结算;50:已结算;',
    `order_status_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单状态名称(不需要手动赋值)',
    `order_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单总金额（生成订单时的总价）',
    `technical_services_fee` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '技术服务费,order_amount字段会加上技术服务费',
    `order_express_fee` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '运费',
    `goods_num` smallint NOT NULL DEFAULT '0' COMMENT '商品总数量(不需要手动赋值)',
    `order_profit` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单利润',
    `member_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '会员昵称',
    `member_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '会员手机号',
    `payment_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '应该支付总金额（生成订单时免去一切优惠券、积分抵扣、微币抵扣后需要支付的价格）',
    `payment_sn` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付订单号（调用支付服务创建支付时返回）',
    `payment_type` tinyint NOT NULL DEFAULT '0' COMMENT '支付方式： 1：银联支付 2：支付宝支付 3：微信支付 4.个人余额 5.小巴余额',
    `payment_time` datetime DEFAULT NULL COMMENT '支付时间',
    `payment_status` smallint NOT NULL DEFAULT '0' COMMENT '订单支付状态0:未支付;1:已支付等待支付结果;2:支付成功;3:支付失败;',
    `deduct_coin` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '微币抵扣金额',
    `return_credits` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '赠送积分（订单积分）',
    `return_credits_status` tinyint NOT NULL DEFAULT '0' COMMENT '赠送积分状态：-1:赠送失败；0：未赠送；1：已赠送;2:赠送中;',
    `return_credits_failed_cause` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '积分领取失败原因',
    `is_append_price` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否补差价; 默认0=不需要; 1=需要',
    `append_price_status` tinyint NOT NULL DEFAULT '0' COMMENT '补差价状态 0:未支付;1:已支付等待支付结果;2:支付成功;3:支付失败;',
    `is_after_sales` smallint NOT NULL DEFAULT '0' COMMENT '介入退款流程 支付服务申请退款=1,支付服务退款完成=2,支付服务退款失败=3',
    `cancel_operator` tinyint DEFAULT NULL COMMENT '取消操作：1 用户取消;2 系统取消(未付款自动取消)；3：平台管理取消；4：第三方取消；',
    `cancel_time` datetime DEFAULT NULL COMMENT '取消时间，取消成功才有值(不需要手动赋值)',
    `cancel_cause` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '取消原因（如果时自动取消也需要注明为超时自动取消）',
    `is_user_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户删除(0：未删除； 1：已删除)',
    `user_delete_time` datetime DEFAULT NULL COMMENT '用户删除时间，删除成功才有值(不需要手动赋值)',
    `payment_limit_time` datetime DEFAULT NULL COMMENT '支付时限，一般下单后有15分钟支付时限 ，必须在此时间前发起支付出票，超过这个时间订单将自动取消',
    `settle_time` datetime DEFAULT NULL COMMENT '结算时间，结算成功才有值(不需要手动赋值)',
    `complete_time` datetime DEFAULT NULL COMMENT '完成时间，完成才有值(不需要手动赋值)',
    `invoice_status` tinyint NOT NULL DEFAULT '1' COMMENT '发票状态 1：未开具 2：申请中 3：已开',
    `admin_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注（后台订单备注）',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `is_deduct_coin` tinyint DEFAULT NULL COMMENT '是否使用抵扣; 默认0=不使用; 1使用红包抵扣;2余额抵扣;',
    `deduct_balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '余额抵扣金额',
    `deduct_type` tinyint NOT NULL DEFAULT '0' COMMENT '抵扣类型:0-未抵扣 1-红包抵扣 2-个人余额抵扣 3-小巴余额',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_serial_number` (`order_serial_number`),
    KEY `idx_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单';

/*Table structure for table `greet_order_detail` */

DROP TABLE IF EXISTS `greet_order_detail`;

CREATE TABLE `greet_order_detail` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id` int NOT NULL DEFAULT '0' COMMENT '订单id',
    `appointment_start_time` datetime NOT NULL COMMENT '预约开始时间',
    `appointment_end_time` datetime NOT NULL COMMENT '预约结束时间',
    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '给工厂的备注',
    `express_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '给快递备注',
    `sender_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发件人姓名',
    `sender_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发件人手机号',
    `sender_province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 省',
    `sender_city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 市',
    `sender_county` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 县',
    `sender_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '详细地址',
    `receive_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '收件人姓名',
    `receive_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '收件人手机号',
    `receive_province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 省',
    `receive_city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 市',
    `receive_county` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地区名称 县',
    `receive_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '详细地址',
    `is_packing` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否需要包装 0不需要，1需要',
    `pack_fee` decimal(5,2) NOT NULL DEFAULT '0.00' COMMENT '包装费',
    `out_total_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '第三方订单总价（同步第三方订单时返回）',
    `send_delivery_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '送程快递信息（同步第三方订单时返回）',
    `receive_delivery_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '返程快递信息（同步第三方订单时返回）',
    `refund_delivery_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '退款快递信息（同步第三方订单时返回）',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `city_name` varchar(122) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_id` (`order_id`),
    KEY `idx_sender_name` (`sender_name`),
    KEY `idx_sender_mobile` (`sender_mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='补差价订单';

/*Table structure for table `greet_order_goods` */

DROP TABLE IF EXISTS `greet_order_goods`;

CREATE TABLE `greet_order_goods` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id` int NOT NULL DEFAULT '0' COMMENT '订单id',
    `category_id` int NOT NULL DEFAULT '0' COMMENT '商品分类id',
    `category_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
    `goods_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
    `goods_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品名称',
    `goods_picture` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品图片',
    `goods_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
    `goods_num` smallint NOT NULL DEFAULT '0' COMMENT '商品数量',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_goods_id` (`goods_id`),
    KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单商品';

/*Table structure for table `greet_order_payment` */

DROP TABLE IF EXISTS `greet_order_payment`;

CREATE TABLE `greet_order_payment` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id` int NOT NULL DEFAULT '0' COMMENT '订单id|补款单id',
    `member_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
    `business_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务类型',
    `order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '销巴订单号|补款单号，未加索引，不能做条件字句',
    `payment_status` smallint NOT NULL DEFAULT '0' COMMENT '支付单状态1：已创建待支付；2：已支付；3：支付失败；4:支付异常,需要人工介入;',
    `payment_sn` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付订单号（调用支付服务创建支付时返回）',
    `payment_amount` decimal(10,2) DEFAULT NULL COMMENT '支付单支付金额',
    `payment_type` tinyint NOT NULL DEFAULT '0' COMMENT '支付方式： 1：银联支付 2：支付宝支付 3：微信支付 4.个人余额 5.小巴余额',
    `payment_params` json DEFAULT NULL COMMENT '创建支付订单时的入参',
    `payment_result` json DEFAULT NULL COMMENT '创建支付订单时的响应',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_order_id` (`order_id`),
    KEY `idx_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单支付订单信息';
