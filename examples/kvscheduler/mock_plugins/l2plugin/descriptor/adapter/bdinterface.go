// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/examples/kvscheduler/mock_plugins/l2plugin/model"
)

////////// type-safe key-value pair with metadata //////////

type BDInterfaceKVWithMetadata struct {
	Key      string
	Value    *mock_l2.BridgeDomain_Interface
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type BDInterfaceDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *mock_l2.BridgeDomain_Interface) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *mock_l2.BridgeDomain_Interface) error
	Create               func(key string, value *mock_l2.BridgeDomain_Interface) (metadata interface{}, err error)
	Delete               func(key string, value *mock_l2.BridgeDomain_Interface, metadata interface{}) error
	Update               func(key string, oldValue, newValue *mock_l2.BridgeDomain_Interface, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *mock_l2.BridgeDomain_Interface, metadata interface{}) bool
	Retrieve             func(correlate []BDInterfaceKVWithMetadata) ([]BDInterfaceKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *mock_l2.BridgeDomain_Interface) []KeyValuePair
	Dependencies         func(key string, value *mock_l2.BridgeDomain_Interface) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type BDInterfaceDescriptorAdapter struct {
	descriptor *BDInterfaceDescriptor
}

func NewBDInterfaceDescriptor(typedDescriptor *BDInterfaceDescriptor) *KVDescriptor {
	adapter := &BDInterfaceDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *BDInterfaceDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castBDInterfaceValue(key, oldValue)
	typedNewValue, err2 := castBDInterfaceValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *BDInterfaceDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castBDInterfaceValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *BDInterfaceDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castBDInterfaceValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *BDInterfaceDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castBDInterfaceValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castBDInterfaceValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castBDInterfaceMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *BDInterfaceDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castBDInterfaceValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castBDInterfaceMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *BDInterfaceDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castBDInterfaceValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castBDInterfaceValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castBDInterfaceMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *BDInterfaceDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []BDInterfaceKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castBDInterfaceValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castBDInterfaceMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			BDInterfaceKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *BDInterfaceDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castBDInterfaceValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *BDInterfaceDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castBDInterfaceValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castBDInterfaceValue(key string, value proto.Message) (*mock_l2.BridgeDomain_Interface, error) {
	typedValue, ok := value.(*mock_l2.BridgeDomain_Interface)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castBDInterfaceMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
