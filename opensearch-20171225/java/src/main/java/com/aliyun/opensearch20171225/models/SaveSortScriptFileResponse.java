// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.opensearch20171225.models;

import com.aliyun.tea.*;

public class SaveSortScriptFileResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public SaveSortScriptFileResponseBody body;

    public static SaveSortScriptFileResponse build(java.util.Map<String, ?> map) throws Exception {
        SaveSortScriptFileResponse self = new SaveSortScriptFileResponse();
        return TeaModel.build(map, self);
    }

}
