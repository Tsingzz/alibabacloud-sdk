// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.CS20151215.Models
{
    public class DeleteClusterNodesBody : TeaModel {
        [NameInMap("release_node")]
        [Validation(Required=false)]
        public string ReleaseNode { get; set; }

        [NameInMap("nodes")]
        [Validation(Required=false)]
        public List<string> Nodes { get; set; }

    }

}
