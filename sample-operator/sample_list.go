package main

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SampleList is a list of samples from the TPR.
type SampleList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	Metadata metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of third party objects
	Items []Sample `json:"items"`
}

// There is known issue with TPR in client-go:
//   https://github.com/kubernetes/client-go/issues/8
// Workarounds:
// - We include `Metadata` field in object explicitly.
// - we have the code below to work around a known problem with third-party resources and ugorji.

type SampleListCopy SampleList
type SampleCopy Sample

func (p *Sample) UnmarshalJSON(data []byte) error {
	tmp := SampleCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := Sample(tmp)
	*p = tmp2
	return nil
}

func (pl *SampleList) UnmarshalJSON(data []byte) error {
	tmp := SampleListCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := SampleList(tmp)
	*pl = tmp2
	return nil
}
