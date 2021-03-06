// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.Polardb20170801.Models
{
    public class DescribeBackupsResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("TotalRecordCount")]
        [Validation(Required=true)]
        public string TotalRecordCount { get; set; }

        [NameInMap("PageNumber")]
        [Validation(Required=true)]
        public string PageNumber { get; set; }

        [NameInMap("PageRecordCount")]
        [Validation(Required=true)]
        public string PageRecordCount { get; set; }

        [NameInMap("Items")]
        [Validation(Required=true)]
        public DescribeBackupsResponseItems Items { get; set; }
        public class DescribeBackupsResponseItems : TeaModel {
            [NameInMap("Backup")]
            [Validation(Required=true)]
            public List<DescribeBackupsResponseItemsBackup> Backup { get; set; }
            public class DescribeBackupsResponseItemsBackup : TeaModel {
                    public string BackupId { get; set; }
                    public string DBClusterId { get; set; }
                    public string BackupStatus { get; set; }
                    public string BackupStartTime { get; set; }
                    public string BackupEndTime { get; set; }
                    public string BackupType { get; set; }
                    public string BackupMode { get; set; }
                    public string BackupMethod { get; set; }
                    public string StoreStatus { get; set; }
                    public string BackupSetSize { get; set; }
                    public string ConsistentTime { get; set; }
                    public string BackupsLevel { get; set; }
                    public string IsAvail { get; set; }
            }
        };

    }

}
