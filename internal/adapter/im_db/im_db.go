package im_db

import (
	"context"
	"testWorkmate/internal/common/types/error_with_codes"
	"testWorkmate/internal/model/task_model"
)

// ImitationDb is a mock database for testing purposes.
type ImitationDb struct {
	table map[uint64]task_model.Task
}

func NewImitationDb() *ImitationDb {
	return &ImitationDb{
		table: make(map[uint64]task_model.Task),
	}
}

func (m *ImitationDb) Exec(ctx context.Context, values *task_model.Task, dest *task_model.Task) error {
	m.table[values.ID] = *values
	*dest = *values
	return nil
}

func (m *ImitationDb) Get(ctx context.Context, dest *task_model.Task, id uint64) error {
	task, ok := m.table[id]
	if !ok {
		err := error_with_codes.ErrorTaskNotFound
		return err
	}
	*dest = task
	return nil
}

func (m *ImitationDb) Delete(ctx context.Context, id uint64) error {
	_, ok := m.table[id]
	if !ok {
		err := error_with_codes.ErrorTaskNotFound
		return err
	}
	delete(m.table, id)
	return nil
}
