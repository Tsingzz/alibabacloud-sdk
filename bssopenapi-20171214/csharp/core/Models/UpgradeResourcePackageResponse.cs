// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.BssOpenApi20171214.Models
{
    public class UpgradeResourcePackageResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("OrderId")]
        [Validation(Required=true)]
        public long OrderId { get; set; }

        [NameInMap("Success")]
        [Validation(Required=true)]
        public bool? Success { get; set; }

        [NameInMap("Code")]
        [Validation(Required=true)]
        public string Code { get; set; }

        [NameInMap("Message")]
        [Validation(Required=true)]
        public string Message { get; set; }

        [NameInMap("Data")]
        [Validation(Required=true)]
        public UpgradeResourcePackageResponseData Data { get; set; }
        public class UpgradeResourcePackageResponseData : TeaModel {
            [NameInMap("OrderId")]
            [Validation(Required=true)]
            public long OrderId { get; set; }
            [NameInMap("InstanceId")]
            [Validation(Required=true)]
            public string InstanceId { get; set; }
        };

    }

}
