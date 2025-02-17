// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateBlockBruteForce(ctx context.Context, arg CreateBlockBruteForceParams) (BlockBruteForce, error)
	CreateBudgetPlan(ctx context.Context, arg CreateBudgetPlanParams) (BudgetPlan, error)
	CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error)
	CreateGoal(ctx context.Context, arg CreateGoalParams) (Goal, error)
	CreateLKPlanWiseSession(ctx context.Context, arg CreateLKPlanWiseSessionParams) (LKPlanWiseSession, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (TransactionHistory, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	DeleteBlockBruteForce(ctx context.Context, id uuid.UUID) (BlockBruteForce, error)
	DeleteBudgetPlan(ctx context.Context, id uuid.UUID) error
	DeleteExpense(ctx context.Context, id uuid.UUID) error
	DeleteGoal(ctx context.Context, id uuid.UUID) error
	DeleteLKPlanWiseSession(ctx context.Context, id uuid.UUID) (LKPlanWiseSession, error)
	DeleteRole(ctx context.Context, id uuid.UUID) error
	DeleteTransaction(ctx context.Context, id uuid.UUID) error
	GetAccountByEmail(ctx context.Context, email string) (Account, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
	GetAllAccounts(ctx context.Context) ([]Account, error)
	GetAllBudgetPlans(ctx context.Context) ([]BudgetPlan, error)
	GetAllExpenses(ctx context.Context) ([]Expense, error)
	GetAllGoals(ctx context.Context) ([]Goal, error)
	GetAllRoles(ctx context.Context) ([]Role, error)
	GetAllTransactions(ctx context.Context) ([]TransactionHistory, error)
	GetBlockBruteForceByEmail(ctx context.Context, email string) (BlockBruteForce, error)
	GetBlockBruteForceById(ctx context.Context, id uuid.UUID) (BlockBruteForce, error)
	GetBudgetPlanById(ctx context.Context, id uuid.UUID) (BudgetPlan, error)
	GetExpenseById(ctx context.Context, id uuid.UUID) (Expense, error)
	GetGoalById(ctx context.Context, id uuid.UUID) (Goal, error)
	GetLKPlanWiseSessionById(ctx context.Context, id uuid.UUID) (LKPlanWiseSession, error)
	GetLKPlanWiseSessionForAuth(ctx context.Context, arg GetLKPlanWiseSessionForAuthParams) (LKPlanWiseSession, error)
	GetRoleById(ctx context.Context, id uuid.UUID) (Role, error)
	GetTransactionById(ctx context.Context, id uuid.UUID) (TransactionHistory, error)
	PagedAccounts(ctx context.Context, arg PagedAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateBlockBruteForce(ctx context.Context, arg UpdateBlockBruteForceParams) (BlockBruteForce, error)
	UpdateBudgetPlan(ctx context.Context, arg UpdateBudgetPlanParams) (BudgetPlan, error)
	UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error)
	UpdateGoal(ctx context.Context, arg UpdateGoalParams) (Goal, error)
	UpdateLKPlanWiseSession(ctx context.Context, arg UpdateLKPlanWiseSessionParams) (LKPlanWiseSession, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (TransactionHistory, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
