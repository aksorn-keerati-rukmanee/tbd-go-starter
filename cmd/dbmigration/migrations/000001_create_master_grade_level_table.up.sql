CREATE OR REPLACE FUNCTION public.uuid_generate_v4()
    RETURNS uuid
    LANGUAGE c
    PARALLEL SAFE STRICT
    AS '$libdir/uuid-ossp', $function$uuid_generate_v4$function$
;

CREATE SCHEMA workbook;

CREATE TABLE workbook.master_grade_level (
    "gradeLevelId" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "nameTh" varchar(255) NULL,
    "nameEn" varchar(255) NULL,
    slug varchar(50) NULL,
    "group" varchar(50) NULL,
    "position" int4 NOT NULL DEFAULT 0,
    "gradeLevelAbbrev" varchar(50) NULL,
    "isActive" bool NOT NULL DEFAULT true,
    "createdAt" timestamptz NOT NULL DEFAULT now(),
    "createdBy" uuid NULL,
    "updatedAt" timestamptz NULL,
    "updatedBy" uuid NULL,
    "deletedAt" timestamptz NULL,
    "deletedBy" uuid NULL,
    CONSTRAINT master_grade_level_pkey PRIMARY KEY ("gradeLevelId")
);