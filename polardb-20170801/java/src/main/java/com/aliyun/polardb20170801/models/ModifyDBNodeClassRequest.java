// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.polardb20170801.models;

import com.aliyun.tea.*;

public class ModifyDBNodeClassRequest extends TeaModel {
    @NameInMap("AccessKeyId")
    public String accessKeyId;

    @NameInMap("OwnerId")
    public Long ownerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("ResourceOwnerId")
    public Long resourceOwnerId;

    @NameInMap("OwnerAccount")
    public String ownerAccount;

    @NameInMap("DBClusterId")
    @Validation(required = true)
    public String DBClusterId;

    @NameInMap("ModifyType")
    @Validation(required = true)
    public String modifyType;

    @NameInMap("DBNodeTargetClass")
    @Validation(required = true)
    public String DBNodeTargetClass;

    @NameInMap("ClientToken")
    public String clientToken;

    public static ModifyDBNodeClassRequest build(java.util.Map<String, ?> map) throws Exception {
        ModifyDBNodeClassRequest self = new ModifyDBNodeClassRequest();
        return TeaModel.build(map, self);
    }

}
