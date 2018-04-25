// Code generated by lister-gen. DO NOT EDIT.

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/chargeback/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScheduledReportLister helps list ScheduledReports.
type ScheduledReportLister interface {
	// List lists all ScheduledReports in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledReport, err error)
	// ScheduledReports returns an object that can list and get ScheduledReports.
	ScheduledReports(namespace string) ScheduledReportNamespaceLister
	ScheduledReportListerExpansion
}

// scheduledReportLister implements the ScheduledReportLister interface.
type scheduledReportLister struct {
	indexer cache.Indexer
}

// NewScheduledReportLister returns a new ScheduledReportLister.
func NewScheduledReportLister(indexer cache.Indexer) ScheduledReportLister {
	return &scheduledReportLister{indexer: indexer}
}

// List lists all ScheduledReports in the indexer.
func (s *scheduledReportLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledReport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledReport))
	})
	return ret, err
}

// ScheduledReports returns an object that can list and get ScheduledReports.
func (s *scheduledReportLister) ScheduledReports(namespace string) ScheduledReportNamespaceLister {
	return scheduledReportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduledReportNamespaceLister helps list and get ScheduledReports.
type ScheduledReportNamespaceLister interface {
	// List lists all ScheduledReports in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledReport, err error)
	// Get retrieves the ScheduledReport from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ScheduledReport, error)
	ScheduledReportNamespaceListerExpansion
}

// scheduledReportNamespaceLister implements the ScheduledReportNamespaceLister
// interface.
type scheduledReportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScheduledReports in the indexer for a given namespace.
func (s scheduledReportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledReport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledReport))
	})
	return ret, err
}

// Get retrieves the ScheduledReport from the indexer for a given namespace and name.
func (s scheduledReportNamespaceLister) Get(name string) (*v1alpha1.ScheduledReport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scheduledreport"), name)
	}
	return obj.(*v1alpha1.ScheduledReport), nil
}
