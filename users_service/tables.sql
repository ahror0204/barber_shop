CREATE TYPE gender AS('male', 'female');
CREATE TABLE "customers"(
    "id" UUID PRIMARY KEY NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "gender" gender NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "image_url" TEXT NOT NULL
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
);

CREATE TABLE "salon"(
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "rating" INTEGER,
    "address" TEXT NOT NULL,
    "latitude" VARCHAR(255) NOT NULL,
    "longitude" VARCHAR(255) NOT NULL,
    "start_time" TIMESTAMP,
    "end_time" TIMESTAMP,
    "image_url" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE "staff"(
    "id" UUID PRIMARY KEY NOT NULL,
    "salon_id" UUID NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "image_url" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE "personal_work_images"(
    "id" UUID PRIMARY KEY NOT NULL,
    "staff_id" UUID NOT NULL,
    "image_url" TEXT NOT NULL,
    "sequence_number" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "services"(
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "salon_id" UUID NOT NULL,
    "price" DECIMAL(8, 2) NOT NULL,
    "type" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "orders"(
    "id" UUID PRIMARY KEY NOT NULL,
    "customer_id" UUID NOT NULL,
    "staff_id" UUID NOT NULL,
    "salon_id" UUID NOT NULL,
    "total_amount" INTEGER NOT NULL,
    "discount" DECIMAL(8, 2),
    "order_start-time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "order_end-time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "order_day" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "order_month" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "order_items"(
    "id" UUID PRIMARY KEY NOT NULL,
    "service_id" UUID NOT NULL,
    "order_id" UUID NOT NULL,
    "service_name" INTEGER NOT NULL,
    "price" DECIMAL(8, 2) NOT NULL
);