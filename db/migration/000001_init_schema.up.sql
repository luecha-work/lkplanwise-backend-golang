CREATE TABLE "Accounts" (
  "Id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "FirstName" varchar(100),
  "LastName" varchar(100),
  "Email" varchar(255),
  "PasswordHash" text,
  "DateOfBirth" date,
  "RoleId" uuid NOT NULL,
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "Roles" (
  "Id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "RoleCode" varchar(50) UNIQUE NOT NULL,
  "RoleName" varchar(50),
  "CreatedAt" timestamptz,
  "UpdatedAt" timestamptz,
  "CreatedBy" varchar(100),
  "UpdatedBy" varchar(100)
);

CREATE TABLE "Expense" (
  "ExpenseId" uuid UNIQUE PRIMARY KEY NOT NULL,
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
  "GoalId" uuid UNIQUE PRIMARY KEY NOT NULL,
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
  "TransactionId" uuid UNIQUE PRIMARY KEY NOT NULL,
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
  "BudgetId" uuid UNIQUE PRIMARY KEY NOT NULL,
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

ALTER TABLE "Accounts" ADD FOREIGN KEY ("RoleId") REFERENCES "Roles" ("Id");

ALTER TABLE "Expense" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");

ALTER TABLE "Goal" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");

ALTER TABLE "TransactionHistory" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");

ALTER TABLE "BudgetPlan" ADD FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id");
