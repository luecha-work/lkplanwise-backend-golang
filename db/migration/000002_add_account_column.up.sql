ALTER TABLE "Accounts" ADD COLUMN "IsMailVerified" bool NOT NULL DEFAULT false;
ALTER TABLE "Accounts" ADD COLUMN "IsLocked" bool NOT NULL DEFAULT false;