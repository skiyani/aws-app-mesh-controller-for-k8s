package virtualservice

import (
	appmesh "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"testing"
)

func TestVirtualNodeReferenceIndexFunc(t *testing.T) {
	type args struct {
		obj runtime.Object
	}
	tests := []struct {
		name string
		args args
		want []types.NamespacedName
	}{
		{
			name: "using virtualNodeProvider - with namespace",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: &appmesh.VirtualServiceProvider{
							VirtualNode: &appmesh.VirtualNodeServiceProvider{
								VirtualNodeRef: appmesh.VirtualNodeReference{
									Namespace: aws.String("other-ns"),
									Name:      "vn",
								},
							},
						},
					},
				},
			},
			want: []types.NamespacedName{
				{
					Namespace: "other-ns",
					Name:      "vn",
				},
			},
		},
		{
			name: "using virtualNodeProvider - without namespace",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: &appmesh.VirtualServiceProvider{
							VirtualNode: &appmesh.VirtualNodeServiceProvider{
								VirtualNodeRef: appmesh.VirtualNodeReference{
									Name: "vn",
								},
							},
						},
					},
				},
			},
			want: []types.NamespacedName{
				{
					Namespace: "my-ns",
					Name:      "vn",
				},
			},
		},
		{
			name: "using virtualRouterProvider",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: &appmesh.VirtualServiceProvider{
							VirtualRouter: &appmesh.VirtualRouterServiceProvider{
								VirtualRouterRef: appmesh.VirtualRouterReference{
									Namespace: aws.String("other-ns"),
									Name:      "vr",
								},
							},
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "using no provider",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: nil,
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VirtualNodeReferenceIndexFunc(tt.args.obj)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVirtualRouterReferenceIndexFunc(t *testing.T) {
	type args struct {
		obj runtime.Object
	}
	tests := []struct {
		name string
		args args
		want []types.NamespacedName
	}{
		{
			name: "using virtualRouterProvider - with namespace",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: &appmesh.VirtualServiceProvider{
							VirtualRouter: &appmesh.VirtualRouterServiceProvider{
								VirtualRouterRef: appmesh.VirtualRouterReference{
									Namespace: aws.String("other-ns"),
									Name:      "vr",
								},
							},
						},
					},
				},
			},
			want: []types.NamespacedName{
				{
					Namespace: "other-ns",
					Name:      "vr",
				},
			},
		},
		{
			name: "using virtualRouterProvider - without namespace",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: &appmesh.VirtualServiceProvider{
							VirtualRouter: &appmesh.VirtualRouterServiceProvider{
								VirtualRouterRef: appmesh.VirtualRouterReference{
									Name: "vr",
								},
							},
						},
					},
				},
			},
			want: []types.NamespacedName{
				{
					Namespace: "my-ns",
					Name:      "vr",
				},
			},
		},
		{
			name: "using virtualNodeProvider",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: &appmesh.VirtualServiceProvider{
							VirtualNode: &appmesh.VirtualNodeServiceProvider{
								VirtualNodeRef: appmesh.VirtualNodeReference{
									Namespace: aws.String("other-ns"),
									Name:      "vn",
								},
							},
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "using no provider",
			args: args{
				obj: &appmesh.VirtualService{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "my-ns",
					},
					Spec: appmesh.VirtualServiceSpec{
						Provider: nil,
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VirtualRouterReferenceIndexFunc(tt.args.obj)
			assert.Equal(t, tt.want, got)
		})
	}
}
