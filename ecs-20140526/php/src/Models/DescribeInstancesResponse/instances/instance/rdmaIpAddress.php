<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Ecs\V20140526\Models\DescribeInstancesResponse\instances\instance;

use AlibabaCloud\Tea\Model;

class rdmaIpAddress extends Model
{
    /**
     * @description IpAddress
     *
     * @var array
     */
    public $ipAddress;
    protected $_name = [
        'ipAddress' => 'IpAddress',
    ];

    public function validate()
    {
        Model::validateRequired('ipAddress', $this->ipAddress, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->ipAddress) {
            $res['IpAddress'] = [];
            if (null !== $this->ipAddress) {
                $res['IpAddress'] = $this->ipAddress;
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return rdmaIpAddress
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['IpAddress'])) {
            if (!empty($map['IpAddress'])) {
                $model->ipAddress = [];
                $model->ipAddress = $map['IpAddress'];
            }
        }

        return $model;
    }
}
