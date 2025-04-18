// Code generated by MockGen. DO NOT EDIT.
// Source: internal/mocks/kafka_producer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	gomock "github.com/golang/mock/gomock"
)

// MockKafkaProducer is a mock of KafkaProducer interface.
type MockKafkaProducer struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaProducerMockRecorder
}

// MockKafkaProducerMockRecorder is the mock recorder for MockKafkaProducer.
type MockKafkaProducerMockRecorder struct {
	mock *MockKafkaProducer
}

// NewMockKafkaProducer creates a new mock instance.
func NewMockKafkaProducer(ctrl *gomock.Controller) *MockKafkaProducer {
	mock := &MockKafkaProducer{ctrl: ctrl}
	mock.recorder = &MockKafkaProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaProducer) EXPECT() *MockKafkaProducerMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *MockKafkaProducer) Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", msg, deliveryChan)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockKafkaProducerMockRecorder) Produce(msg, deliveryChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockKafkaProducer)(nil).Produce), msg, deliveryChan)
}
