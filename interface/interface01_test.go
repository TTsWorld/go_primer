package _interface

import (
	"fmt"
	"testing"
)

var _ ContributionDao = (*ContributionDaoImpl)(nil)
var DefaultContributionDao ContributionDao

func init() {
	DefaultContributionDao = NewContributionDao()
}

func NewContributionDao() ContributionDao {
	if DefaultContributionDao != nil {
		return DefaultContributionDao
	}
	return &ContributionDaoImpl{}
}
func NewContributionDao02() ContributionDao {
	if DefaultContributionDao != nil {
		return DefaultContributionDao
	}
	return ContributionDaoImpl{}
}

type ContributionDao interface {
	Print()
}
type ContributionDaoImpl struct {
}

func (p ContributionDaoImpl) Print() {
	fmt.Println(p)
}

func NewContributionDao03() *ContributionDaoImpl {
	if DefaultContributionDao != nil {
		return DefaultContributionDao.(*ContributionDaoImpl)
	}
	return &ContributionDaoImpl{}
}

func NewContributionDao04() ContributionDaoImpl {
	if DefaultContributionDao != nil {
		return DefaultContributionDao.(ContributionDaoImpl)
	}
	return ContributionDaoImpl{}
}

func TestI0401(t *testing.T) {
	dd := NewContributionDao()
	dd.Print()

}
