// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170321.models;

import com.aliyun.tea.*;

public class SetDefaultWatermarkResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    public static SetDefaultWatermarkResponse build(java.util.Map<String, ?> map) throws Exception {
        SetDefaultWatermarkResponse self = new SetDefaultWatermarkResponse();
        return TeaModel.build(map, self);
    }

}
