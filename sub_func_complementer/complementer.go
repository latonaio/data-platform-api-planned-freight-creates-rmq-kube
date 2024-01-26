package sub_func_complementer

import (
	"context"
	dpfm_api_input_reader "data-platform-api-planned-freight-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-planned-freight-creates-rmq-kube/config"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SubFuncComplementer struct {
	ctx context.Context
	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
	db  *database.Mysql
}

func NewSubFuncComplementer(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *SubFuncComplementer {
	return &SubFuncComplementer{
		ctx: ctx,
		c:   c,
		rmq: rmq,
		db:  db,
	}
}

func (c *SubFuncComplementer) ComplementHeader(input *dpfm_api_input_reader.SDC, subfuncSDC *SDC, l *logger.Logger) error {
	s := &SDC{}
	numRange, err := c.ComplementHeader(input, l)
	if err != nil {
		return xerrors.Errorf("complement plannedFreight error: %w", err)
	}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Headers"], input)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}

	err = c.IncrementLatestNumber(numRange, l)
	if err != nil {
		return xerrors.Errorf("increment latest number error: %w")
	}

	subfuncSDC.SubfuncResult = s.SubfuncResult
	subfuncSDC.SubfuncError = s.SubfuncError
	subfuncSDC.Message.Header = s.Message.Header
	return nil
}

func (c *SubFuncComplementer) IncrementLatestNumber(nr *NumberRange, l *logger.Logger) error {
	_, err := c.db.Query(
		`UPDATE DataPlatformCommonSettingsMysqlKube.data_platform_number_range_latest_number_data
		SET LatestNumber = ?
		WHERE  (NumberRangeID, ServiceLabel, FieldNameWithNumberRange) = ((?,?,?));`, nr.LatestNumber, nr.NumberRangeID, nr.ServiceLabel, nr.FieldNameWithNumberRange,
	)
	if err != nil {
		return xerrors.Errorf("DB Query error: %w", err)
	}

	return nil
}
