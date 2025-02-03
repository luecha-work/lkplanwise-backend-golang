package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/lkplanwise-api/db/sqlc"
)

func CheckBlockedBruteForce(ctx *gin.Context, store db.Store, username string) (bool, error) {
	bruteForce, err := store.GetBlockBruteForceByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	if bruteForce.Status == "L" && bruteForce.UnLockTime.Time.After(time.Now()) {
		remainingTime := time.Until(bruteForce.UnLockTime.Time)
		return true, errors.New("Account is locked, please try again in " + remainingTime.String())
	}

	return false, nil
}

func ManageBlockBruteForce(
	ctx *gin.Context,
	store db.Store,
	username string) (db.BlockBruteForce, error) {

	bruteForce, err := store.GetBlockBruteForceByUsername(ctx, username)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return CreateNewBruteForce(ctx, store, username)
		}
		return db.BlockBruteForce{}, err
	} else if bruteForce.Count.Int32 >= 4 {
		return lockBruteForce(ctx, store, bruteForce)
	}

	return incrementBruteForce(ctx, store, bruteForce)
}

func CreateNewBruteForce(ctx *gin.Context, store db.Store, username string) (db.BlockBruteForce, error) {
	newBruteForce, err := store.CreateBlockBruteForce(ctx, db.CreateBlockBruteForceParams{
		UserName:   username,
		Count:      pgtype.Int4{Int32: 1, Valid: true},
		Status:     "U",
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
	updatedBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
		ID:         bruteForce.Id,
		Username:   pgtype.Text{String: bruteForce.UserName, Valid: true},
		Count:      pgtype.Int4{Int32: bruteForce.Count.Int32 + 1, Valid: true},
		Status:     pgtype.Text{String: "L", Valid: true},
		Lockedtime: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Unlocktime: pgtype.Timestamptz{Time: time.Now().Add(15 * time.Minute), Valid: true},
		Updatedby:  pgtype.Text{String: "system", Valid: true},
		Updatedat:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return db.BlockBruteForce{}, err
	}

	if _, err = lockAccount(ctx, store, bruteForce.UserName); err != nil {
		return db.BlockBruteForce{}, err
	}

	return updatedBruteForce, nil
}

func checkForUnLockBruteForce(ctx *gin.Context, store db.Store, username string) (db.BlockBruteForce, error) {
	bruteForce, err := store.GetBlockBruteForceByUsername(ctx, username)
	if err != nil && !errors.Is(err, db.ErrRecordNotFound) {
		return db.BlockBruteForce{}, err
	}

	if bruteForce.Status == "L" && bruteForce.UnLockTime.Time.Before(time.Now()) {
		updatedBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
			ID:         bruteForce.Id,
			Count:      pgtype.Int4{Int32: 0, Valid: true},
			Status:     pgtype.Text{String: "U", Valid: true},
			Lockedtime: pgtype.Timestamptz{Valid: false},
			Unlocktime: pgtype.Timestamptz{Valid: false},
			Updatedby:  pgtype.Text{String: "system", Valid: true},
			Updatedat:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
		if err != nil {
			return db.BlockBruteForce{}, err
		}

		return updatedBruteForce, nil
	}

	return db.BlockBruteForce{}, nil
}

func incrementBruteForce(ctx *gin.Context, store db.Store, bruteForce db.BlockBruteForce) (db.BlockBruteForce, error) {
	newBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
		ID:       bruteForce.Id,
		Username: pgtype.Text{String: bruteForce.UserName, Valid: true},
		Status:   pgtype.Text{String: bruteForce.Status, Valid: true},
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
