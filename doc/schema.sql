CREATE TABLE "Accounts" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "FirstName" varchar(100),
  "LastName" varchar(100),
  "UserName" varchar(100) NOT NULL UNIQUE,
  "Email" varchar(100),
  "PasswordHash" text,
  "DateOfBirth" varchar,
  "RoleId" uuid NOT NULL,
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "Roles" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "RoleCode" varchar(50) UNIQUE NOT NULL,
  "RoleName" varchar(50),
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "Expense" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "AccountId" uuid NOT NULL,
  "Category" varchar(100),
  "Amount" decimal(10,2),
  "Date" timestamptz,
  "Description" text,
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "Goal" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "AccountId" uuid NOT NULL,
  "GoalType" varchar(100),
  "TargetAmount" decimal(10,2),
  "CurrentAmount" decimal(10,2),
  "Deadline" timestamptz,
  "Progress" decimal(5,2),
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "TransactionHistory" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "AccountId" uuid NOT NULL,
  "TransactionType" varchar(50),
  "Amount" decimal(10,2),
  "Description" text,
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "BudgetPlan" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "AccountId" uuid NOT NULL,
  "Month" varchar(20),
  "TotalIncome" decimal(10,2),
  "TotalExpenses" decimal(10,2),
  "SavingsGoal" decimal(10,2),
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "LKPlanWiseSession" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "AccountId" uuid,
  "LoginAt" timestamptz,
  "Platform" varchar(100),
  "Os" varchar(100),
  "Browser" varchar(100),
  "LoginIp" varchar(100) NOT NULL,
  "IssuedTime" timestamptz,
  "ExpirationTime" timestamptz,
  "SessionStatus" varchar(1) NOT NULL,
  "Token" text,
  "RefreshTokenAt" timestamptz,
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "BlockBruteForce" (
  "Id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "UserName" varchar(100) NOT NULL,
  "Count" int,
  "Status" varchar(1) NOT NULL,
  "LockedTime" timestamptz,
  "UnLockTime" timestamptz,
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

ALTER TABLE "Accounts" ADD FOREIGN KEY ("RoleId") REFERENCES "Roles" ("Id");

ALTER TABLE "Expense" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");

ALTER TABLE "Goal" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");

ALTER TABLE "TransactionHistory" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");

ALTER TABLE "BudgetPlan" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");