package users

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Directory struct {
	logger  *logrus.Logger
	db      *sql.DB
	sb      squirrel.StatementBuilderType
	querier Querier
}

func NewDirectory(logger *logrus.Logger, db *sql.DB) *Directory {
	return &Directory{
		logger:  logger,
		db:      db,
		sb:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
		querier: New(db),
	}
}

func (d *Directory) Close() error {
	return d.db.Close()
}

func (d *Directory) AddExperiment(ctx context.Context, arg AddExperimentParams) (Experiment, error) {
	experiment, err := d.querier.AddExperiment(ctx, arg)
	if err != nil {
		return experiment, status.Errorf(codes.Internal,
			"unexpected error AddExperiment: %s", err.Error())
	}
	return experiment, nil
}

func (d *Directory) UpdateExperiment(ctx context.Context, arg UpdateExperimentParams) (Experiment, error) {
	experiment, err := d.querier.UpdateExperiment(ctx, arg)
	if err != nil {
		return experiment, status.Errorf(codes.Internal, "unexpected error UpdateExperiment: %s", err.Error())
	}
	return experiment, nil
}

func (d *Directory) GetExperiments(ctx context.Context) ([]Experiment, error) {
	experiments, err := d.querier.GetExperiments(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetExperiments: %s", err.Error())
	}
	return experiments, nil
}

func (d *Directory) GetExperiment(ctx context.Context, id int32) (Experiment, error) {
	experiment, err := d.querier.GetExperiment(ctx, id)
	if err != nil {
		return experiment, status.Errorf(codes.Internal, "unexpected error GetExperiment: %s", err.Error())
	}
	return experiment, nil
}

func (d *Directory) GetExperimentByName(ctx context.Context, name string) (Experiment, error) {
	experiment, err := d.querier.GetExperimentByName(ctx, name)
	if err != nil {
		return experiment, status.Errorf(codes.Internal, "unexpected error GetExperimentByName: %s", err.Error())
	}
	return experiment, nil
}

func (d *Directory) GetExperimentsByActiveState(ctx context.Context, activeState bool) ([]Experiment, error) {
	experiments, err := d.querier.GetExperimentsByActiveState(ctx, activeState)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetExperimentByName: %s", err.Error())
	}
	return experiments, nil
}

func (d *Directory) GetExperimentsBySameConditions(ctx context.Context,
	arg GetExperimentsBySameConditionsParams) ([]Experiment, error) {
	experiments, err := d.querier.GetExperimentsBySameConditions(ctx, arg)
	if err != nil {
		return nil,
			status.Errorf(codes.Internal, "unexpected error GetExperimentsBySameConditions: %s",
				err.Error())
	}
	return experiments, nil
}

func (d *Directory) DeleteExperiment(ctx context.Context, id int32) (Experiment, error) {
	exp, err := d.querier.DeleteExperiment(ctx, id)
	if err != nil {
		return exp, status.Errorf(codes.Internal, "unexpected error DeleteExperiment: %s", err.Error())
	}
	return exp, nil
}

func (d *Directory) GetExperimentsExceeded(ctx context.Context, endTime time.Time) ([]Experiment, error) {
	experiments, err := d.querier.GetExperimentsExceeded(ctx, endTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetExperimentsExceeded: %s", err.Error())
	}
	return experiments, nil
}

func (d *Directory) DisableExperiment(ctx context.Context, id int32) (Experiment, error) {
	experiment, err := d.querier.DisableExperiment(ctx, id)
	if err != nil {
		return experiment,
			status.Errorf(codes.Internal, "unexpected error DisableExperiment: %s", err.Error())
	}
	return experiment, nil
}

func (d *Directory) EnableExperiment(ctx context.Context, id int32) (Experiment, error) {
	experiment, err := d.querier.EnableExperiment(ctx, id)
	if err != nil {
		return experiment,
			status.Errorf(codes.Internal, "unexpected error EnableExperiment: %s", err.Error())
	}
	return experiment, nil
}
