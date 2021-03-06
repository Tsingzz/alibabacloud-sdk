// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170321.models;

import com.aliyun.tea.*;

public class GetMediaAuditAudioResultDetailRequest extends TeaModel {
    @NameInMap("AccessKeyId")
    public String accessKeyId;

    @NameInMap("OwnerId")
    public String ownerId;

    @NameInMap("ResourceOwnerId")
    public String resourceOwnerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("OwnerAccount")
    public String ownerAccount;

    @NameInMap("MediaId")
    @Validation(required = true)
    public String mediaId;

    @NameInMap("PageNo")
    public Integer pageNo;

    public static GetMediaAuditAudioResultDetailRequest build(java.util.Map<String, ?> map) throws Exception {
        GetMediaAuditAudioResultDetailRequest self = new GetMediaAuditAudioResultDetailRequest();
        return TeaModel.build(map, self);
    }

}
