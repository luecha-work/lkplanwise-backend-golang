CREATE TABLE "VerifyEmails" (
  "Id" bigserial PRIMARY KEY,
  "UserName" varchar NOT NULL,
  "Email" varchar NOT NULL,
  "SecretCode" varchar NOT NULL,
  "IsUsed" bool NOT NULL DEFAULT false,
  "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
  "ExpiredAt" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "VerifyEmails" ADD FOREIGN KEY ("UserName") REFERENCES "Accounts" ("UserName");