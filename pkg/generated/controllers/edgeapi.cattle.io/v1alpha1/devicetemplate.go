/*
Copyright 2020 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/cnrancher/edge-api-server/pkg/apis/edgeapi.cattle.io/v1alpha1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type DeviceTemplateHandler func(string, *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error)

type DeviceTemplateController interface {
	generic.ControllerMeta
	DeviceTemplateClient

	OnChange(ctx context.Context, name string, sync DeviceTemplateHandler)
	OnRemove(ctx context.Context, name string, sync DeviceTemplateHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() DeviceTemplateCache
}

type DeviceTemplateClient interface {
	Create(*v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error)
	Update(*v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error)
	UpdateStatus(*v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.DeviceTemplate, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha1.DeviceTemplateList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceTemplate, err error)
}

type DeviceTemplateCache interface {
	Get(namespace, name string) (*v1alpha1.DeviceTemplate, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha1.DeviceTemplate, error)

	AddIndexer(indexName string, indexer DeviceTemplateIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha1.DeviceTemplate, error)
}

type DeviceTemplateIndexer func(obj *v1alpha1.DeviceTemplate) ([]string, error)

type deviceTemplateController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewDeviceTemplateController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) DeviceTemplateController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &deviceTemplateController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromDeviceTemplateHandlerToHandler(sync DeviceTemplateHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha1.DeviceTemplate
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha1.DeviceTemplate))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *deviceTemplateController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha1.DeviceTemplate))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateDeviceTemplateDeepCopyOnChange(client DeviceTemplateClient, obj *v1alpha1.DeviceTemplate, handler func(obj *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error)) (*v1alpha1.DeviceTemplate, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *deviceTemplateController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *deviceTemplateController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *deviceTemplateController) OnChange(ctx context.Context, name string, sync DeviceTemplateHandler) {
	c.AddGenericHandler(ctx, name, FromDeviceTemplateHandlerToHandler(sync))
}

func (c *deviceTemplateController) OnRemove(ctx context.Context, name string, sync DeviceTemplateHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromDeviceTemplateHandlerToHandler(sync)))
}

func (c *deviceTemplateController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *deviceTemplateController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *deviceTemplateController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *deviceTemplateController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *deviceTemplateController) Cache() DeviceTemplateCache {
	return &deviceTemplateCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *deviceTemplateController) Create(obj *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error) {
	result := &v1alpha1.DeviceTemplate{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *deviceTemplateController) Update(obj *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error) {
	result := &v1alpha1.DeviceTemplate{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *deviceTemplateController) UpdateStatus(obj *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error) {
	result := &v1alpha1.DeviceTemplate{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *deviceTemplateController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *deviceTemplateController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.DeviceTemplate, error) {
	result := &v1alpha1.DeviceTemplate{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *deviceTemplateController) List(namespace string, opts metav1.ListOptions) (*v1alpha1.DeviceTemplateList, error) {
	result := &v1alpha1.DeviceTemplateList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *deviceTemplateController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *deviceTemplateController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha1.DeviceTemplate, error) {
	result := &v1alpha1.DeviceTemplate{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type deviceTemplateCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *deviceTemplateCache) Get(namespace, name string) (*v1alpha1.DeviceTemplate, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha1.DeviceTemplate), nil
}

func (c *deviceTemplateCache) List(namespace string, selector labels.Selector) (ret []*v1alpha1.DeviceTemplate, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeviceTemplate))
	})

	return ret, err
}

func (c *deviceTemplateCache) AddIndexer(indexName string, indexer DeviceTemplateIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha1.DeviceTemplate))
		},
	}))
}

func (c *deviceTemplateCache) GetByIndex(indexName, key string) (result []*v1alpha1.DeviceTemplate, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha1.DeviceTemplate, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha1.DeviceTemplate))
	}
	return result, nil
}

type DeviceTemplateStatusHandler func(obj *v1alpha1.DeviceTemplate, status v1alpha1.DeviceTemplateStatus) (v1alpha1.DeviceTemplateStatus, error)

type DeviceTemplateGeneratingHandler func(obj *v1alpha1.DeviceTemplate, status v1alpha1.DeviceTemplateStatus) ([]runtime.Object, v1alpha1.DeviceTemplateStatus, error)

func RegisterDeviceTemplateStatusHandler(ctx context.Context, controller DeviceTemplateController, condition condition.Cond, name string, handler DeviceTemplateStatusHandler) {
	statusHandler := &deviceTemplateStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromDeviceTemplateHandlerToHandler(statusHandler.sync))
}

func RegisterDeviceTemplateGeneratingHandler(ctx context.Context, controller DeviceTemplateController, apply apply.Apply,
	condition condition.Cond, name string, handler DeviceTemplateGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &deviceTemplateGeneratingHandler{
		DeviceTemplateGeneratingHandler: handler,
		apply:                           apply,
		name:                            name,
		gvk:                             controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterDeviceTemplateStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type deviceTemplateStatusHandler struct {
	client    DeviceTemplateClient
	condition condition.Cond
	handler   DeviceTemplateStatusHandler
}

func (a *deviceTemplateStatusHandler) sync(key string, obj *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		var newErr error
		obj.Status = newStatus
		obj, newErr = a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
	}
	return obj, err
}

type deviceTemplateGeneratingHandler struct {
	DeviceTemplateGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *deviceTemplateGeneratingHandler) Remove(key string, obj *v1alpha1.DeviceTemplate) (*v1alpha1.DeviceTemplate, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1alpha1.DeviceTemplate{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *deviceTemplateGeneratingHandler) Handle(obj *v1alpha1.DeviceTemplate, status v1alpha1.DeviceTemplateStatus) (v1alpha1.DeviceTemplateStatus, error) {
	objs, newStatus, err := a.DeviceTemplateGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
