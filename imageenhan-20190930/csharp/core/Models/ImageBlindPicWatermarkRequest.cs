// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.Imageenhan20190930.Models
{
    public class ImageBlindPicWatermarkRequest : TeaModel {
        [NameInMap("FunctionType")]
        [Validation(Required=true)]
        public string FunctionType { get; set; }

        [NameInMap("LogoURL")]
        [Validation(Required=false)]
        public string LogoURL { get; set; }

        [NameInMap("WatermarkImageURL")]
        [Validation(Required=false)]
        public string WatermarkImageURL { get; set; }

        [NameInMap("OutputFileType")]
        [Validation(Required=false)]
        public string OutputFileType { get; set; }

        [NameInMap("QualityFactor")]
        [Validation(Required=true)]
        public int? QualityFactor { get; set; }

        [NameInMap("OriginImageURL")]
        [Validation(Required=true)]
        public string OriginImageURL { get; set; }

    }

}
