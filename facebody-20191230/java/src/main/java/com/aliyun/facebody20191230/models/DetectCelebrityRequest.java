// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.facebody20191230.models;

import com.aliyun.tea.*;

public class DetectCelebrityRequest extends TeaModel {
    @NameInMap("ImageURL")
    @Validation(required = true)
    public String imageURL;

    public static DetectCelebrityRequest build(java.util.Map<String, ?> map) throws Exception {
        DetectCelebrityRequest self = new DetectCelebrityRequest();
        return TeaModel.build(map, self);
    }

}
