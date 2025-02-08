package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lkplanwise-api/constant"
	db "github.com/lkplanwise-api/db/sqlc"
)

func CheckBlockedBruteForce(ctx *gin.Context, store db.Store, email string) (bool, error) {
	bruteForce, err := store.GetBlockBruteForceByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if bruteForce.Status == constant.LockedBruteForce && bruteForce.UnLockTime.Time.After(time.Now()) {
		remainingTime := time.Until(bruteForce.UnLockTime.Time)
		return true, errors.New("Account is locked, please try again in " + remainingTime.String())
	}

	return false, nil
}

func ManageBlockBruteForce(
	ctx *gin.Context,
	store db.Store,
	email string) (db.BlockBruteForce, error) {

	bruteForce, err := store.GetBlockBruteForceByEmail(ctx, email)
	//TODO: Check if the BruteForce not found create a new BruteForce
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return CreateNewBruteForce(ctx, store, email)
		}
		return db.BlockBruteForce{}, err
	} else if bruteForce.Count.Int32 >= 4 {
		//TODO: If the BruteForce there have been more than 5 incorrect login attempts, Lock the BruteForce.
		return lockBruteForce(ctx, store, bruteForce)
	}

	//TODO: If the BruteForce is not locked, increment the count
	return incrementBruteForce(ctx, store, bruteForce)
}

func CreateNewBruteForce(ctx *gin.Context, store db.Store, email string) (db.BlockBruteForce, error) {
	newBruteForce, err := store.CreateBlockBruteForce(ctx, db.CreateBlockBruteForceParams{
		Email:      email,
		Count:      pgtype.Int4{Int32: 1, Valid: true},
		Status:     constant.UnlockedBruteForce,
		LockedTime: pgtype.Timestamptz{Valid: false},
		UnLockTime: pgtype.Timestamptz{Valid: false},
		CreatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		CreatedBy:  pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		return db.BlockBruteForce{}, err
	}
	return newBruteForce, nil
}

func lockBruteForce(ctx *gin.Context, store db.Store, bruteForce db.BlockBruteForce) (db.BlockBruteForce, error) {
	//TODO: Lock BruteForce
	updatedBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
		ID:         bruteForce.Id,
		Email:      pgtype.Text{String: bruteForce.Email, Valid: true},
		Count:      pgtype.Int4{Int32: bruteForce.Count.Int32 + 1, Valid: true},
		Status:     pgtype.Text{String: constant.LockedBruteForce, Valid: true},
		Lockedtime: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Unlocktime: pgtype.Timestamptz{Time: time.Now().Add(15 * time.Minute), Valid: true},
		Updatedby:  pgtype.Text{String: "system", Valid: true},
		Updatedat:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return db.BlockBruteForce{}, err
	}

	//TODO: Lock account
	if _, err = lockedAccount(ctx, store, bruteForce.Email); err != nil {
		return db.BlockBruteForce{}, err
	}

	return updatedBruteForce, nil
}

func checkForUnLockBruteForce(ctx *gin.Context, store db.Store, email string) (db.BlockBruteForce, error) {
	bruteForce, err := store.GetBlockBruteForceByEmail(ctx, email)
	if err != nil && !errors.Is(err, db.ErrRecordNotFound) {
		return db.BlockBruteForce{}, err
	}

	//TODO: Unlock BruteForce and clear the count
	if bruteForce.Status == constant.LockedBruteForce && bruteForce.UnLockTime.Time.Before(time.Now()) {
		updatedBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
			ID:         bruteForce.Id,
			Count:      pgtype.Int4{Int32: 0, Valid: true},
			Status:     pgtype.Text{String: constant.UnlockedBruteForce, Valid: true},
			Lockedtime: pgtype.Timestamptz{Valid: false},
			Unlocktime: pgtype.Timestamptz{Valid: false},
			Updatedby:  pgtype.Text{String: "system", Valid: true},
			Updatedat:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
		if err != nil {
			return db.BlockBruteForce{}, err
		}

		//TODO: Unlock account
		if _, err = unLockAccount(ctx, store, bruteForce.Email); err != nil && !errors.Is(err, db.ErrRecordNotFound) {
			return db.BlockBruteForce{}, err
		}

		return updatedBruteForce, nil
	}

	return db.BlockBruteForce{}, nil
}

func incrementBruteForce(ctx *gin.Context, store db.Store, bruteForce db.BlockBruteForce) (db.BlockBruteForce, error) {
	newBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
		ID:     bruteForce.Id,
		Email:  pgtype.Text{String: bruteForce.Email, Valid: true},
		Status: pgtype.Text{String: bruteForce.Status, Valid: true},
		Lockedtime: pgtype.Timestamptz{
			Time:  bruteForce.LockedTime.Time,
			Valid: bruteForce.LockedTime.Valid,
		},
		Unlocktime: pgtype.Timestamptz{
			Time:  bruteForce.UnLockTime.Time,
			Valid: bruteForce.UnLockTime.Valid,
		},
		Count:     pgtype.Int4{Int32: bruteForce.Count.Int32 + 1, Valid: true},
		Updatedat: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Updatedby: pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		return db.BlockBruteForce{}, err
	}

	return newBruteForce, nil
}
