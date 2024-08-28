CREATE TABLE "notes" (
    "id" BIGSERIAL PRIMARY KEY,
    "text" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP
);