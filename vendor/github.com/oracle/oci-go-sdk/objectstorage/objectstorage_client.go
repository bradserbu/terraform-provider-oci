// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oracle/oci-go-sdk/common"
)

//ObjectStorageClient a client for ObjectStorage
type ObjectStorageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

func buildSigner(configProvider common.ConfigurationProvider) common.HTTPRequestSigner {
	objStorageHeaders := []string{"date", "(request-target)", "host"}
	defaultBodyHeaders := []string{"content-length", "content-type", "x-content-sha256"}
	shouldHashBody := func(r *http.Request) bool {
		return r.Method == http.MethodPost
	}
	signer := common.RequestSignerWithBodyHashingPredicate(configProvider, objStorageHeaders, defaultBodyHeaders, shouldHashBody)
	return signer
}

// NewObjectStorageClientWithConfigurationProvider Creates a new default ObjectStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewObjectStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ObjectStorageClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}
	baseClient.Signer = buildSigner(configProvider)

	client = ObjectStorageClient{BaseClient: baseClient}
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ObjectStorageClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "objectstorage", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ObjectStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ObjectStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortMultipartUpload Aborts an in-progress multipart upload and deletes all parts that have been uploaded.
func (client ObjectStorageClient) AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest, options ...common.RetryPolicyOption) (response AbortMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CommitMultipartUpload Commits a multipart upload, which involves checking part numbers and ETags of the parts, to create an aggregate object.
func (client ObjectStorageClient) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest, options ...common.RetryPolicyOption) (response CommitMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateBucket Creates a bucket in the given namespace with a bucket name and optional user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) CreateBucket(ctx context.Context, request CreateBucketRequest, options ...common.RetryPolicyOption) (response CreateBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateMultipartUpload Starts a new multipart upload to a specific object in the given bucket in the given namespace.
func (client ObjectStorageClient) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest, options ...common.RetryPolicyOption) (response CreateMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreatePreauthenticatedRequest Create a pre-authenticated request specific to the bucket
func (client ObjectStorageClient) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest, options ...common.RetryPolicyOption) (response CreatePreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/p/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteBucket Deletes a bucket if it is already empty. If the bucket is not empty, use DeleteObject first.
func (client ObjectStorageClient) DeleteBucket(ctx context.Context, request DeleteBucketRequest, options ...common.RetryPolicyOption) (response DeleteBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteObject Deletes an object.
func (client ObjectStorageClient) DeleteObject(ctx context.Context, request DeleteObjectRequest, options ...common.RetryPolicyOption) (response DeleteObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeletePreauthenticatedRequest Deletes the bucket level pre-authenticateted request
func (client ObjectStorageClient) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest, options ...common.RetryPolicyOption) (response DeletePreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/p/{parId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetBucket Gets the current representation of the given bucket in the given namespace.
func (client ObjectStorageClient) GetBucket(ctx context.Context, request GetBucketRequest, options ...common.RetryPolicyOption) (response GetBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetNamespace Gets the name of the namespace for the user making the request. An account name must be unique, must start with a
// letter, and can have up to 15 lowercase letters and numbers. You cannot use spaces or special characters.
func (client ObjectStorageClient) GetNamespace(ctx context.Context, request GetNamespaceRequest, options ...common.RetryPolicyOption) (response GetNamespaceResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetObject Gets the metadata and body of an object.
func (client ObjectStorageClient) GetObject(ctx context.Context, request GetObjectRequest, options ...common.RetryPolicyOption) (response GetObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		KeepResponseBodyOpen: true,
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetPreauthenticatedRequest Get the bucket level pre-authenticateted request
func (client ObjectStorageClient) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest, options ...common.RetryPolicyOption) (response GetPreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/{parId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// HeadBucket Efficiently checks if a bucket exists and gets the current ETag for the bucket.
func (client ObjectStorageClient) HeadBucket(ctx context.Context, request HeadBucketRequest, options ...common.RetryPolicyOption) (response HeadBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// HeadObject Gets the user-defined metadata and entity tag for an object.
func (client ObjectStorageClient) HeadObject(ctx context.Context, request HeadObjectRequest, options ...common.RetryPolicyOption) (response HeadObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListBuckets Gets a list of all `BucketSummary`s in a compartment. A `BucketSummary` contains only summary fields for the bucket
// and does not contain fields like the user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListBuckets(ctx context.Context, request ListBucketsRequest, options ...common.RetryPolicyOption) (response ListBucketsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListMultipartUploadParts Lists the parts of an in-progress multipart upload.
func (client ObjectStorageClient) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest, options ...common.RetryPolicyOption) (response ListMultipartUploadPartsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListMultipartUploads Lists all in-progress multipart uploads for the given bucket in the given namespace.
func (client ObjectStorageClient) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest, options ...common.RetryPolicyOption) (response ListMultipartUploadsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListObjects Lists the objects in a bucket.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListObjects(ctx context.Context, request ListObjectsRequest, options ...common.RetryPolicyOption) (response ListObjectsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListPreauthenticatedRequests List pre-authenticated requests for the bucket
func (client ObjectStorageClient) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest, options ...common.RetryPolicyOption) (response ListPreauthenticatedRequestsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// PutObject Creates a new object or overwrites an existing one.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) PutObject(ctx context.Context, request PutObjectRequest, options ...common.RetryPolicyOption) (response PutObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateBucket Performs a partial or full update of a bucket's user-defined metadata.
func (client ObjectStorageClient) UpdateBucket(ctx context.Context, request UpdateBucketRequest, options ...common.RetryPolicyOption) (response UpdateBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UploadPart Uploads a single part of a multipart upload.
func (client ObjectStorageClient) UploadPart(ctx context.Context, request UploadPartRequest, options ...common.RetryPolicyOption) (response UploadPartResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}
