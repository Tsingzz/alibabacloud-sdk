// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.cloudwifi_pop20191118.models;

import com.aliyun.tea.*;

public class EffectApConfigRequest extends TeaModel {
    @NameInMap("AppName")
    @Validation(required = true)
    public String appName;

    @NameInMap("AppCode")
    @Validation(required = true)
    public String appCode;

    public static EffectApConfigRequest build(java.util.Map<String, ?> map) throws Exception {
        EffectApConfigRequest self = new EffectApConfigRequest();
        return TeaModel.build(map, self);
    }

}
