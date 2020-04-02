<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Objectdet\V20191230\Objectdet\DetectMainBodyResponse;

use AlibabaCloud\SDK\Objectdet\V20191230\Objectdet\data\location;
use AlibabaCloud\Tea\Model;

class data extends Model
{
    /**
     * @description location
     *
     * @var data.location
     */
    public $location;
    protected $_name = [
        'location' => 'Location',
    ];

    public function validate()
    {
        Model::validateRequired('location', $this->location, true);
    }

    public function toMap()
    {
        $res             = [];
        $res['Location'] = null !== $this->location ? $this->location->toMap() : null;

        return $res;
    }

    /**
     * @param array $map
     *
     * @return data
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Location'])) {
            $model->location = data\location::fromMap($map['Location']);
        }

        return $model;
    }
}
