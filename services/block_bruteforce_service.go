package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/lkplanwise-api/db/sqlc"
)

func ManageBlockBruteForce(
	ctx *gin.Context,
	store db.Store,
	username string) (db.BlockBruteForce, error) {

	bruteForce, err := store.GetBlockBruteForceByUsername(ctx, username)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return createNewBruteForce(ctx, store, username)
		}
		return db.BlockBruteForce{}, err
	} else if bruteForce.Count.Int32 >= 4 {
		return lockBruteForce(ctx, store, bruteForce)
	}

	return incrementBruteForce(ctx, store, bruteForce)
}

func createNewBruteForce(ctx *gin.Context, store db.Store, username string) (db.BlockBruteForce, error) {
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
	newBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
		Id:         bruteForce.Id,
		UserName:   bruteForce.UserName,
		Count:      pgtype.Int4{Int32: bruteForce.Count.Int32 + 1, Valid: true},
		Status:     "L",
		LockedTime: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UnLockTime: pgtype.Timestamptz{Time: time.Now().Add(15 * time.Minute), Valid: true},
		UpdatedBy:  pgtype.Text{String: "system", Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return db.BlockBruteForce{}, err
	}
	return newBruteForce, nil
}

func incrementBruteForce(ctx *gin.Context, store db.Store, bruteForce db.BlockBruteForce) (db.BlockBruteForce, error) {
	newBruteForce, err := store.UpdateBlockBruteForce(ctx, db.UpdateBlockBruteForceParams{
		Id:       bruteForce.Id,
		UserName: bruteForce.UserName,
		Status:   bruteForce.Status,
		LockedTime: pgtype.Timestamptz{
			Time:  bruteForce.LockedTime.Time,
			Valid: bruteForce.LockedTime.Valid,
		},
		UnLockTime: pgtype.Timestamptz{
			Time:  bruteForce.UnLockTime.Time,
			Valid: bruteForce.UnLockTime.Valid,
		},
		Count:     pgtype.Int4{Int32: bruteForce.Count.Int32 + 1, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedBy: pgtype.Text{String: "system", Valid: true},
	})
	if err != nil {
		return db.BlockBruteForce{}, err
	}

	return newBruteForce, nil
}
