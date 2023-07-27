/*
 Navicat Premium Data Transfer

 Source Server         : 企迅云-免啦街Dev
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 10.68.74.250:5432
 Source Catalog        : mianlajiedb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 27/07/2023 09:53:54
*/


-- ----------------------------
-- Table structure for oss_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."oss_app_config";
CREATE TABLE "public"."oss_app_config" (
  "id" int8 NOT NULL,
  "app_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "available_number" int8,
  "current_limiting" int8,
  "use_number" int8,
  "start_time" timestamptz(6) NOT NULL,
  "end_time" timestamptz(6) NOT NULL,
  "status" int2 NOT NULL,
  "remark" text COLLATE "pg_catalog"."default",
  "union_main_id" int8,
  "bucket_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."oss_app_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."oss_app_config"."app_name" IS '应用名称';
COMMENT ON COLUMN "public"."oss_app_config"."available_number" IS '可用总量';
COMMENT ON COLUMN "public"."oss_app_config"."current_limiting" IS '总量';
COMMENT ON COLUMN "public"."oss_app_config"."use_number" IS '使用量';
COMMENT ON COLUMN "public"."oss_app_config"."start_time" IS '生效时间';
COMMENT ON COLUMN "public"."oss_app_config"."end_time" IS '失效时间';
COMMENT ON COLUMN "public"."oss_app_config"."status" IS '状态：0禁用 1正常';
COMMENT ON COLUMN "public"."oss_app_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."oss_app_config"."union_main_id" IS '主体id';
COMMENT ON COLUMN "public"."oss_app_config"."bucket_id" IS '存储空间id';

-- ----------------------------
-- Table structure for oss_bucket_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."oss_bucket_config";
CREATE TABLE "public"."oss_bucket_config" (
  "id" int8 NOT NULL,
  "bucket_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "endpoint" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "storage_type" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "redundant_type" varchar(64) COLLATE "pg_catalog"."default",
  "monthly_flow" int8,
  "visits_num" int4,
  "union_main_id" int8,
  "owner_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "provider_no" varchar(32) COLLATE "pg_catalog"."default",
  "state" int2
)
;
ALTER TABLE "public"."oss_bucket_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."oss_bucket_config"."bucket_name" IS '存储空间名称';
COMMENT ON COLUMN "public"."oss_bucket_config"."endpoint" IS 'bucket调用域名';
COMMENT ON COLUMN "public"."oss_bucket_config"."storage_type" IS '存储类型';
COMMENT ON COLUMN "public"."oss_bucket_config"."redundant_type" IS '冗余类型';
COMMENT ON COLUMN "public"."oss_bucket_config"."monthly_flow" IS '当月流量';
COMMENT ON COLUMN "public"."oss_bucket_config"."visits_num" IS '访问次数';
COMMENT ON COLUMN "public"."oss_bucket_config"."union_main_id" IS '主体ID';
COMMENT ON COLUMN "public"."oss_bucket_config"."owner_id" IS '拥有者ID';
COMMENT ON COLUMN "public"."oss_bucket_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."oss_bucket_config"."state" IS '状态：0禁用 1正常';

-- ----------------------------
-- Table structure for oss_service_provider_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."oss_service_provider_config";
CREATE TABLE "public"."oss_service_provider_config" (
  "id" int8 NOT NULL,
  "provider_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "provider_no" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "access_key_id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "access_key_secret" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "token" varchar(255) COLLATE "pg_catalog"."default",
  "base_path" varchar(255) COLLATE "pg_catalog"."default",
  "endpoint" varchar(255) COLLATE "pg_catalog"."default",
  "remark" text COLLATE "pg_catalog"."default",
  "status" int2,
  "ext_json" json,
  "region" varchar(32) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."oss_service_provider_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."oss_service_provider_config"."provider_name" IS '渠道商名称';
COMMENT ON COLUMN "public"."oss_service_provider_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."oss_service_provider_config"."access_key_id" IS '身份标识';
COMMENT ON COLUMN "public"."oss_service_provider_config"."access_key_secret" IS '身份认证密钥';
COMMENT ON COLUMN "public"."oss_service_provider_config"."token" IS '安全令牌';
COMMENT ON COLUMN "public"."oss_service_provider_config"."base_path" IS '域名';
COMMENT ON COLUMN "public"."oss_service_provider_config"."endpoint" IS 'bucket调用域名';
COMMENT ON COLUMN "public"."oss_service_provider_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."oss_service_provider_config"."status" IS '状态：0禁用 1启用';
COMMENT ON COLUMN "public"."oss_service_provider_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."oss_service_provider_config"."region" IS '地域';

-- ----------------------------
-- Primary Key structure for table oss_app_config
-- ----------------------------
ALTER TABLE "public"."oss_app_config" ADD CONSTRAINT "oss_app_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table oss_bucket_config
-- ----------------------------
ALTER TABLE "public"."oss_bucket_config" ADD CONSTRAINT "oss_bucket_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table oss_service_provider_config
-- ----------------------------
ALTER TABLE "public"."oss_service_provider_config" ADD CONSTRAINT "oss_service_provider_config_pkey" PRIMARY KEY ("id");
