// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.nas20170626.models;

import com.aliyun.tea.*;

public class CreateAccessRuleResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("AccessRuleId")
    @Validation(required = true)
    public String accessRuleId;

    public static CreateAccessRuleResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessRuleResponse self = new CreateAccessRuleResponse();
        return TeaModel.build(map, self);
    }

}
