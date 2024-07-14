package pgstore

import (
	"context"
	"fmt"
	"journey/internal/api/spec"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (q *Queries) CreateNewTrip(ctx context.Context, pool *pgxpool.Pool, params spec.CreateNewTripRequest) (uuid.UUID, error) {
	tx, err := pool.Begin(ctx)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to begin trx for CreateTrip: %w", err)
	}

	defer func() { _ = tx.Rollback(ctx) }()

	qtx := q.WithTx(tx)

	tripID, err := qtx.InsertTrip(ctx, InsertTripParams{
		Destination: params.Destination,
		OwnerEmail: string(params.OwnerEmail),
		OwnerName: params.OwnerName,
		StartsAt: pgtype.Timestamp{Valid: true, Time: params.StartsAt},
		EndsAt: pgtype.Timestamp{Valid: true, Time: params.EndsAt},
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to insert trip for CreateTrip: %w", err)
	}

	participants := make([]InviteParticipantsToTripParams, len(params.EmailsToInvite))

	for i, email := range params.EmailsToInvite {
		participants[i] = InviteParticipantsToTripParams{
			TripID: tripID,
			Email: string(email),
		}
	}

	if _, err := qtx.InviteParticipantsToTrip(ctx, participants); err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to insert participants for CreateTrip: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.UUID{}, fmt.Errorf("pgstore: failed to commit tx for CreateTrip: %w", err)
	}

	return tripID, nil
}