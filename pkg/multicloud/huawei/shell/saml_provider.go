// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/onecloud/pkg/multicloud/huawei"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type SAMLProviderListOptions struct {
	}
	shellutils.R(&SAMLProviderListOptions{}, "saml-provider-list", "List saml provider", func(cli *huawei.SRegion, args *SAMLProviderListOptions) error {
		result, err := cli.GetClient().ListSAMLProviders()
		if err != nil {
			return err
		}
		printList(result, 0, 0, 0, nil)
		return nil
	})

	type SAMLProviderIdOptions struct {
		ID string
	}

	shellutils.R(&SAMLProviderIdOptions{}, "saml-provider-delete", "Delete saml provider", func(cli *huawei.SRegion, args *SAMLProviderIdOptions) error {
		return cli.GetClient().DeleteSAMLProvider(args.ID)
	})

	shellutils.R(&SAMLProviderIdOptions{}, "saml-provider-protocol-list", "List saml provider protocol", func(cli *huawei.SRegion, args *SAMLProviderIdOptions) error {
		result, err := cli.GetClient().GetSAMLProviderProtocols(args.ID)
		if err != nil {
			return err
		}
		printList(result, 0, 0, 0, nil)
		return nil
	})

	shellutils.R(&SAMLProviderIdOptions{}, "saml-provider-metadata-show", "Show saml provider metadata", func(cli *huawei.SRegion, args *SAMLProviderIdOptions) error {
		result, err := cli.GetClient().GetSAMLProviderMetadata(args.ID)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	type SAMLProviderMetadataOptions struct {
		ID       string
		METADATA string
	}

	shellutils.R(&SAMLProviderMetadataOptions{}, "saml-provider-metadata-update", "Update saml provider metadata", func(cli *huawei.SRegion, args *SAMLProviderMetadataOptions) error {
		return cli.GetClient().UpdateSAMLProviderMetadata(args.ID, args.METADATA)
	})

	type MappingListOptions struct {
	}

	shellutils.R(&MappingListOptions{}, "saml-provider-mapping-list", "List saml provider mapping", func(cli *huawei.SRegion, args *MappingListOptions) error {
		mappings, err := cli.GetClient().ListSAMLProviderMappings()
		if err != nil {
			return err
		}
		printList(mappings, 0, 0, 0, nil)
		return nil
	})

	type MappingInitOptions struct {
		SAML_PROVIDER string
	}

	shellutils.R(&MappingInitOptions{}, "saml-provider-mapping-init", "Init saml provider mapping", func(cli *huawei.SRegion, args *MappingInitOptions) error {
		return cli.GetClient().InitSAMLProviderMapping(args.SAML_PROVIDER)
	})

}
