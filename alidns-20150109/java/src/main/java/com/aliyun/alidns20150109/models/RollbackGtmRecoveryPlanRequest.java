// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.alidns20150109.models;

import com.aliyun.tea.*;

public class RollbackGtmRecoveryPlanRequest extends TeaModel {
    @NameInMap("Lang")
    public String lang;

    @NameInMap("RecoveryPlanId")
    @Validation(required = true)
    public Long recoveryPlanId;

    public static RollbackGtmRecoveryPlanRequest build(java.util.Map<String, ?> map) throws Exception {
        RollbackGtmRecoveryPlanRequest self = new RollbackGtmRecoveryPlanRequest();
        return TeaModel.build(map, self);
    }

}
