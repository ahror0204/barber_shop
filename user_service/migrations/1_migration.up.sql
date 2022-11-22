
CREATE TABLE "users"(
    "id" UUID PRIMARY KEY NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "user_name" VARCHAR(50) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL UNIQUE,
    "gender"  VARCHAR(10) CHECK ("gender" IN('male', 'female')) NOT NULL,
    "image_url" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE "salon"(
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
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
    "salon_id" UUID NOT NULL REFERENCES salon(id),
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "image_url" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE "personal_work_images"(
    "id" UUID PRIMARY KEY NOT NULL,
    "staff_id" UUID NOT NULL REFERENCES staff(id),
    "image_url" TEXT NOT NULL,
    "sequence_number" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);