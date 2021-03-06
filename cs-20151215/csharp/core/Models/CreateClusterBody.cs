// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.CS20151215.Models
{
    public class CreateClusterBody : TeaModel {
        [NameInMap("cluster_type")]
        [Validation(Required=true)]
        public string ClusterType { get; set; }

        [NameInMap("name")]
        [Validation(Required=true)]
        public string Name { get; set; }

        [NameInMap("region_id")]
        [Validation(Required=true)]
        public string RegionId { get; set; }

        [NameInMap("kubernetes_version")]
        [Validation(Required=false)]
        public string KubernetesVersion { get; set; }

        [NameInMap("snat_entry")]
        [Validation(Required=true)]
        public bool? SnatEntry { get; set; }

        [NameInMap("endpoint_public_access")]
        [Validation(Required=false)]
        public bool? EndpointPublicAccess { get; set; }

        [NameInMap("ssh_flags")]
        [Validation(Required=false)]
        public bool? SshFlags { get; set; }

        [NameInMap("cloud_monitor_flags")]
        [Validation(Required=false)]
        public bool? CloudMonitorFlags { get; set; }

        [NameInMap("deletion_protection")]
        [Validation(Required=false)]
        public bool? DeletionProtection { get; set; }

        [NameInMap("node_cidr_mask")]
        [Validation(Required=false)]
        public string NodeCidrMask { get; set; }

        [NameInMap("proxy_mode")]
        [Validation(Required=false)]
        public string ProxyMode { get; set; }

        [NameInMap("os_type")]
        [Validation(Required=false)]
        public string OsType { get; set; }

        [NameInMap("platform")]
        [Validation(Required=false)]
        public string Platform { get; set; }

        [NameInMap("node_port_range")]
        [Validation(Required=false)]
        public string NodePortRange { get; set; }

        [NameInMap("key_pair")]
        [Validation(Required=true)]
        public string KeyPair { get; set; }

        [NameInMap("login_password")]
        [Validation(Required=true)]
        public string LoginPassword { get; set; }

        [NameInMap("master_instance_charge_type")]
        [Validation(Required=false)]
        public string MasterInstanceChargeType { get; set; }

        [NameInMap("worker_instance_charge_type")]
        [Validation(Required=false)]
        public string WorkerInstanceChargeType { get; set; }

        [NameInMap("master_period")]
        [Validation(Required=false)]
        public int? MasterPeriod { get; set; }

        [NameInMap("worker_period")]
        [Validation(Required=false)]
        public int? WorkerPeriod { get; set; }

        [NameInMap("master_period_unit")]
        [Validation(Required=false)]
        public string MasterPeriodUnit { get; set; }

        [NameInMap("worker_period_unit")]
        [Validation(Required=false)]
        public string WorkerPeriodUnit { get; set; }

        [NameInMap("master_auto_renew")]
        [Validation(Required=false)]
        public bool? MasterAutoRenew { get; set; }

        [NameInMap("master_auto_renew_period")]
        [Validation(Required=false)]
        public int? MasterAutoRenewPeriod { get; set; }

        [NameInMap("worker_auto_renew")]
        [Validation(Required=false)]
        public bool? WorkerAutoRenew { get; set; }

        [NameInMap("worker_auto_renew_period")]
        [Validation(Required=false)]
        public int? WorkerAutoRenewPeriod { get; set; }

        [NameInMap("cpu_policy")]
        [Validation(Required=false)]
        public string CpuPolicy { get; set; }

        [NameInMap("master_count")]
        [Validation(Required=false)]
        public int? MasterCount { get; set; }

        [NameInMap("master_system_disk_category")]
        [Validation(Required=true)]
        public string MasterSystemDiskCategory { get; set; }

        [NameInMap("master_system_disk_size")]
        [Validation(Required=true)]
        public int? MasterSystemDiskSize { get; set; }

        [NameInMap("runtime")]
        [Validation(Required=false)]
        public Dictionary<string, string> Runtime { get; set; }

        [NameInMap("num_of_nodes")]
        [Validation(Required=true)]
        public int? NumOfNodes { get; set; }

        [NameInMap("worker_system_disk_category")]
        [Validation(Required=true)]
        public string WorkerSystemDiskCategory { get; set; }

        [NameInMap("worker_system_disk_size")]
        [Validation(Required=true)]
        public int? WorkerSystemDiskSize { get; set; }

        [NameInMap("worker_data_disk")]
        [Validation(Required=false)]
        public bool? WorkerDataDisk { get; set; }

        [NameInMap("vpcid")]
        [Validation(Required=false)]
        public string Vpcid { get; set; }

        [NameInMap("security_group_id")]
        [Validation(Required=false)]
        public string SecurityGroupId { get; set; }

        [NameInMap("container_cidr")]
        [Validation(Required=false)]
        public string ContainerCidr { get; set; }

        [NameInMap("service_cidr")]
        [Validation(Required=false)]
        public string ServiceCidr { get; set; }

        [NameInMap("disable_rollback")]
        [Validation(Required=false)]
        public bool? DisableRollback { get; set; }

        [NameInMap("timeout_mins")]
        [Validation(Required=false)]
        public int? TimeoutMins { get; set; }

        [NameInMap("vswitch_ids")]
        [Validation(Required=false)]
        public string VswitchIds { get; set; }

        [NameInMap("private_zone")]
        [Validation(Required=false)]
        public string PrivateZone { get; set; }

        [NameInMap("zone_id")]
        [Validation(Required=false)]
        public string ZoneId { get; set; }

        [NameInMap("profile")]
        [Validation(Required=false)]
        public string Profile { get; set; }

        [NameInMap("is_enterprise_security_group")]
        [Validation(Required=false)]
        public bool? IsEnterpriseSecurityGroup { get; set; }

        [NameInMap("tags")]
        [Validation(Required=false)]
        public List<CreateClusterBodyTags> Tags { get; set; }
        public class CreateClusterBodyTags : TeaModel {
            [NameInMap("key")]
            [Validation(Required=false)]
            public string Key { get; set; }

            [NameInMap("value")]
            [Validation(Required=false)]
            public string Value { get; set; }

        }

        [NameInMap("addons")]
        [Validation(Required=false)]
        public List<CreateClusterBodyAddons> Addons { get; set; }
        public class CreateClusterBodyAddons : TeaModel {
            [NameInMap("name")]
            [Validation(Required=false)]
            public string Name { get; set; }

            [NameInMap("config")]
            [Validation(Required=false)]
            public string Config { get; set; }

            [NameInMap("disabled")]
            [Validation(Required=false)]
            public bool? Disabled { get; set; }

        }

        [NameInMap("taints")]
        [Validation(Required=false)]
        public List<CreateClusterBodyTaints> Taints { get; set; }
        public class CreateClusterBodyTaints : TeaModel {
            [NameInMap("key")]
            [Validation(Required=false)]
            public string Key { get; set; }

            [NameInMap("value")]
            [Validation(Required=false)]
            public string Value { get; set; }

            [NameInMap("effect")]
            [Validation(Required=false)]
            public string Effect { get; set; }

        }

        [NameInMap("worker_data_disks")]
        [Validation(Required=false)]
        public List<CreateClusterBodyWorkerDataDisks> WorkerDataDisks { get; set; }
        public class CreateClusterBodyWorkerDataDisks : TeaModel {
            [NameInMap("category")]
            [Validation(Required=false)]
            public string Category { get; set; }

            [NameInMap("size")]
            [Validation(Required=false)]
            public string Size { get; set; }

            [NameInMap("encrypted")]
            [Validation(Required=false)]
            public string Encrypted { get; set; }

        }

        [NameInMap("master_vswitch_ids")]
        [Validation(Required=false)]
        public List<string> MasterVswitchIds { get; set; }

        [NameInMap("master_instance_types")]
        [Validation(Required=false)]
        public List<string> MasterInstanceTypes { get; set; }

        [NameInMap("worker_instance_types")]
        [Validation(Required=false)]
        public List<string> WorkerInstanceTypes { get; set; }

        [NameInMap("worker_vswitch_ids")]
        [Validation(Required=false)]
        public List<string> WorkerVswitchIds { get; set; }

        [NameInMap("pod_vswitch_ids")]
        [Validation(Required=false)]
        public List<string> PodVswitchIds { get; set; }

    }

}
