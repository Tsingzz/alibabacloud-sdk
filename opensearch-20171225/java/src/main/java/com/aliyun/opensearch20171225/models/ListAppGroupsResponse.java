// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.opensearch20171225.models;

import com.aliyun.tea.*;

public class ListAppGroupsResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public ListAppGroupsResponseBody body;

    public static ListAppGroupsResponse build(java.util.Map<String, ?> map) throws Exception {
        ListAppGroupsResponse self = new ListAppGroupsResponse();
        return TeaModel.build(map, self);
    }

}
