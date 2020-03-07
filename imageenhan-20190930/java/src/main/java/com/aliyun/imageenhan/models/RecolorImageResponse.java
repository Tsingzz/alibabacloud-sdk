// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.imageenhan.models;

import com.aliyun.tea.*;

public class RecolorImageResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("Data")
    @Validation(required = true)
    public RecolorImageResponseData data;

    public static RecolorImageResponse build(java.util.Map<String, ?> map) throws Exception {
        RecolorImageResponse self = new RecolorImageResponse();
        return TeaModel.build(map, self);
    }

    public static class RecolorImageResponseDataImageList extends TeaModel {
        @NameInMap("ImageList")
        @Validation(required = true)
        public String imageList;

        public static RecolorImageResponseDataImageList build(java.util.Map<String, ?> map) throws Exception {
            RecolorImageResponseDataImageList self = new RecolorImageResponseDataImageList();
            return TeaModel.build(map, self);
        }

    }

    public static class RecolorImageResponseData extends TeaModel {
        @NameInMap("ImageList")
        @Validation(required = true)
        public java.util.List<RecolorImageResponseDataImageList> imageList;

        public static RecolorImageResponseData build(java.util.Map<String, ?> map) throws Exception {
            RecolorImageResponseData self = new RecolorImageResponseData();
            return TeaModel.build(map, self);
        }

    }

}
