// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.bssopenapi20171214.models;

import com.aliyun.tea.*;

public class QuerySplitItemBillResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("Success")
    @Validation(required = true)
    public Boolean success;

    @NameInMap("Code")
    @Validation(required = true)
    public String code;

    @NameInMap("Message")
    @Validation(required = true)
    public String message;

    @NameInMap("Data")
    @Validation(required = true)
    public QuerySplitItemBillResponseData data;

    public static QuerySplitItemBillResponse build(java.util.Map<String, ?> map) throws Exception {
        QuerySplitItemBillResponse self = new QuerySplitItemBillResponse();
        return TeaModel.build(map, self);
    }

    public static class QuerySplitItemBillResponseDataItemsItem extends TeaModel {
        @NameInMap("InstanceID")
        @Validation(required = true)
        public String instanceID;

        @NameInMap("BillingType")
        @Validation(required = true)
        public String billingType;

        @NameInMap("CostUnit")
        @Validation(required = true)
        public String costUnit;

        @NameInMap("ProductCode")
        @Validation(required = true)
        public String productCode;

        @NameInMap("ProductType")
        @Validation(required = true)
        public String productType;

        @NameInMap("SubscriptionType")
        @Validation(required = true)
        public String subscriptionType;

        @NameInMap("ProductName")
        @Validation(required = true)
        public String productName;

        @NameInMap("ProductDetail")
        @Validation(required = true)
        public String productDetail;

        @NameInMap("OwnerID")
        @Validation(required = true)
        public String ownerID;

        @NameInMap("BillingItem")
        @Validation(required = true)
        public String billingItem;

        @NameInMap("ListPrice")
        @Validation(required = true)
        public String listPrice;

        @NameInMap("ListPriceUnit")
        @Validation(required = true)
        public String listPriceUnit;

        @NameInMap("Usage")
        @Validation(required = true)
        public String usage;

        @NameInMap("UsageUnit")
        @Validation(required = true)
        public String usageUnit;

        @NameInMap("DeductedByResourcePackage")
        @Validation(required = true)
        public String deductedByResourcePackage;

        @NameInMap("PretaxGrossAmount")
        @Validation(required = true)
        public Double pretaxGrossAmount;

        @NameInMap("InvoiceDiscount")
        @Validation(required = true)
        public Double invoiceDiscount;

        @NameInMap("DeductedByCoupons")
        @Validation(required = true)
        public Double deductedByCoupons;

        @NameInMap("PretaxAmount")
        @Validation(required = true)
        public Double pretaxAmount;

        @NameInMap("DeductedByCashCoupons")
        @Validation(required = true)
        public Double deductedByCashCoupons;

        @NameInMap("DeductedByPrepaidCard")
        @Validation(required = true)
        public Double deductedByPrepaidCard;

        @NameInMap("PaymentAmount")
        @Validation(required = true)
        public Double paymentAmount;

        @NameInMap("OutstandingAmount")
        @Validation(required = true)
        public Double outstandingAmount;

        @NameInMap("Currency")
        @Validation(required = true)
        public String currency;

        @NameInMap("NickName")
        @Validation(required = true)
        public String nickName;

        @NameInMap("ResourceGroup")
        @Validation(required = true)
        public String resourceGroup;

        @NameInMap("Tag")
        @Validation(required = true)
        public String tag;

        @NameInMap("InstanceConfig")
        @Validation(required = true)
        public String instanceConfig;

        @NameInMap("InstanceSpec")
        @Validation(required = true)
        public String instanceSpec;

        @NameInMap("InternetIP")
        @Validation(required = true)
        public String internetIP;

        @NameInMap("IntranetIP")
        @Validation(required = true)
        public String intranetIP;

        @NameInMap("Region")
        @Validation(required = true)
        public String region;

        @NameInMap("Zone")
        @Validation(required = true)
        public String zone;

        @NameInMap("Item")
        @Validation(required = true)
        public String item;

        @NameInMap("ServicePeriod")
        @Validation(required = true)
        public String servicePeriod;

        @NameInMap("BillingDate")
        @Validation(required = true)
        public String billingDate;

        @NameInMap("SplitItemID")
        @Validation(required = true)
        public String splitItemID;

        @NameInMap("SplitItemName")
        @Validation(required = true)
        public String splitItemName;

        public static QuerySplitItemBillResponseDataItemsItem build(java.util.Map<String, ?> map) throws Exception {
            QuerySplitItemBillResponseDataItemsItem self = new QuerySplitItemBillResponseDataItemsItem();
            return TeaModel.build(map, self);
        }

    }

    public static class QuerySplitItemBillResponseDataItems extends TeaModel {
        @NameInMap("Item")
        @Validation(required = true)
        public java.util.List<QuerySplitItemBillResponseDataItemsItem> item;

        public static QuerySplitItemBillResponseDataItems build(java.util.Map<String, ?> map) throws Exception {
            QuerySplitItemBillResponseDataItems self = new QuerySplitItemBillResponseDataItems();
            return TeaModel.build(map, self);
        }

    }

    public static class QuerySplitItemBillResponseData extends TeaModel {
        @NameInMap("BillingCycle")
        @Validation(required = true)
        public String billingCycle;

        @NameInMap("AccountID")
        @Validation(required = true)
        public String accountID;

        @NameInMap("AccountName")
        @Validation(required = true)
        public String accountName;

        @NameInMap("TotalCount")
        @Validation(required = true)
        public Integer totalCount;

        @NameInMap("PageNum")
        @Validation(required = true)
        public Integer pageNum;

        @NameInMap("PageSize")
        @Validation(required = true)
        public Integer pageSize;

        @NameInMap("Items")
        @Validation(required = true)
        public QuerySplitItemBillResponseDataItems items;

        public static QuerySplitItemBillResponseData build(java.util.Map<String, ?> map) throws Exception {
            QuerySplitItemBillResponseData self = new QuerySplitItemBillResponseData();
            return TeaModel.build(map, self);
        }

    }

}
