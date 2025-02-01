ALTER TABLE "BudgetPlan" DROP CONSTRAINT "BudgetPlan_AccountId_fkey";
ALTER TABLE "TransactionHistory" DROP CONSTRAINT "TransactionHistory_AccountId_fkey";
ALTER TABLE "Goal" DROP CONSTRAINT "Goal_AccountId_fkey";
ALTER TABLE "Expense" DROP CONSTRAINT "Expense_AccountId_fkey";
ALTER TABLE "Accounts" DROP CONSTRAINT "Accounts_RoleId_fkey";

DROP TABLE IF EXISTS "BlockBruteForce";
DROP TABLE IF EXISTS "LKPlanWiseSession";
DROP TABLE IF EXISTS "BudgetPlan";
DROP TABLE IF EXISTS "TransactionHistory";
DROP TABLE IF EXISTS "Goal";
DROP TABLE IF EXISTS "Expense";
DROP TABLE IF EXISTS "Accounts";
DROP TABLE IF EXISTS "Roles";